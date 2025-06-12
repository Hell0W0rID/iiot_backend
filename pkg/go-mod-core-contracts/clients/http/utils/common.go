//
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"iiot-backend/pkg/go-mod-core-contracts/clients/interfaces"
	"iiot-backend/pkg/go-mod-core-contracts/common"
	dtosCommon "iiot-backend/pkg/go-mod-core-contracts/dtos/common"
	"iiot-backend/pkg/go-mod-core-contracts/errors"

	"github.com/google/uuid"
)

// FromContext allows for the retrieval of the specified key's value from the supplied Context.
// If the value is not found, an empty string is returned.
func FromContext(ctx context.Context, key string) string {
	hdr, ok := ctx.Value(key).(string)
	if !ok {
		hdr = ""
	}
	return hdr
}

// correlatedId gets Correlation ID from supplied context. If no Correlation ID header is
// present in the supplied context, one will be created along with a value.
func correlatedId(ctx context.Context) string {
	correlation := FromContext(ctx, common.CorrelationHeader)
	if len(correlation) == 0 {
		correlation = uuid.New().String()
	}
	return correlation
}

// Helper method to get the body from the response after making the request
func getBody(resp *http.Response) ([]byte, errors.IIOT) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return body, errors.NewCommonIIOT(errors.KindIOError, "failed to get the body from the response", err)
	}
	return body, nil
}

// Helper method to make the request and return the response
func makeRequest(req *http.Request, authInjector interfaces.AuthenticationInjector) (*http.Response, errors.IIOT) {
	if authInjector != nil {
		if err := authInjector.AddAuthenticationData(req); err != nil {
			return nil, errors.NewCommonIIOTWrapper(err)
		}
	}

	client := &http.Client{Transport: authInjector.RoundTripper()}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.NewCommonIIOT(errors.KindServiceUnavailable, "failed to send a http request", err)
	}
	if resp == nil {
		return nil, errors.NewCommonIIOT(errors.KindServerError, "the response should not be a nil", nil)
	}
	return resp, nil
}

func CreateRequest(ctx context.Context, httpMethod string, baseUrl string, requestPath string, requestParams url.Values) (*http.Request, errors.IIOT) {
	u, err := parseBaseUrlAndRequestPath(baseUrl, requestPath)
	if err != nil {
		return nil, errors.NewCommonIIOT(errors.KindServerError, "failed to parse baseUrl and requestPath", err)
	}
	if requestParams != nil {
		u.RawQuery = requestParams.Encode()
	}
	req, err := http.NewRequest(httpMethod, u.String(), nil)
	if err != nil {
		return nil, errors.NewCommonIIOT(errors.KindServerError, "failed to create a http request", err)
	}
	req.Header.Set(common.CorrelationHeader, correlatedId(ctx))
	return req, nil
}

func CreateRequestWithRawDataAndParams(ctx context.Context, httpMethod string, baseUrl string, requestPath string, requestParams url.Values, data interface{}) (*http.Request, errors.IIOT) {
	u, err := parseBaseUrlAndRequestPath(baseUrl, requestPath)
	if err != nil {
		return nil, errors.NewCommonIIOT(errors.KindServerError, "failed to parse baseUrl and requestPath", err)
	}
	if requestParams != nil {
		u.RawQuery = requestParams.Encode()
	}
	jsonEncodedData, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewCommonIIOT(errors.KindContractInvalid, "failed to encode input data to JSON", err)
	}

	content := FromContext(ctx, common.ContentType)
	if content == "" {
		content = common.ContentTypeJSON
	}

	req, err := http.NewRequest(httpMethod, u.String(), bytes.NewReader(jsonEncodedData))
	if err != nil {
		return nil, errors.NewCommonIIOT(errors.KindServerError, "failed to create a http request", err)
	}
	req.Header.Set(common.ContentType, content)
	req.Header.Set(common.CorrelationHeader, correlatedId(ctx))
	return req, nil
}

func CreateRequestWithRawData(ctx context.Context, httpMethod string, baseUrl string, requestPath string, requestParams url.Values, data interface{}) (*http.Request, errors.IIOT) {
	u, err := parseBaseUrlAndRequestPath(baseUrl, requestPath)
	if err != nil {
		return nil, errors.NewCommonIIOT(errors.KindServerError, "failed to parse baseUrl and requestPath", err)
	}
	if requestParams != nil {
		u.RawQuery = requestParams.Encode()
	}

	jsonEncodedData, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewCommonIIOT(errors.KindContractInvalid, "failed to encode input data to JSON", err)
	}

	content := FromContext(ctx, common.ContentType)
	if content == "" {
		content = common.ContentTypeJSON
	}

	req, err := http.NewRequest(httpMethod, u.String(), bytes.NewReader(jsonEncodedData))
	if err != nil {
		return nil, errors.NewCommonIIOT(errors.KindServerError, "failed to create a http request", err)
	}
	req.Header.Set(common.ContentType, content)
	req.Header.Set(common.CorrelationHeader, correlatedId(ctx))
	return req, nil
}

func CreateRequestWithRawDataAndHeaders(ctx context.Context, httpMethod string, baseUrl string, requestPath string, requestParams url.Values, data any, headers map[string]string) (*http.Request, errors.IIOT) {
	req, err := CreateRequestWithRawData(ctx, httpMethod, baseUrl, requestPath, requestParams, data)
	if err != nil {
		return nil, errors.NewCommonIIOTWrapper(err)
	}

	// Add the additional headers from request
	for name, value := range headers {
		req.Header.Set(name, value)
	}

	return req, nil
}

func CreateRequestWithEncodedData(ctx context.Context, httpMethod string, baseUrl string, requestPath string, data []byte, encoding string) (*http.Request, errors.IIOT) {
	u, err := parseBaseUrlAndRequestPath(baseUrl, requestPath)
	if err != nil {
		return nil, errors.NewCommonIIOT(errors.KindServerError, "failed to parse baseUrl and requestPath", err)
	}

	content := encoding
	if content == "" {
		content = FromContext(ctx, common.ContentType)
	}

	req, err := http.NewRequest(httpMethod, u.String(), bytes.NewReader(data))
	if err != nil {
		return nil, errors.NewCommonIIOT(errors.KindServerError, "failed to create a http request", err)
	}
	req.Header.Set(common.ContentType, content)
	req.Header.Set(common.CorrelationHeader, correlatedId(ctx))
	return req, nil
}

// CreateRequestFromFilePath creates multipart/form-data request with the specified file
func CreateRequestFromFilePath(ctx context.Context, httpMethod string, baseUrl string, requestPath string, filePath string) (*http.Request, errors.IIOT) {
	u, err := parseBaseUrlAndRequestPath(baseUrl, requestPath)
	if err != nil {
		return nil, errors.NewCommonIIOT(errors.KindServerError, "failed to parse baseUrl and requestPath", err)
	}

	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.NewCommonIIOT(errors.KindIOError, fmt.Sprintf("fail to read file from %s", filePath), err)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	formFileWriter, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, errors.NewCommonIIOT(errors.KindServerError, "fail to create form data", err)
	}
	_, err = io.Copy(formFileWriter, bytes.NewReader(fileContents))
	if err != nil {
		return nil, errors.NewCommonIIOT(errors.KindIOError, "fail to copy file to form data", err)
	}
	writer.Close()

	req, err := http.NewRequest(httpMethod, u.String(), body)
	if err != nil {
		return nil, errors.NewCommonIIOT(errors.KindServerError, "failed to create a http request", err)
	}
	req.Header.Set(common.ContentType, writer.FormDataContentType())
	req.Header.Set(common.CorrelationHeader, correlatedId(ctx))
	return req, nil
}

// SendRequest will make a request with raw data to the specified URL.
// It returns the body as a byte array if successful and an error otherwise.
func SendRequest(ctx context.Context, req *http.Request, authInjector interfaces.AuthenticationInjector) ([]byte, errors.IIOT) {
	resp, err := makeRequest(req, authInjector)
	if err != nil {
		return nil, errors.NewCommonIIOTWrapper(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := getBody(resp)
	if err != nil {
		return nil, errors.NewCommonIIOTWrapper(err)
	}

	if resp.StatusCode <= http.StatusMultiStatus {
		return bodyBytes, nil
	}

	var errMsg string
	var errResp dtosCommon.BaseResponse
	// If the bodyBytes can be unmarshalled to BaseResponse DTO, use the BaseResponse.Message field as the error message
	// Otherwise, use the whole bodyBytes string as the error message
	baseRespErr := json.Unmarshal(bodyBytes, &errResp)
	if baseRespErr == nil {
		errMsg = errResp.Message
	} else {
		errMsg = string(bodyBytes)
	}

	// Handle error response
	msg := fmt.Sprintf("request failed, status code: %d, err: %s", resp.StatusCode, errMsg)
	errKind := errors.KindMapping(resp.StatusCode)
	return bodyBytes, errors.NewCommonIIOT(errKind, msg, nil)
}

// EscapeAndJoinPath escape and join the path variables
func EscapeAndJoinPath(apiRoutePath string, pathVariables ...string) string {
	elements := make([]string, len(pathVariables)+1)
	elements[0] = apiRoutePath // we don't need to escape the route path like /device, /reading, ...,etc.
	for i, e := range pathVariables {
		elements[i+1] = common.URLEncode(e)
	}
	return path.Join(elements...)
}

func parseBaseUrlAndRequestPath(baseUrl, requestPath string) (*url.URL, error) {
	fullPath, err := url.JoinPath(baseUrl, requestPath)
	if err != nil {
		return nil, err
	}
	return url.Parse(fullPath)
}
