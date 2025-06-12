/*******************************************************************************
 *******************************************************************************/

package interfaces

import "iiot-backend/pkg/go-mod-bootstrap/config"

// CredentialsProvider interface provides an abstraction for obtaining credentials.
type CredentialsProvider interface {
	// GetDatabaseCredentials retrieves database credentials.
	GetDatabaseCredentials(database config.Database) (config.Credentials, error)
}
