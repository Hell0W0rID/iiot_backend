//
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"iiot-backend/pkg/go-mod-core-contracts/dtos/requests"
	"iiot-backend/pkg/go-mod-core-contracts/dtos/responses"
	"iiot-backend/pkg/go-mod-core-contracts/errors"
)

// KVSClient defines the interface for interactions with the kvs endpoint on the IIOT core-keeper service.
type KVSClient interface {
	// UpdateValuesByKey updates values of the specified key and the child keys defined in the request payload.
	// If no key exists at the given path, the key(s) will be created.
	// If 'flatten' is true, the request json object would be flattened before storing into database.
	UpdateValuesByKey(ctx context.Context, key string, flatten bool, reqs requests.UpdateKeysRequest) (responses.KeysResponse, errors.IIOT)
	// ValuesByKey returns the values of the specified key prefix.
	ValuesByKey(ctx context.Context, key string) (responses.MultiKeyValueResponse, errors.IIOT)
	// ListKeys returns the list of the keys with the specified key prefix.
	ListKeys(ctx context.Context, key string) (responses.KeysResponse, errors.IIOT)
	// DeleteKey deletes the specified key.
	DeleteKey(ctx context.Context, key string) (responses.KeysResponse, errors.IIOT)
	// DeleteKeysByPrefix deletes all keys with the specified prefix.
	DeleteKeysByPrefix(ctx context.Context, key string) (responses.KeysResponse, errors.IIOT)
}
