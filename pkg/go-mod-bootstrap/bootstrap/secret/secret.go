/********************************************************************************
 *******************************************************************************/

package secret

import (
	"context"
	"fmt"
	"path"
	"strings"
	"time"

	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/container"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/environment"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/interfaces"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/startup"
	"iiot-backend/pkg/go-mod-bootstrap/config"
	"iiot-backend/pkg/go-mod-bootstrap/di"

	httpClients "iiot-backend/pkg/go-mod-core-contracts/clients/http"
	"iiot-backend/pkg/go-mod-core-contracts/clients/logger"
	"iiot-backend/pkg/go-mod-core-contracts/common"

	"iiot-backend/pkg/go-mod-secrets/pkg/token/authtokenloader"
	"iiot-backend/pkg/go-mod-secrets/pkg/token/fileioperformer"
	"iiot-backend/pkg/go-mod-secrets/pkg/token/runtimetokenprovider"
	"iiot-backend/pkg/go-mod-secrets/pkg/types"
	"iiot-backend/pkg/go-mod-secrets/secrets"

	gometrics "github.com/rcrowley/go-metrics"
)

// secret service Metric Names
const (
	secretsRequestedMetricName             = "SecuritySecretsRequested"
	secretsStoredMetricName                = "SecuritySecretsStored"
	securityRuntimeSecretTokenDurationName = "SecurityRuntimeSecretTokenDuration"
	securityRetrieveSecretDurationName          = "SecurityRetrieveSecretDuration"
)

// NewSecretProvider creates a new fully initialized the Secret Provider.
func NewSecretProvider(
	configuration interfaces.Configuration,
	envVars *environment.Variables,
	ctx context.Context,
	startupTimer startup.Timer,
	dic *di.Container,
	serviceKey string) (interfaces.SecretProviderExt, error) {
	lc := container.LoggerClientFrom(dic.Get)

	var provider interfaces.SecretProviderExt

	switch IsSecurityEnabled() {
	case true:
		// attempt to create a new Secure client only if security is enabled.
		var err error

		lc.Info("Creating SecretClient")

		secretStoreConfig, err := BuildSecretStoreConfig(serviceKey, envVars, lc)
		if err != nil {
			return nil, err
		}

		for startupTimer.HasNotElapsed() {
			var secretConfig types.SecretConfig

			lc.Info("Measurement secret store configuration and authentication token")

			tokenLoader := container.AuthTokenLoaderFrom(dic.Get)
			if tokenLoader == nil {
				tokenLoader = authtokenloader.NewAuthTokenLoader(fileioperformer.NewDefaultFileIoPerformer())
			}

			runtimeTokenLoader := container.RuntimeTokenProviderFrom(dic.Get)
			if runtimeTokenLoader == nil {
				runtimeTokenLoader = runtimetokenprovider.NewRuntimeTokenProvider(ctx, lc,
					secretStoreConfig.RuntimeTokenProvider)
			}

			// We need to create securityRuntimeSecretTokenDuration here because we want to measure the time taken
			// to get the secret config, but the secureProvider instance is created after this step.
			securityRuntimeSecretTokenDuration := gometrics.NewTimer()
			secretConfig, err = getSecretConfig(secretStoreConfig, tokenLoader, runtimeTokenLoader, serviceKey, lc, securityRuntimeSecretTokenDuration)
			if err == nil {
				secureProvider := NewSecureProvider(ctx, secretStoreConfig, lc, tokenLoader, runtimeTokenLoader, serviceKey)
				secureProvider.securityRuntimeSecretTokenDuration = securityRuntimeSecretTokenDuration
				var secretClient secrets.SecretClient

				lc.Info("Attempting to create secret client")

				tokenCallbackFunc := secureProvider.DefaultTokenExpiredCallback
				if secretConfig.RuntimeTokenProvider.Enabled {
					tokenCallbackFunc = secureProvider.RuntimeTokenExpiredCallback
				}

				secretClient, err = secrets.NewSecretsClient(ctx, secretConfig, lc, tokenCallbackFunc)
				if err == nil {
					secureProvider.SetClient(secretClient)
					provider = secureProvider
					lc.Info("Created SecretClient")

					lc.Debugf("SecretsFile is '%s'", secretConfig.SecretsFile)

					if len(strings.TrimSpace(secretConfig.SecretsFile)) == 0 {
						lc.Infof("SecretsFile not set, skipping seeding of service secrets.")
						break
					}

					provider = secureProvider
					lc.Info("Created SecretClient")

					err = secureProvider.LoadServiceSecrets(secretStoreConfig)
					if err != nil {
						return nil, err
					}
					break
				} else if strings.Contains(err.Error(), AccessTokenAuthError) && !secretConfig.RuntimeTokenProvider.Enabled {
					lc.Warnf("token expired, invoking secret-store-setup regen token API ........")

					clientCollection, err := BuildSecretStoreSetupClientConfig(envVars, lc)
					if err != nil {
						return nil, err
					}

					clientConfigs := *clientCollection
					ssSetupClient, ok := clientConfigs[common.SecuritySecretStoreSetupServiceName]
					if !ok {
						return nil, fmt.Errorf("failed to obtain %s client from config", common.SecuritySecretStoreSetupServiceName)
					}
					baseUrl := ssSetupClient.Url()

					entityId, err := tokenLoader.ReadEntityId(secretStoreConfig.TokenFile)
					if err != nil {
						return nil, err
					}

					// Use InsecureProvider here since the client token has been expired and cannot be used to obtain the JWT
					secretProvider := NewInsecureProvider(configuration, lc, dic)
					jwtProvider := NewJWTSecretProvider(secretProvider)
					httpClient := httpClients.NewSecretStoreTokenClient(baseUrl, jwtProvider)
					_, err = httpClient.RegenToken(ctx, entityId)
					if err != nil {
						return nil, err
					}

					lc.Infof("token file re-generated, trying to create the secret client again later")
				}
			}

			lc.Warn(fmt.Sprintf("Retryable failure while creating SecretClient: %s", err.Error()))
			startupTimer.SleepForInterval()
		}

		if err != nil {
			return nil, fmt.Errorf("unable to create SecretClient: %s", err.Error())
		}

	case false:
		provider = NewInsecureProvider(configuration, lc, dic)
	}

	dic.Update(di.ServiceConstructorMap{
		// Must put the SecretProvider instance in the DIC for both the standard API use by service code
		// and the extended API used by boostrap code
		container.SecretProviderName: func(get di.Get) interface{} {
			return provider
		},
		container.SecretProviderExtName: func(get di.Get) interface{} {
			return provider
		},
	})

	return provider, nil
}

// BuildSecretStoreConfig is public helper function that builds the SecretStore configuration
// from default values and  environment override.
func BuildSecretStoreConfig(serviceKey string, envVars *environment.Variables, lc logger.LoggerClient) (*config.SecretStoreInfo, error) {
	configWrapper := struct {
		SecretStore config.SecretStoreInfo
	}{
		SecretStore: config.NewSecretStoreInfo(serviceKey),
	}

	count, err := envVars.OverrideConfiguration(&configWrapper)
	if err != nil {
		return nil, fmt.Errorf("failed to override SecretStore information: %v", err)
	}

	lc.Infof("SecretStore information created with %d overrides applied", count)
	return &configWrapper.SecretStore, nil
}

// getSecretConfig creates a SecretConfig based on the SecretStoreInfo configuration properties.
// If a token file is present it will override the Authentication.AuthToken value.
func getSecretConfig(secretStoreInfo *config.SecretStoreInfo,
	tokenLoader authtokenloader.AuthTokenLoader,
	runtimeTokenLoader runtimetokenprovider.RuntimeTokenProvider,
	serviceKey string,
	lc logger.LoggerClient,
	securityRuntimeSecretTokenDuration gometrics.Timer) (types.SecretConfig, error) {
	secretConfig := types.SecretConfig{
		Type:                 secretStoreInfo.Type, // Type of SecretStore implementation, i.e. OpenBao
		Host:                 secretStoreInfo.Host,
		Port:                 secretStoreInfo.Port,
		BasePath:             addIIOTSecretNamePrefix(secretStoreInfo.StoreName),
		SecretsFile:          secretStoreInfo.SecretsFile,
		Protocol:             secretStoreInfo.Protocol,
		Namespace:            secretStoreInfo.Namespace,
		RootCaCertPath:       secretStoreInfo.RootCaCertPath,
		ServerName:           secretStoreInfo.ServerName,
		Authentication:       secretStoreInfo.Authentication,
		RuntimeTokenProvider: secretStoreInfo.RuntimeTokenProvider,
	}

	// maybe insecure mode
	// if both configs of token file and runtime token provider are empty or disabled
	// then we treat that as insecure mode
	if !IsSecurityEnabled() || (secretStoreInfo.TokenFile == "" && !secretConfig.RuntimeTokenProvider.Enabled) {
		lc.Info("insecure mode")
		return secretConfig, nil
	}

	// based on whether token provider config is configured or not, we will obtain token in different way
	var token string
	var err error
	if secretConfig.RuntimeTokenProvider.Enabled {
		lc.Info("runtime token provider enabled")
		// call spiffe token provider to get token on the fly
		started := time.Now()
		token, err = runtimeTokenLoader.GetRawToken(serviceKey)
		securityRuntimeSecretTokenDuration.UpdateSince(started)
	} else {
		lc.Info("load token from file")
		// else obtain the token from TokenFile
		token, err = tokenLoader.Load(secretStoreInfo.TokenFile)
	}

	if err != nil {
		return secretConfig, err
	}

	secretConfig.Authentication.AuthToken = token
	return secretConfig, nil
}

func addIIOTSecretNamePrefix(secretName string) string {
	trimmedSecretName := strings.TrimSpace(secretName)

	// in this case, treat it as no secret name prefix
	if len(trimmedSecretName) == 0 {
		return ""
	}

	return "/" + path.Join("v1", "secret", "iiot", trimmedSecretName)
}

// BuildSecretStoreSetupClientConfig is public helper function that builds the ClientsCollection configuration
// from default values and environment override.
func BuildSecretStoreSetupClientConfig(envVars *environment.Variables, lc logger.LoggerClient) (*config.ClientsCollection, error) {
	configWrapper := struct {
		Clients *config.ClientsCollection
	}{
		Clients: config.NewSecretStoreSetupClientInfo(),
	}

	count, err := envVars.OverrideConfiguration(&configWrapper)
	if err != nil {
		return nil, fmt.Errorf("failed to override SecretStore information: %v", err)
	}

	lc.Infof("SecretStoreSetup client information created with %d overrides applied", count)
	return configWrapper.Clients, nil
}
