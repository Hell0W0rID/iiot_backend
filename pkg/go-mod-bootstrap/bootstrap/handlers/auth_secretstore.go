//
//
// SPDX-License-Identifier: Apache-2.0

package handlers

import (
	"fmt"

	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/interfaces"
	"iiot-backend/pkg/go-mod-core-contracts/clients/logger"
	"iiot-backend/pkg/go-mod-core-contracts/errors"

	"github.com/labstack/echo/v4"
)

// SecretStoreAuthenticationHandlerFunc verifies the JWT with a OpenBao-based JWT authentication check
func SecretStoreAuthenticationHandlerFunc(secretProvider interfaces.SecretProviderExt, lc logger.LoggerClient, token string, c echo.Context) errors.IIOT {
	r := c.Request()

	validToken, err := secretProvider.IsJWTValid(token)
	if err != nil {
		lc.Errorf("Error checking JWT validity by the secret provider: %v ", err)
		return errors.NewCommonIIOT(errors.KindServerError, "Error checking JWT validity by the secret provider", err)
	} else if !validToken {
		lc.Warnf("Request to '%s' UNAUTHORIZED", r.URL.Path)
		return errors.NewCommonIIOT(errors.KindUnauthorized, fmt.Sprintf("Request to '%s' UNAUTHORIZED", r.URL.Path), err)
	}
	lc.Debugf("Request to '%s' authorized", r.URL.Path)
	return nil
}
