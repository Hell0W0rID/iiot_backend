/*******************************************************************************
 *******************************************************************************/

package pkg

import "net/http"

// Caller interface used to abstract the implementation details for issuing an HTTP request. This allows for easier testing by the way of mocks.
type Caller interface {
	Do(req *http.Request) (*http.Response, error)
}

// TokenExpiredCallback is the callback function to handle the case when the secret store token has already expired
type TokenExpiredCallback func(expiredToken string) (replacementToken string, retry bool)
