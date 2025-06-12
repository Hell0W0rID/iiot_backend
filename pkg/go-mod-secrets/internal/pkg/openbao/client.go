/*******************************************************************************
 *******************************************************************************/

package openbao

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"
	"sync"

	"iiot-backend/pkg/go-mod-secrets/pkg"
	"iiot-backend/pkg/go-mod-secrets/pkg/types"

	"iiot-backend/pkg/go-mod-core-contracts/clients/logger"
)

// Client defines the behavior for interacting with the OpenBao REST secret key/value store via HTTP(S).
type Client struct {
	Config     types.SecretConfig
	HttpCaller pkg.Caller
	lc         logger.LoggerClient
	context    context.Context
	// secretStoreTokenToCancelFuncMap is an internal map with token as key and the context.cancel function as value
	secretStoreTokenToCancelFuncMap secretStoreTokenToCancelFuncMap
	mapMutex                        sync.Mutex
	tokenExpiredCallback            pkg.TokenExpiredCallback
}

// NewClient constructs a secret store *Client which communicates with OpenBao via HTTP(S)
// lc is any logging client that implements the loggingClient interface;
// today IIOT's logger.LoggerClient from go-mod-core-contracts satisfies this implementation
func NewClient(config types.SecretConfig, requester pkg.Caller, forSecrets bool, lc logger.LoggerClient) (*Client, error) {

	if forSecrets && config.Authentication.AuthToken == "" {
		return nil, pkg.NewErrSecretStore("AuthToken is required in config")
	}

	var err error
	if requester == nil {
		requester, err = createHTTPClient(config)
		if err != nil {
			return nil, err
		}
	}

	secretStoreClient := Client{
		Config:                          config,
		HttpCaller:                      requester,
		lc:                              lc,
		mapMutex:                        sync.Mutex{},
		secretStoreTokenToCancelFuncMap: make(secretStoreTokenToCancelFuncMap),
	}

	return &secretStoreClient, err
}

func createHTTPClient(config types.SecretConfig) (pkg.Caller, error) {

	if config.RootCaCertPath == "" {
		return http.DefaultClient, nil
	}

	// Read and load the CA Root certificate so the client will be able to use TLS without skipping the verification of
	// the cert received by the server.
	caCert, err := os.ReadFile(config.RootCaCertPath)
	if err != nil {
		return nil, ErrCaRootCert{
			path:        config.RootCaCertPath,
			description: err.Error(),
		}
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:    caCertPool,
				ServerName: config.ServerName,
				MinVersion: tls.VersionTLS12,
			},
		},
	}, nil
}
