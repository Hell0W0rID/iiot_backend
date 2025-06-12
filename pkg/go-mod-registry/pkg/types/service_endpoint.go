//
//
//
//
// Unless required by applicable law or agreed to in writing, software
//

package types

// ServiceEndpoint defines the service information returned by GetServiceEndpoint() need to connect to the target service
type ServiceEndpoint struct {
	ServiceId string
	Host      string
	Port      int
}
