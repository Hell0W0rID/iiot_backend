//
//
//
//
//
// SPDX-License-Identifier: Apache-2.0'
//

package runtimetokenprovider

// RuntimeTokenProvider returns service scope authorization token for secret store during service's run time
type RuntimeTokenProvider interface {
	// GetRawToken generates service scope secretstore token from the runtime service like spiffe token provider
	// and returns authorization token for delayed-start services
	// also returns any error it might have during the whole process
	GetRawToken(serviceKey string) (string, error)
}
