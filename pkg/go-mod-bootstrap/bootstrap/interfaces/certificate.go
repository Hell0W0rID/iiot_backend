/*******************************************************************************
 *******************************************************************************/

package interfaces

import "iiot-backend/pkg/go-mod-bootstrap/config"

// CertificateProvider interface provides an abstraction for obtaining certificate pair.
type CertificateProvider interface {
	// GetCertificateKeyPair retrieves certificate pair.
	GetCertificateKeyPair(path string) (config.CertKeyPair, error)
}
