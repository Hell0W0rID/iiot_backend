/*******************************************************************************
 *******************************************************************************/

package openbao

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"

	"iiot-backend/pkg/go-mod-secrets/pkg/types"
)

// GetMockTokenServer returns a stub http test server for dealing with token lookup-self and renew-self API calls
func GetMockTokenServer(tokenDataMap *sync.Map) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		urlPath := req.URL.String()
		if req.Method == http.MethodGet && urlPath == "/v1/auth/token/lookup-self" {
			token := req.Header.Get(AuthTypeHeader)
			sampleTokenLookup, exists := tokenDataMap.Load(token)
			if !exists {
				rw.WriteHeader(403)
				_, _ = rw.Write([]byte("permission denied"))
			} else {
				resp := sampleTokenLookup.(TokenLookupResponse)
				if ret, err := json.Marshal(resp); err != nil {
					rw.WriteHeader(500)
					_, _ = rw.Write([]byte(err.Error()))
				} else {
					rw.WriteHeader(200)
					_, _ = rw.Write(ret)
				}
			}
		} else if req.Method == http.MethodPost && urlPath == "/v1/auth/token/renew-self" {
			token := req.Header.Get(AuthTypeHeader)
			sampleTokenLookup, exists := tokenDataMap.Load(token)
			if !exists {
				rw.WriteHeader(403)
				_, _ = rw.Write([]byte("permission denied"))
			} else {
				currentTTL := sampleTokenLookup.(TokenLookupResponse).Data.Ttl
				if currentTTL <= 0 {
					// already expired
					rw.WriteHeader(403)
					_, _ = rw.Write([]byte("permission denied"))
				} else {
					tokenPeriod := sampleTokenLookup.(TokenLookupResponse).Data.Period

					tokenDataMap.Store(token, TokenLookupResponse{
						Data: types.TokenMetadata{
							Renewable: true,
							Ttl:       tokenPeriod,
							Period:    tokenPeriod,
						},
					})
					rw.WriteHeader(200)
					_, _ = rw.Write([]byte("token renewed"))
				}
			}
		} else {
			rw.WriteHeader(404)
			_, _ = rw.Write([]byte(fmt.Sprintf("Unknown urlPath: %s", urlPath)))
		}
	}))
	return server
}
