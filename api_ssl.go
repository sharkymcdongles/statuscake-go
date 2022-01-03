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
 * API version: 1.0.0-beta.1
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
	"strings"
)

// Linger please
var _ context.Context

// SslAPI describes the necessary methods to adhere to this interface.
type SslAPI interface {
	CreateSslTest(ctx context.Context) APICreateSslTestRequest
	CreateSslTestWithData(ctx context.Context, m map[string]interface{}) APICreateSslTestRequest
	CreateSslTestExecute(r APICreateSslTestRequest) (APIResponse, error)
	DeleteSslTest(ctx context.Context, testId string) APIDeleteSslTestRequest
	DeleteSslTestExecute(r APIDeleteSslTestRequest) error
	GetSslTest(ctx context.Context, testId string) APIGetSslTestRequest
	GetSslTestExecute(r APIGetSslTestRequest) (SSLTestResponse, error)
	ListSslTests(ctx context.Context) APIListSslTestsRequest
	ListSslTestsExecute(r APIListSslTestsRequest) (SSLTests, error)
	UpdateSslTest(ctx context.Context, testId string) APIUpdateSslTestRequest
	UpdateSslTestWithData(ctx context.Context, testId string, m map[string]interface{}) APIUpdateSslTestRequest
	UpdateSslTestExecute(r APIUpdateSslTestRequest) error
}

// SslService Ssl service.
type SslService service

// APICreateSslTestRequest represents a request type.
type APICreateSslTestRequest struct {
	ctx              context.Context
	APIService       SslAPI
	websiteUrl       *string
	checkRate        *SSLTestCheckRate
	alertAt          *[]int32
	alertAtCsv       *string
	alertBroken      *bool
	alertExpiry      *bool
	alertMixed       *bool
	alertReminder    *bool
	contactGroups    *[]string
	contactGroupsCsv *string
	followRedirects  *bool
	hostname         *string
	paused           *bool
	userAgent        *string
}

// WebsiteURL sets websiteUrl on the request type.
func (r APICreateSslTestRequest) WebsiteURL(websiteUrl string) APICreateSslTestRequest {
	r.websiteUrl = &websiteUrl
	return r
}

// CheckRate sets checkRate on the request type.
func (r APICreateSslTestRequest) CheckRate(checkRate SSLTestCheckRate) APICreateSslTestRequest {
	r.checkRate = &checkRate
	return r
}

// AlertAt sets alertAt on the request type.
func (r APICreateSslTestRequest) AlertAt(alertAt []int32) APICreateSslTestRequest {
	r.alertAt = &alertAt
	return r
}

// AlertAtCsv sets alertAtCsv on the request type.
func (r APICreateSslTestRequest) AlertAtCsv(alertAtCsv string) APICreateSslTestRequest {
	r.alertAtCsv = &alertAtCsv
	return r
}

// AlertBroken sets alertBroken on the request type.
func (r APICreateSslTestRequest) AlertBroken(alertBroken bool) APICreateSslTestRequest {
	r.alertBroken = &alertBroken
	return r
}

// AlertExpiry sets alertExpiry on the request type.
func (r APICreateSslTestRequest) AlertExpiry(alertExpiry bool) APICreateSslTestRequest {
	r.alertExpiry = &alertExpiry
	return r
}

// AlertMixed sets alertMixed on the request type.
func (r APICreateSslTestRequest) AlertMixed(alertMixed bool) APICreateSslTestRequest {
	r.alertMixed = &alertMixed
	return r
}

// AlertReminder sets alertReminder on the request type.
func (r APICreateSslTestRequest) AlertReminder(alertReminder bool) APICreateSslTestRequest {
	r.alertReminder = &alertReminder
	return r
}

// ContactGroups sets contactGroups on the request type.
func (r APICreateSslTestRequest) ContactGroups(contactGroups []string) APICreateSslTestRequest {
	r.contactGroups = &contactGroups
	return r
}

// ContactGroupsCsv sets contactGroupsCsv on the request type.
func (r APICreateSslTestRequest) ContactGroupsCsv(contactGroupsCsv string) APICreateSslTestRequest {
	r.contactGroupsCsv = &contactGroupsCsv
	return r
}

// FollowRedirects sets followRedirects on the request type.
func (r APICreateSslTestRequest) FollowRedirects(followRedirects bool) APICreateSslTestRequest {
	r.followRedirects = &followRedirects
	return r
}

// Hostname sets hostname on the request type.
func (r APICreateSslTestRequest) Hostname(hostname string) APICreateSslTestRequest {
	r.hostname = &hostname
	return r
}

// Paused sets paused on the request type.
func (r APICreateSslTestRequest) Paused(paused bool) APICreateSslTestRequest {
	r.paused = &paused
	return r
}

// UserAgent sets userAgent on the request type.
func (r APICreateSslTestRequest) UserAgent(userAgent string) APICreateSslTestRequest {
	r.userAgent = &userAgent
	return r
}

// Execute executes the request.
func (r APICreateSslTestRequest) Execute() (APIResponse, error) {
	return r.APIService.CreateSslTestExecute(r)
}

// CreateSslTest Create an SSL check.
func (a *SslService) CreateSslTest(ctx context.Context) APICreateSslTestRequest {
	return APICreateSslTestRequest{
		ctx:        ctx,
		APIService: a,
	}
}

// CreateSslTestWithData Create an SSL check.
// The use of this method is discouraged as it does not provide the level of
// type safety afforded by the field methods on the request type.
func (a *SslService) CreateSslTestWithData(ctx context.Context, m map[string]interface{}) APICreateSslTestRequest {
	r := a.CreateSslTest(ctx)

	if prop, ok := m["website_url"].(string); ok {
		r.websiteUrl = &prop
	}

	if prop, ok := m["check_rate"].(SSLTestCheckRate); ok {
		r.checkRate = &prop
	}

	if prop, ok := m["alert_at"].([]int32); ok {
		r.alertAt = &prop
	}

	if prop, ok := m["alert_at_csv"].(string); ok {
		r.alertAtCsv = &prop
	}

	if prop, ok := m["alert_broken"].(bool); ok {
		r.alertBroken = &prop
	}

	if prop, ok := m["alert_expiry"].(bool); ok {
		r.alertExpiry = &prop
	}

	if prop, ok := m["alert_mixed"].(bool); ok {
		r.alertMixed = &prop
	}

	if prop, ok := m["alert_reminder"].(bool); ok {
		r.alertReminder = &prop
	}

	if prop, ok := m["contact_groups"].([]string); ok {
		r.contactGroups = &prop
	}

	if prop, ok := m["contact_groups_csv"].(string); ok {
		r.contactGroupsCsv = &prop
	}

	if prop, ok := m["follow_redirects"].(bool); ok {
		r.followRedirects = &prop
	}

	if prop, ok := m["hostname"].(string); ok {
		r.hostname = &prop
	}

	if prop, ok := m["paused"].(bool); ok {
		r.paused = &prop
	}

	if prop, ok := m["user_agent"].(string); ok {
		r.userAgent = &prop
	}

	return r
}

// Execute executes the request.
func (a *SslService) CreateSslTestExecute(r APICreateSslTestRequest) (APIResponse, error) {
	var (
		requestBody          interface{}
		requestFormFieldName string
		requestFileName      string
		requestFileBytes     []byte
		returnValue          APIResponse
	)

	basePath, err := a.client.ServerURLWithContext(r.ctx, "SslService.CreateSslTest")
	if err != nil {
		return returnValue, err
	}

	requestPath := basePath + "/ssl"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if r.websiteUrl == nil {
		return returnValue, errorf("websiteUrl is required and must be specified")
	}

	if r.checkRate == nil {
		return returnValue, errorf("checkRate is required and must be specified")
	}

	// Determine the Content-Type header.
	contentTypes := []string{"application/x-www-form-urlencoded"}

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

	formParams.Add("website_url", parameterToString(*r.websiteUrl))
	formParams.Add("check_rate", parameterToString(*r.checkRate))

	if r.alertAt != nil {
		// Explicity empty array. This indictes the consumer intended to pass an
		// empty value and therefore likely want to nullify the field.
		if len(*r.alertAt) == 0 {
			formParams.Add("alert_at[]", "")
		}
		for _, val := range *r.alertAt {
			formParams.Add("alert_at[]", parameterToString(val))
		}
	}

	if r.alertAtCsv != nil {
		formParams.Add("alert_at_csv", parameterToString(*r.alertAtCsv))
	}

	if r.alertBroken != nil {
		formParams.Add("alert_broken", parameterToString(*r.alertBroken))
	}

	if r.alertExpiry != nil {
		formParams.Add("alert_expiry", parameterToString(*r.alertExpiry))
	}

	if r.alertMixed != nil {
		formParams.Add("alert_mixed", parameterToString(*r.alertMixed))
	}

	if r.alertReminder != nil {
		formParams.Add("alert_reminder", parameterToString(*r.alertReminder))
	}

	if r.contactGroups != nil {
		// Explicity empty array. This indictes the consumer intended to pass an
		// empty value and therefore likely want to nullify the field.
		if len(*r.contactGroups) == 0 {
			formParams.Add("contact_groups[]", "")
		}
		for _, val := range *r.contactGroups {
			formParams.Add("contact_groups[]", parameterToString(val))
		}
	}

	if r.contactGroupsCsv != nil {
		formParams.Add("contact_groups_csv", parameterToString(*r.contactGroupsCsv))
	}

	if r.followRedirects != nil {
		formParams.Add("follow_redirects", parameterToString(*r.followRedirects))
	}

	if r.hostname != nil {
		formParams.Add("hostname", parameterToString(*r.hostname))
	}

	if r.paused != nil {
		formParams.Add("paused", parameterToString(*r.paused))
	}

	if r.userAgent != nil {
		formParams.Add("user_agent", parameterToString(*r.userAgent))
	}
	req, err := a.client.prepareRequest(r.ctx, requestPath, http.MethodPost, requestBody, headerParams, queryParams, formParams, requestFormFieldName, requestFileName, requestFileBytes)
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

// APIDeleteSslTestRequest represents a request type.
type APIDeleteSslTestRequest struct {
	ctx        context.Context
	APIService SslAPI
	testId     string
}

// Execute executes the request.
func (r APIDeleteSslTestRequest) Execute() error {
	return r.APIService.DeleteSslTestExecute(r)
}

// DeleteSslTest Delete an SSL check.
func (a *SslService) DeleteSslTest(ctx context.Context, testId string) APIDeleteSslTestRequest {
	return APIDeleteSslTestRequest{
		ctx:        ctx,
		APIService: a,
		testId:     testId,
	}
}

// DeleteSslTestWithData Delete an SSL check.
// The use of this method is discouraged as it does not provide the level of
// type safety afforded by the field methods on the request type.
func (a *SslService) DeleteSslTestWithData(ctx context.Context, testId string, m map[string]interface{}) APIDeleteSslTestRequest {
	r := a.DeleteSslTest(ctx, testId)
	return r
}

// Execute executes the request.
func (a *SslService) DeleteSslTestExecute(r APIDeleteSslTestRequest) error {
	var (
		requestBody          interface{}
		requestFormFieldName string
		requestFileName      string
		requestFileBytes     []byte
	)

	basePath, err := a.client.ServerURLWithContext(r.ctx, "SslService.DeleteSslTest")
	if err != nil {
		return err
	}

	requestPath := basePath + "/ssl/{test_id}"
	requestPath = strings.Replace(requestPath, "{"+"test_id"+"}", url.PathEscape(parameterToString(r.testId)), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

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

	req, err := a.client.prepareRequest(r.ctx, requestPath, http.MethodDelete, requestBody, headerParams, queryParams, formParams, requestFormFieldName, requestFileName, requestFileBytes)
	if err != nil {
		return err
	}

	res, err := a.client.callAPI(req)
	if err != nil || res == nil {
		return err
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	res.Body = ioutil.NopCloser(bytes.NewBuffer(responseBody))
	if err != nil {
		return err
	}

	responseContentType := res.Header.Get("Content-Type")

	if res.StatusCode >= 300 {
		var v APIError
		if err := a.client.decode(&v, responseBody, responseContentType); err != nil {
			return APIError{
				Status:  res.StatusCode,
				Message: "failed to deserialise error response",
				parent:  err,
			}
		}
		v.Status = res.StatusCode
		return v
	}

	return nil
}

// APIGetSslTestRequest represents a request type.
type APIGetSslTestRequest struct {
	ctx        context.Context
	APIService SslAPI
	testId     string
}

// Execute executes the request.
func (r APIGetSslTestRequest) Execute() (SSLTestResponse, error) {
	return r.APIService.GetSslTestExecute(r)
}

// GetSslTest Retrieve an SSL check.
func (a *SslService) GetSslTest(ctx context.Context, testId string) APIGetSslTestRequest {
	return APIGetSslTestRequest{
		ctx:        ctx,
		APIService: a,
		testId:     testId,
	}
}

// GetSslTestWithData Retrieve an SSL check.
// The use of this method is discouraged as it does not provide the level of
// type safety afforded by the field methods on the request type.
func (a *SslService) GetSslTestWithData(ctx context.Context, testId string, m map[string]interface{}) APIGetSslTestRequest {
	r := a.GetSslTest(ctx, testId)
	return r
}

// Execute executes the request.
func (a *SslService) GetSslTestExecute(r APIGetSslTestRequest) (SSLTestResponse, error) {
	var (
		requestBody          interface{}
		requestFormFieldName string
		requestFileName      string
		requestFileBytes     []byte
		returnValue          SSLTestResponse
	)

	basePath, err := a.client.ServerURLWithContext(r.ctx, "SslService.GetSslTest")
	if err != nil {
		return returnValue, err
	}

	requestPath := basePath + "/ssl/{test_id}"
	requestPath = strings.Replace(requestPath, "{"+"test_id"+"}", url.PathEscape(parameterToString(r.testId)), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

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

// APIListSslTestsRequest represents a request type.
type APIListSslTestsRequest struct {
	ctx        context.Context
	APIService SslAPI
}

// Execute executes the request.
func (r APIListSslTestsRequest) Execute() (SSLTests, error) {
	return r.APIService.ListSslTestsExecute(r)
}

// ListSslTests Get all SSL checks.
func (a *SslService) ListSslTests(ctx context.Context) APIListSslTestsRequest {
	return APIListSslTestsRequest{
		ctx:        ctx,
		APIService: a,
	}
}

// ListSslTestsWithData Get all SSL checks.
// The use of this method is discouraged as it does not provide the level of
// type safety afforded by the field methods on the request type.
func (a *SslService) ListSslTestsWithData(ctx context.Context, m map[string]interface{}) APIListSslTestsRequest {
	r := a.ListSslTests(ctx)
	return r
}

// Execute executes the request.
func (a *SslService) ListSslTestsExecute(r APIListSslTestsRequest) (SSLTests, error) {
	var (
		requestBody          interface{}
		requestFormFieldName string
		requestFileName      string
		requestFileBytes     []byte
		returnValue          SSLTests
	)

	basePath, err := a.client.ServerURLWithContext(r.ctx, "SslService.ListSslTests")
	if err != nil {
		return returnValue, err
	}

	requestPath := basePath + "/ssl"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

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

// APIUpdateSslTestRequest represents a request type.
type APIUpdateSslTestRequest struct {
	ctx              context.Context
	APIService       SslAPI
	testId           string
	checkRate        *SSLTestCheckRate
	alertAt          *[]int32
	alertAtCsv       *string
	alertBroken      *bool
	alertExpiry      *bool
	alertMixed       *bool
	alertReminder    *bool
	contactGroups    *[]string
	contactGroupsCsv *string
	followRedirects  *bool
	hostname         *string
	paused           *bool
	userAgent        *string
}

// CheckRate sets checkRate on the request type.
func (r APIUpdateSslTestRequest) CheckRate(checkRate SSLTestCheckRate) APIUpdateSslTestRequest {
	r.checkRate = &checkRate
	return r
}

// AlertAt sets alertAt on the request type.
func (r APIUpdateSslTestRequest) AlertAt(alertAt []int32) APIUpdateSslTestRequest {
	r.alertAt = &alertAt
	return r
}

// AlertAtCsv sets alertAtCsv on the request type.
func (r APIUpdateSslTestRequest) AlertAtCsv(alertAtCsv string) APIUpdateSslTestRequest {
	r.alertAtCsv = &alertAtCsv
	return r
}

// AlertBroken sets alertBroken on the request type.
func (r APIUpdateSslTestRequest) AlertBroken(alertBroken bool) APIUpdateSslTestRequest {
	r.alertBroken = &alertBroken
	return r
}

// AlertExpiry sets alertExpiry on the request type.
func (r APIUpdateSslTestRequest) AlertExpiry(alertExpiry bool) APIUpdateSslTestRequest {
	r.alertExpiry = &alertExpiry
	return r
}

// AlertMixed sets alertMixed on the request type.
func (r APIUpdateSslTestRequest) AlertMixed(alertMixed bool) APIUpdateSslTestRequest {
	r.alertMixed = &alertMixed
	return r
}

// AlertReminder sets alertReminder on the request type.
func (r APIUpdateSslTestRequest) AlertReminder(alertReminder bool) APIUpdateSslTestRequest {
	r.alertReminder = &alertReminder
	return r
}

// ContactGroups sets contactGroups on the request type.
func (r APIUpdateSslTestRequest) ContactGroups(contactGroups []string) APIUpdateSslTestRequest {
	r.contactGroups = &contactGroups
	return r
}

// ContactGroupsCsv sets contactGroupsCsv on the request type.
func (r APIUpdateSslTestRequest) ContactGroupsCsv(contactGroupsCsv string) APIUpdateSslTestRequest {
	r.contactGroupsCsv = &contactGroupsCsv
	return r
}

// FollowRedirects sets followRedirects on the request type.
func (r APIUpdateSslTestRequest) FollowRedirects(followRedirects bool) APIUpdateSslTestRequest {
	r.followRedirects = &followRedirects
	return r
}

// Hostname sets hostname on the request type.
func (r APIUpdateSslTestRequest) Hostname(hostname string) APIUpdateSslTestRequest {
	r.hostname = &hostname
	return r
}

// Paused sets paused on the request type.
func (r APIUpdateSslTestRequest) Paused(paused bool) APIUpdateSslTestRequest {
	r.paused = &paused
	return r
}

// UserAgent sets userAgent on the request type.
func (r APIUpdateSslTestRequest) UserAgent(userAgent string) APIUpdateSslTestRequest {
	r.userAgent = &userAgent
	return r
}

// Execute executes the request.
func (r APIUpdateSslTestRequest) Execute() error {
	return r.APIService.UpdateSslTestExecute(r)
}

// UpdateSslTest Update an SSL check.
func (a *SslService) UpdateSslTest(ctx context.Context, testId string) APIUpdateSslTestRequest {
	return APIUpdateSslTestRequest{
		ctx:        ctx,
		APIService: a,
		testId:     testId,
	}
}

// UpdateSslTestWithData Update an SSL check.
// The use of this method is discouraged as it does not provide the level of
// type safety afforded by the field methods on the request type.
func (a *SslService) UpdateSslTestWithData(ctx context.Context, testId string, m map[string]interface{}) APIUpdateSslTestRequest {
	r := a.UpdateSslTest(ctx, testId)

	if prop, ok := m["check_rate"].(SSLTestCheckRate); ok {
		r.checkRate = &prop
	}

	if prop, ok := m["alert_at"].([]int32); ok {
		r.alertAt = &prop
	}

	if prop, ok := m["alert_at_csv"].(string); ok {
		r.alertAtCsv = &prop
	}

	if prop, ok := m["alert_broken"].(bool); ok {
		r.alertBroken = &prop
	}

	if prop, ok := m["alert_expiry"].(bool); ok {
		r.alertExpiry = &prop
	}

	if prop, ok := m["alert_mixed"].(bool); ok {
		r.alertMixed = &prop
	}

	if prop, ok := m["alert_reminder"].(bool); ok {
		r.alertReminder = &prop
	}

	if prop, ok := m["contact_groups"].([]string); ok {
		r.contactGroups = &prop
	}

	if prop, ok := m["contact_groups_csv"].(string); ok {
		r.contactGroupsCsv = &prop
	}

	if prop, ok := m["follow_redirects"].(bool); ok {
		r.followRedirects = &prop
	}

	if prop, ok := m["hostname"].(string); ok {
		r.hostname = &prop
	}

	if prop, ok := m["paused"].(bool); ok {
		r.paused = &prop
	}

	if prop, ok := m["user_agent"].(string); ok {
		r.userAgent = &prop
	}

	return r
}

// Execute executes the request.
func (a *SslService) UpdateSslTestExecute(r APIUpdateSslTestRequest) error {
	var (
		requestBody          interface{}
		requestFormFieldName string
		requestFileName      string
		requestFileBytes     []byte
	)

	basePath, err := a.client.ServerURLWithContext(r.ctx, "SslService.UpdateSslTest")
	if err != nil {
		return err
	}

	requestPath := basePath + "/ssl/{test_id}"
	requestPath = strings.Replace(requestPath, "{"+"test_id"+"}", url.PathEscape(parameterToString(r.testId)), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// Determine the Content-Type header.
	contentTypes := []string{"application/x-www-form-urlencoded"}

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

	if r.checkRate != nil {
		formParams.Add("check_rate", parameterToString(*r.checkRate))
	}

	if r.alertAt != nil {
		// Explicity empty array. This indictes the consumer intended to pass an
		// empty value and therefore likely want to nullify the field.
		if len(*r.alertAt) == 0 {
			formParams.Add("alert_at[]", "")
		}
		for _, val := range *r.alertAt {
			formParams.Add("alert_at[]", parameterToString(val))
		}
	}

	if r.alertAtCsv != nil {
		formParams.Add("alert_at_csv", parameterToString(*r.alertAtCsv))
	}

	if r.alertBroken != nil {
		formParams.Add("alert_broken", parameterToString(*r.alertBroken))
	}

	if r.alertExpiry != nil {
		formParams.Add("alert_expiry", parameterToString(*r.alertExpiry))
	}

	if r.alertMixed != nil {
		formParams.Add("alert_mixed", parameterToString(*r.alertMixed))
	}

	if r.alertReminder != nil {
		formParams.Add("alert_reminder", parameterToString(*r.alertReminder))
	}

	if r.contactGroups != nil {
		// Explicity empty array. This indictes the consumer intended to pass an
		// empty value and therefore likely want to nullify the field.
		if len(*r.contactGroups) == 0 {
			formParams.Add("contact_groups[]", "")
		}
		for _, val := range *r.contactGroups {
			formParams.Add("contact_groups[]", parameterToString(val))
		}
	}

	if r.contactGroupsCsv != nil {
		formParams.Add("contact_groups_csv", parameterToString(*r.contactGroupsCsv))
	}

	if r.followRedirects != nil {
		formParams.Add("follow_redirects", parameterToString(*r.followRedirects))
	}

	if r.hostname != nil {
		formParams.Add("hostname", parameterToString(*r.hostname))
	}

	if r.paused != nil {
		formParams.Add("paused", parameterToString(*r.paused))
	}

	if r.userAgent != nil {
		formParams.Add("user_agent", parameterToString(*r.userAgent))
	}
	req, err := a.client.prepareRequest(r.ctx, requestPath, http.MethodPut, requestBody, headerParams, queryParams, formParams, requestFormFieldName, requestFileName, requestFileBytes)
	if err != nil {
		return err
	}

	res, err := a.client.callAPI(req)
	if err != nil || res == nil {
		return err
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	res.Body = ioutil.NopCloser(bytes.NewBuffer(responseBody))
	if err != nil {
		return err
	}

	responseContentType := res.Header.Get("Content-Type")

	if res.StatusCode >= 300 {
		var v APIError
		if err := a.client.decode(&v, responseBody, responseContentType); err != nil {
			return APIError{
				Status:  res.StatusCode,
				Message: "failed to deserialise error response",
				parent:  err,
			}
		}
		v.Status = res.StatusCode
		return v
	}

	return nil
}
