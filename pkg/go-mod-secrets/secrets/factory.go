/*******************************************************************************
 *******************************************************************************/

package secrets

import (
	"context"
	"fmt"

	"iiot-backend/pkg/go-mod-secrets/internal/pkg/openbao"
	"iiot-backend/pkg/go-mod-secrets/pkg"
	"iiot-backend/pkg/go-mod-secrets/pkg/types"

	"iiot-backend/pkg/go-mod-core-contracts/clients/logger"
)

const DefaultSecretStore = "openbao"

// NewSecretsClient creates a new instance of a SecretClient based on the passed in configuration.
// The SecretClient allows access to secret(s) for the configured token.
func NewSecretsClient(ctx context.Context, config types.SecretConfig, lc logger.LoggerClient, callback pkg.TokenExpiredCallback) (SecretClient, error) {
	if ctx == nil {
		return nil, pkg.NewErrSecretStore("background ctx is required and cannot be nil")
	}

	// Currently only have one secret store type implementation, so no need to have/check type.

	switch config.Type {
	// Currently only have one secret store type implementation, so type isn't actual set in configuration
	case DefaultSecretStore:
		return openbao.NewSecretsClient(ctx, config, lc, callback)
	default:
		return nil, fmt.Errorf("invalid secrets client type of '%s'", config.Type)
	}
}

// NewSecretStoreClient creates a new instance of a SecretClient based on the passed in configuration.
// The SecretStoreClient provides management functionality to manage the secret store.
func NewSecretStoreClient(config types.SecretConfig, lc logger.LoggerClient, requester pkg.Caller) (SecretStoreClient, error) {
	switch config.Type {
	case DefaultSecretStore:
		return openbao.NewClient(config, requester, false, lc)

	default:
		return nil, fmt.Errorf("invalid secret store client type of '%s'", config.Type)
	}
}
