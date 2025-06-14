/*******************************************************************************
 *******************************************************************************/

// Package messaging contains common constants and utilities functions used when setting up Secure MessageBus.
// A common bootstrap handler can not be here due to the fact that it would pull in dependency on go-mod-messaging
// which is dependent on ZMQ. This causes all services that use go-mod-boostrap to them have a dependency on ZMQ which
// breaks the security bootstrapping.
package messaging

import (
	"crypto/x509"
	"errors"
	"fmt"

	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/container"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/secret"
	"iiot-backend/pkg/go-mod-bootstrap/config"
	"iiot-backend/pkg/go-mod-bootstrap/di"
	"iiot-backend/pkg/go-mod-core-contracts/clients/logger"
)

const (
	AuthModeNone             = "none"
	AuthModeUsernamePassword = "usernamepassword"
	AuthModeCert             = "clientcert"
	AuthModeCA               = "cacert"

	SecretUsernameKey = "username"
	SecretPasswordKey = "password"
	SecretClientKey   = "clientkey"
	SecretClientCert  = AuthModeCert
	SecretCACert      = AuthModeCA

	OptionsUsernameKey     = "Username"
	OptionsPasswordKey     = "Password"
	OptionsCertPEMBlockKey = "CertPEMBlock"
	OptionsKeyPEMBlockKey  = "KeyPEMBlock"
	OptionsCaPEMBlockKey   = "CaPEMBlock"
)

type SecretDataProvider interface {
	// RetrieveSecret retrieves secrets from the service's SecretStore at the specified path.
	RetrieveSecret(path string, keys ...string) (map[string]string, error)
}

type SecretData struct {
	Username     string
	Password     string
	KeyPemBlock  []byte
	CertPemBlock []byte
	CaPemBlock   []byte
}

func SetOptionsAuthData(messageBusInfo *config.MessageBusInfo, lc logger.LoggerClient, dic *di.Container) error {
	lc.Infof("Setting options for secure MessageBus with AuthMode='%s' and SecretName='%s",
		messageBusInfo.AuthMode,
		messageBusInfo.SecretName)

	secretProvider := container.SecretProviderFrom(dic.Get)
	if secretProvider == nil {
		return errors.New("secret provider is missing. Make sure it is specified to be used in bootstrap.Run()")
	}

	secretData, err := RetrieveSecretData(messageBusInfo.AuthMode, messageBusInfo.SecretName, secretProvider)
	if err != nil {
		return fmt.Errorf("Unable to get Secret Data for secure message bus: %w", err)
	}

	if err := ValidateSecretData(messageBusInfo.AuthMode, messageBusInfo.SecretName, secretData); err != nil {
		return fmt.Errorf("Secret Data for secure message bus invalid: %w", err)
	}

	if messageBusInfo.Optional == nil {
		messageBusInfo.Optional = map[string]string{}
	}

	// Since already validated, these are the only modes that can be set at this point.
	switch messageBusInfo.AuthMode {
	case AuthModeUsernamePassword:
		messageBusInfo.Optional[OptionsUsernameKey] = secretData.Username
		messageBusInfo.Optional[OptionsPasswordKey] = secretData.Password
	case AuthModeCert:
		messageBusInfo.Optional[OptionsCertPEMBlockKey] = string(secretData.CertPemBlock)
		messageBusInfo.Optional[OptionsKeyPEMBlockKey] = string(secretData.KeyPemBlock)
	}

	if len(secretData.CaPemBlock) > 0 {
		messageBusInfo.Optional[OptionsCaPEMBlockKey] = string(secretData.CaPemBlock)
	}

	return nil
}

func RetrieveSecretData(authMode string, secretName string, provider SecretDataProvider) (*SecretData, error) {
	// No Auth? No Problem!...No secrets required.
	if authMode == AuthModeNone {
		return nil, nil
	}

	secrets, err := provider.RetrieveSecret(secretName)
	if err != nil {
		return nil, err
	}
	data := &SecretData{
		Username:     secrets[SecretUsernameKey],
		Password:     secrets[SecretPasswordKey],
		KeyPemBlock:  []byte(secrets[SecretClientKey]),
		CertPemBlock: []byte(secrets[SecretClientCert]),
		CaPemBlock:   []byte(secrets[SecretCACert]),
	}

	return data, nil
}

func ValidateSecretData(authMode string, secretName string, secretData *SecretData) error {
	switch authMode {
	case AuthModeUsernamePassword:
		if secret.IsSecurityEnabled() && (secretData.Username == "" || secretData.Password == "") {
			return fmt.Errorf("AuthModeUsernamePassword selected however Username or Password was not found for secret=%s", secretName)
		}

	case AuthModeCert:
		// need both to make a successful connection
		if len(secretData.KeyPemBlock) <= 0 || len(secretData.CertPemBlock) <= 0 {
			return fmt.Errorf("AuthModeCert selected however the key or cert PEM block was not found for secret=%s", secretName)
		}

	case AuthModeCA:
		if len(secretData.CaPemBlock) <= 0 {
			return fmt.Errorf("AuthModeCA selected however no PEM Block was found for secret=%s", secretName)
		}

	case AuthModeNone:
		// Nothing to validate
	default:
		return fmt.Errorf("Invalid AuthMode of '%s' selected", authMode)
	}

	if len(secretData.CaPemBlock) > 0 {
		caCertPool := x509.NewCertPool()
		ok := caCertPool.AppendCertsFromPEM(secretData.CaPemBlock)
		if !ok {
			return errors.New("Error parsing CA Certificate")
		}
	}

	return nil
}
