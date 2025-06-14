package file

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/environment"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/interfaces"
	"iiot-backend/pkg/go-mod-core-contracts/clients/logger"
)

func Load(path string, provider interfaces.SecretProvider, lc logger.LoggerClient) ([]byte, error) {
	var fileBytes []byte
	var err error

	parsedUrl, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("Could not parse file path: %v", err)
	}

	if parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https" {
		client := &http.Client{
			Timeout: environment.GetURIRequestTimeout(lc),
		}
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			return nil, fmt.Errorf("Unable to create new request for remote file: %s: %v", parsedUrl.Redacted(), err)
		}

		// Get httpheader secret
		params := parsedUrl.Query()
		iiotSecretName := params.Get("iiotSecretName")
		if iiotSecretName != "" {
			secrets, err := provider.RetrieveSecret(iiotSecretName)
			if err != nil {
				return nil, err
			}

			// Set request header
			if len(secrets) > 0 && secrets["type"] == "httpheader" {
				if secrets["headername"] != "" && secrets["headercontents"] != "" {
					req.Header.Add(secrets["headername"], secrets["headercontents"])
				} else {
					return nil, fmt.Errorf("Secret headername and headercontents can not be empty")
				}
			} else {
				return nil, fmt.Errorf("Secret type is not httpheader")
			}
		}

		// Run request
		resp, err := client.Do(req)

		if err != nil {
			return nil, fmt.Errorf("Could not get remote file: %s: %v", parsedUrl.Redacted(), err)
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 300 {
			return nil, fmt.Errorf("Invalid status code %d loading remote file: %s", resp.StatusCode, parsedUrl.Redacted())
		}

		fileBytes, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Could not read remote file: %s: %v", parsedUrl.Redacted(), err)
		}
	} else {
		fileBytes, err = os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("Could not read file %s: %v", path, err)
		}
	}

	return fileBytes, nil
}
