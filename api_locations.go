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
 * API version: 1.0.0-beta.3
 * Contact: support@statuscake.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package statuscake

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Linger please
var _ context.Context

// LocationsAPI describes the necessary methods to adhere to this interface.
type LocationsAPI interface {
	ListPagespeedMonitoringLocations(ctx context.Context) APIListPagespeedMonitoringLocationsRequest
	ListPagespeedMonitoringLocationsExecute(r APIListPagespeedMonitoringLocationsRequest) (MonitoringLocations, error)
	ListUptimeMonitoringLocations(ctx context.Context) APIListUptimeMonitoringLocationsRequest
	ListUptimeMonitoringLocationsExecute(r APIListUptimeMonitoringLocationsRequest) (MonitoringLocations, error)
}

// LocationsService Locations service.
type LocationsService service

// APIListPagespeedMonitoringLocationsRequest represents a request type.
type APIListPagespeedMonitoringLocationsRequest struct {
	ctx        context.Context
	APIService LocationsAPI
	best       *bool
	location   *string
}

// Best sets best on the request type.
func (r APIListPagespeedMonitoringLocationsRequest) Best(best bool) APIListPagespeedMonitoringLocationsRequest {
	r.best = &best
	return r
}

// Location sets location on the request type.
func (r APIListPagespeedMonitoringLocationsRequest) Location(location string) APIListPagespeedMonitoringLocationsRequest {
	r.location = &location
	return r
}

// Execute executes the request.
func (r APIListPagespeedMonitoringLocationsRequest) Execute() (MonitoringLocations, error) {
	return r.APIService.ListPagespeedMonitoringLocationsExecute(r)
}

// ListPagespeedMonitoringLocations Get all pagespeed monitoring locations.
func (a *LocationsService) ListPagespeedMonitoringLocations(ctx context.Context) APIListPagespeedMonitoringLocationsRequest {
	return APIListPagespeedMonitoringLocationsRequest{
		ctx:        ctx,
		APIService: a,
	}
}

// ListPagespeedMonitoringLocationsWithData Get all pagespeed monitoring locations.
// The use of this method is discouraged as it does not provide the level of
// type safety afforded by the field methods on the request type.
func (a *LocationsService) ListPagespeedMonitoringLocationsWithData(ctx context.Context, m map[string]interface{}) APIListPagespeedMonitoringLocationsRequest {
	r := a.ListPagespeedMonitoringLocations(ctx)
	return r
}

// Execute executes the request.
func (a *LocationsService) ListPagespeedMonitoringLocationsExecute(r APIListPagespeedMonitoringLocationsRequest) (MonitoringLocations, error) {
	var (
		requestBody          interface{}
		requestFormFieldName string
		requestFileName      string
		requestFileBytes     []byte
		returnValue          MonitoringLocations
	)

	basePath, err := a.client.ServerURLWithContext(r.ctx, "LocationsService.ListPagespeedMonitoringLocations")
	if err != nil {
		return returnValue, err
	}

	requestPath := basePath + "/pagespeed-locations"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if r.best != nil {
		queryParams.Add("best", parameterToString(*r.best))
	}
	if r.location != nil {
		queryParams.Add("location", parameterToString(*r.location))
	}
	// Determine the Content-Type header.
	contentTypes := []string{}

	// Set Content-Type header
	requestContentTypeHeader := selectHeaderContentType(contentTypes)
	if requestContentTypeHeader != "" {
		headerParams["Content-Type"] = requestContentTypeHeader
	}

	// Determine the Accept header.
	accepts := []string{"application/json"}

	// Set Accept header.
	requestAcceptHeader := selectHeaderAccept(accepts)
	if requestAcceptHeader != "" {
		headerParams["Accept"] = requestAcceptHeader
	}

	req, err := a.client.prepareRequest(r.ctx, requestPath, http.MethodGet, requestBody, headerParams, queryParams, formParams, requestFormFieldName, requestFileName, requestFileBytes)
	if err != nil {
		return returnValue, err
	}

	res, err := a.client.callAPI(req)
	if err != nil || res == nil {
		return returnValue, err
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	res.Body = ioutil.NopCloser(bytes.NewBuffer(responseBody))
	if err != nil {
		return returnValue, err
	}

	responseContentType := res.Header.Get("Content-Type")

	if res.StatusCode >= 300 {
		var v APIError
		if err := a.client.decode(&v, responseBody, responseContentType); err != nil {
			return returnValue, APIError{
				Status:  res.StatusCode,
				Message: "failed to deserialise error response",
				parent:  err,
			}
		}
		v.Status = res.StatusCode
		return returnValue, v
	}

	if err := a.client.decode(&returnValue, responseBody, responseContentType); err != nil {
		return returnValue, APIError{
			Status:  res.StatusCode,
			Message: "failed to deserialise response body",
			parent:  err,
		}
	}

	return returnValue, nil
}

// APIListUptimeMonitoringLocationsRequest represents a request type.
type APIListUptimeMonitoringLocationsRequest struct {
	ctx        context.Context
	APIService LocationsAPI
	best       *bool
	location   *string
}

// Best sets best on the request type.
func (r APIListUptimeMonitoringLocationsRequest) Best(best bool) APIListUptimeMonitoringLocationsRequest {
	r.best = &best
	return r
}

// Location sets location on the request type.
func (r APIListUptimeMonitoringLocationsRequest) Location(location string) APIListUptimeMonitoringLocationsRequest {
	r.location = &location
	return r
}

// Execute executes the request.
func (r APIListUptimeMonitoringLocationsRequest) Execute() (MonitoringLocations, error) {
	return r.APIService.ListUptimeMonitoringLocationsExecute(r)
}

// ListUptimeMonitoringLocations Get all uptime monitoring locations.
func (a *LocationsService) ListUptimeMonitoringLocations(ctx context.Context) APIListUptimeMonitoringLocationsRequest {
	return APIListUptimeMonitoringLocationsRequest{
		ctx:        ctx,
		APIService: a,
	}
}

// ListUptimeMonitoringLocationsWithData Get all uptime monitoring locations.
// The use of this method is discouraged as it does not provide the level of
// type safety afforded by the field methods on the request type.
func (a *LocationsService) ListUptimeMonitoringLocationsWithData(ctx context.Context, m map[string]interface{}) APIListUptimeMonitoringLocationsRequest {
	r := a.ListUptimeMonitoringLocations(ctx)
	return r
}

// Execute executes the request.
func (a *LocationsService) ListUptimeMonitoringLocationsExecute(r APIListUptimeMonitoringLocationsRequest) (MonitoringLocations, error) {
	var (
		requestBody          interface{}
		requestFormFieldName string
		requestFileName      string
		requestFileBytes     []byte
		returnValue          MonitoringLocations
	)

	basePath, err := a.client.ServerURLWithContext(r.ctx, "LocationsService.ListUptimeMonitoringLocations")
	if err != nil {
		return returnValue, err
	}

	requestPath := basePath + "/uptime-locations"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if r.best != nil {
		queryParams.Add("best", parameterToString(*r.best))
	}
	if r.location != nil {
		queryParams.Add("location", parameterToString(*r.location))
	}
	// Determine the Content-Type header.
	contentTypes := []string{}

	// Set Content-Type header
	requestContentTypeHeader := selectHeaderContentType(contentTypes)
	if requestContentTypeHeader != "" {
		headerParams["Content-Type"] = requestContentTypeHeader
	}

	// Determine the Accept header.
	accepts := []string{"application/json"}

	// Set Accept header.
	requestAcceptHeader := selectHeaderAccept(accepts)
	if requestAcceptHeader != "" {
		headerParams["Accept"] = requestAcceptHeader
	}

	req, err := a.client.prepareRequest(r.ctx, requestPath, http.MethodGet, requestBody, headerParams, queryParams, formParams, requestFormFieldName, requestFileName, requestFileBytes)
	if err != nil {
		return returnValue, err
	}

	res, err := a.client.callAPI(req)
	if err != nil || res == nil {
		return returnValue, err
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	res.Body = ioutil.NopCloser(bytes.NewBuffer(responseBody))
	if err != nil {
		return returnValue, err
	}

	responseContentType := res.Header.Get("Content-Type")

	if res.StatusCode >= 300 {
		var v APIError
		if err := a.client.decode(&v, responseBody, responseContentType); err != nil {
			return returnValue, APIError{
				Status:  res.StatusCode,
				Message: "failed to deserialise error response",
				parent:  err,
			}
		}
		v.Status = res.StatusCode
		return returnValue, v
	}

	if err := a.client.decode(&returnValue, responseBody, responseContentType); err != nil {
		return returnValue, APIError{
			Status:  res.StatusCode,
			Message: "failed to deserialise response body",
			parent:  err,
		}
	}

	return returnValue, nil
}
