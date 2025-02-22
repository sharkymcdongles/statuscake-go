/*
 * StatusCake API
 *
 * Copyright (c) 2022
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to
 * deal in the Software without restriction, including without limitation the
 * rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
 * sell copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 * FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
 * IN THE SOFTWARE.
 *
 * API version: 1.0.0
 * Contact: support@statuscake.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package statuscake

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"time"
)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)

	// ContextServerIndex uses a server configuration from the index.
	ContextServerIndex = contextKey("serverIndex")

	// ContextOperationServerIndices uses a server configuration from the index mapping.
	ContextOperationServerIndices = contextKey("serverOperationIndices")

	// ContextServerVariables overrides a server configuration variables.
	ContextServerVariables = contextKey("serverVariables")

	// ContextOperationServerVariables overrides a server configuration variables using operation specific values.
	ContextOperationServerVariables = contextKey("serverOperationVariables")
)

// contextKeys are used to identify the type of value in the context. Since
// these are strings, it is possible to get a short description of the context
// key for logging and debugging using key.String().
type contextKey string

func (c contextKey) String() string {
	switch c {
	case ContextServerIndex:
		return "server index"
	case ContextOperationServerIndices:
		return "server operation indicies"
	case ContextServerVariables:
		return "server variables"
	case ContextOperationServerVariables:
		return "server operation variables"
	default:
		return "unknown"
	}
}

// Client manages communication with the StatusCake API API v1.0.0
// In most cases there should be only one, shared, Client.
type Client struct {
	options options
	common  service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	// These are embedded so that from a user perspective they do not appear as
	// separate APIs. Since the operation IDs need to be unique within a
	// specification this will not cause function conflicts.
	ContactGroupsAPI
	LocationsAPI
	MaintenanceWindowsAPI
	PagespeedAPI
	SslAPI
	UptimeAPI
}

type service struct {
	client *Client
}

// NewClient creates a new API client. Additional options may be given to
// override default configuration values.
func NewClient(opts ...Option) *Client {
	c := &Client{}
	c.options = applyOptions(opts)
	c.common.client = c

	// API Services
	c.ContactGroupsAPI = (*ContactGroupsService)(&c.common)
	c.LocationsAPI = (*LocationsService)(&c.common)
	c.MaintenanceWindowsAPI = (*MaintenanceWindowsService)(&c.common)
	c.PagespeedAPI = (*PagespeedService)(&c.common)
	c.SslAPI = (*SslService)(&c.common)
	c.UptimeAPI = (*UptimeService)(&c.common)

	return c
}

// selectHeaderContentType select a content type from the available list.
// TODO: more intelligently select content type.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}

	if contains(contentTypes, "application/json") {
		return "application/json"
	}

	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return.
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// parameterToString converts an interface{} parameter to a string.
func parameterToString(obj interface{}) string {
	if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(jsonBuf), err
}

// callAPI wraps debug information around a HTTP request.
func (c *Client) callAPI(request *http.Request) (*http.Response, error) {
	if c.options.disableRetry {
		return c.do(request)
	}

	ctx := request.Context()
	backoffIdx := 0

	var res *http.Response
	var err error

	for backoffIdx <= c.options.maxRetries {
		res, err = c.do(request)
		if err != nil && backoffIdx == c.options.maxRetries {
			break
		}
		if err != nil {
			backoffFor := c.options.backoff.Backoff(backoffIdx)

			timer := time.NewTimer(backoffFor)
			select {
			case <-timer.C:
				backoffIdx++
				continue
			case <-ctx.Done():
				timer.Stop()
			}
		}

		return res, err
	}

	return res, fmt.Errorf("retries exceeded: %w", err)
}

func (c *Client) do(request *http.Request) (*http.Response, error) {
	if c.options.debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}

		log.Printf("\n%s\n", string(dump))
	}

	res, err := c.options.client.Do(request)
	if err != nil {
		return res, err
	}

	if c.options.debug {
		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			return res, err
		}

		log.Printf("\n%s\n", string(dump))
	}

	return res, err
}

// prepareRequest build the request.
func (c *Client) prepareRequest(ctx context.Context, path string, method string, postBody interface{}, headerParams map[string]string, queryParams url.Values, formParams url.Values, fieldName string, fileName string, fileBytes []byte) (req *http.Request, err error) {
	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// Add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("cannot specify postBody and multipart form at the same time")
		}

		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if !strings.HasPrefix(k, "@") { // Form value
					w.WriteField(k, iv)
					continue
				}

				// File
				if err := addFile(w, k[1:], iv); err != nil {
					return nil, err
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			part, err := w.CreateFormFile(fieldName, filepath.Base(fileName))
			if err != nil {
				return nil, err
			}

			if _, err := part.Write(fileBytes); err != nil {
				return nil, err
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("cannot specify postBody and x-www-form-urlencoded form at the same time")
		}

		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Override request host. This can be useful for testing purposes.
	if c.options.host != "" {
		urlOverride, err := url.Parse(c.options.host)
		if err != nil {
			return nil, err
		}

		url.Scheme = urlOverride.Scheme
		url.Host = urlOverride.Host
	}

	// Generate a new request. It is imperitive the check for `nil` is done here.
	// It may appear that if body is never initialised then it will be `nil` but
	// interfaces in Go are only `nil` if both their value and type are `nil`. If
	// instead we always pass body then the `http` package will panic.
	if body != nil {
		req, err = http.NewRequest(method, url.String(), body)
	} else {
		req, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// Add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		req.Header = headers
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if c.options.credentials != nil {
		c.options.credentials.AddCredentials(req)
	}

	// Add the user agent to the request.
	req.Header.Add(http.CanonicalHeaderKey("User-Agent"), c.options.userAgent)

	// Apply default, static, headers. These will be appended to existing headers
	// set by client options such as authentication credentials, and any header
	// parameters previously applied to the request.
	//
	// Header names are taken verbatim from the consumer and not transoformed
	// into canonical names. If it is important that headers are named in the
	// correct HTTP format this must be done by the caller.
	for header, values := range c.options.headers {
		for _, value := range values {
			req.Header.Add(header, value)
		}
	}

	return req, nil
}

func (c *Client) decode(v interface{}, b []byte, contentType string) error {
	if len(b) == 0 {
		return nil
	}

	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}

	if f, ok := v.(**os.File); ok {
		var err error

		*f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return err
		}

		_, err = (*f).Write(b)
		if err != nil {
			return err
		}
		_, err = (*f).Seek(0, io.SeekStart)

		return err
	}

	if xmlCheck.MatchString(contentType) {
		if err := xml.Unmarshal(b, v); err != nil {
			return err
		}

		return nil
	}

	if jsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err := unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err := json.Unmarshal(b, v); err != nil { // simple model
			return err
		}

		return nil
	}

	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	return err
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(**os.File); ok {
		_, err = bodyBuf.ReadFrom(*fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		err = xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s", contentType)
		return nil, err
	}

	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// ServerURL returns URL based on server settings.
func (c *Client) ServerURL(index int, variables map[string]string) (string, error) {
	return c.options.servers.URL(index, variables)
}

// ServerURLWithContext returns a new server URL given an endpoint.
func (c *Client) ServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	sc, ok := c.options.operationServers[endpoint]
	if !ok {
		sc = c.options.servers
	}

	if ctx == nil {
		return sc.URL(0, nil)
	}

	index, err := getServerOperationIndex(ctx, endpoint)
	if err != nil {
		return "", err
	}

	variables, err := getServerOperationVariables(ctx, endpoint)
	if err != nil {
		return "", err
	}

	return sc.URL(index, variables)
}

func getServerIndex(ctx context.Context) (int, error) {
	si := ctx.Value(ContextServerIndex)
	if si != nil {
		if index, ok := si.(int); ok {
			return index, nil
		}
		return 0, fmt.Errorf("invalid type %T should be int", si)
	}
	return 0, nil
}

func getServerOperationIndex(ctx context.Context, endpoint string) (int, error) {
	osi := ctx.Value(ContextOperationServerIndices)
	if osi != nil {
		if operationIndices, ok := osi.(map[string]int); !ok {
			return 0, fmt.Errorf("invalid type %T should be map[string]int", osi)
		} else {
			index, ok := operationIndices[endpoint]
			if ok {
				return index, nil
			}
		}
	}
	return getServerIndex(ctx)
}

func getServerVariables(ctx context.Context) (map[string]string, error) {
	sv := ctx.Value(ContextServerVariables)
	if sv != nil {
		if variables, ok := sv.(map[string]string); ok {
			return variables, nil
		}
		return nil, fmt.Errorf("ctx value of ContextServerVariables has invalid type %T should be map[string]string", sv)
	}
	return nil, nil
}

func getServerOperationVariables(ctx context.Context, endpoint string) (map[string]string, error) {
	osv := ctx.Value(ContextOperationServerVariables)
	if osv != nil {
		if operationVariables, ok := osv.(map[string]map[string]string); !ok {
			return nil, fmt.Errorf("ctx value of ContextOperationServerVariables has invalid type %T should be map[string]map[string]string", osv)
		} else {
			variables, ok := operationVariables[endpoint]
			if ok {
				return variables, nil
			}
		}
	}
	return getServerVariables(ctx)
}
