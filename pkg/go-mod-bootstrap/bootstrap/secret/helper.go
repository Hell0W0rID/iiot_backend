/*******************************************************************************
 *******************************************************************************/

package secret

import "os"

const (
	EnvSecretStore = "IIOT_SECURITY_SECRET_STORE"
	UsernameKey    = "username"
	PasswordKey    = "password"
	// WildcardName is a special secret name that can be used to register a secret callback for any secret.
	WildcardName = "*"
)

// IsSecurityEnabled determines if security has been enabled.
func IsSecurityEnabled() bool {
	env := os.Getenv(EnvSecretStore)
	return env != "false" // Any other value is considered secure mode enabled
}
