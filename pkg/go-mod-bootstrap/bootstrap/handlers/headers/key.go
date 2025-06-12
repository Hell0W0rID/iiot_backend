//
//
// SPDX-License-Identifier: Apache-2.0

package headers

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"sync"

	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/container"
	"iiot-backend/pkg/go-mod-bootstrap/di"
	"iiot-backend/pkg/go-mod-core-contracts/clients/logger"
	"iiot-backend/pkg/go-mod-core-contracts/errors"

	"github.com/golang-jwt/jwt/v5"
)

// A key cache to store the verification keys by issuer
var (
	keysCache = make(map[string]any)
	mutex     sync.RWMutex
)

// GetVerificationKey returns the verification key obtained from local cache or security-proxy-auth http client
func GetVerificationKey(dic *di.Container, issuer, alg string, ctx context.Context) (any, errors.IIOT) {
	lc := container.LoggerClientFrom(dic.Get)
	var verifyKey any

	// Check if the verification of the issuer already exists
	mutex.RLock()
	key, ok := keysCache[issuer]
	mutex.RUnlock()

	if ok {
		lc.Debugf("obtaining verification key from cache for JWT issuer '%s'", issuer)

		verifyKey = key
	} else {
		lc.Debugf("obtaining verification key from proxy-auth service client for JWT issuer '%s'", issuer)

		authClient := container.SecurityProxyAuthClientFrom(dic.Get)
		keyResponse, iiotErr := authClient.VerificationKeyByIssuer(ctx, issuer)
		if iiotErr != nil {
			if errors.Kind(iiotErr) == errors.KindEntityDoesNotExist {
				return nil, errors.NewCommonIIOT(errors.KindServerError, fmt.Sprintf("verification key not found from proxy-auth service for JWT issuer '%s'", issuer), nil)
			}
			return nil, errors.NewCommonIIOT(errors.KindServerError, fmt.Sprintf("failed to obtain the verification key from proxy-auth service for JWT issuer '%s'", issuer), iiotErr)
		}
		verifyKey, iiotErr = ProcessVerificationKey(keyResponse.KeyData.Key, alg, lc)
		if iiotErr != nil {
			return nil, errors.NewCommonIIOT(errors.KindServerError, fmt.Sprintf("failed to process the verification key from proxy-auth service for JWT issuer '%s'", issuer), iiotErr)
		}

		mutex.Lock()
		keysCache[issuer] = verifyKey
		mutex.Unlock()
	}
	return verifyKey, nil
}

// ProcessVerificationKey handles the verification key retrieved from security-proxy-auth and returns the public key in the appropriate format according to the JWT signing algorithm
func ProcessVerificationKey(keyString string, alg string, lc logger.LoggerClient) (any, errors.IIOT) {
	keyBytes := []byte(keyString)

	switch alg {
	case jwt.SigningMethodHS256.Alg(), jwt.SigningMethodHS384.Alg(), jwt.SigningMethodHS512.Alg():
		binaryKey, err := base64.StdEncoding.DecodeString(keyString)
		if err != nil {
			lc.Debugf("the key is not a valid base64, err: '%v', using the key '%s' without base64 encoding.", err, keyString)
			return keyBytes, nil
		}

		return binaryKey, nil
	case jwt.SigningMethodEdDSA.Alg():
		block, _ := pem.Decode(keyBytes)
		if block == nil || block.Type != "PUBLIC KEY" {
			return nil, errors.NewCommonIIOT(errors.KindServerError, "failed to decode the verification key PEM block", nil)
		}

		edPublicKey := ed25519.PublicKey(block.Bytes)
		return edPublicKey, nil
	case jwt.SigningMethodRS256.Alg(), jwt.SigningMethodRS384.Alg(), jwt.SigningMethodRS512.Alg(),
		jwt.SigningMethodPS256.Alg(), jwt.SigningMethodPS384.Alg(), jwt.SigningMethodPS512.Alg():
		rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(keyBytes)
		if err != nil {
			return nil, errors.NewCommonIIOT(errors.KindServerError, fmt.Sprintf("failed to parse '%s' rsa verification key", alg), err)
		}

		return rsaPublicKey, nil
	case jwt.SigningMethodES256.Alg(), jwt.SigningMethodES384.Alg(), jwt.SigningMethodES512.Alg():
		ecdsaPublicKey, err := jwt.ParseECPublicKeyFromPEM(keyBytes)
		if err != nil {
			return nil, errors.NewCommonIIOT(errors.KindServerError, fmt.Sprintf("failed to parse '%s' es verification key", alg), err)
		}

		return ecdsaPublicKey, nil
	default:
		return nil, errors.NewCommonIIOT(errors.KindContractInvalid, fmt.Sprintf("unsupported signing algorithm '%s'", alg), nil)
	}
}
