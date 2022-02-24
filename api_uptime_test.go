// +build consumer

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
 * API version: 1.0.0-beta.2
 * Contact: support@statuscake.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package statuscake_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/pact-foundation/pact-go/v2/matchers"
	. "github.com/pact-foundation/pact-go/v2/sugar"

	"github.com/StatusCakeDev/statuscake-go"
)

func TestCreateUptimeTest(t *testing.T) {
	t.Run("returns a created status on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing contact group",
				Parameters: map[string]interface{}{
					"group_id": 123,
				},
			}).
			UponReceiving("A request to create a valid uptime test").
			WithRequest(http.MethodPost, S("/v1/uptime")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
				"Content-Type":  []Matcher{S("application/x-www-form-urlencoded")},
			}).
			WithBody("application/x-www-form-urlencoded", []byte(
				"basic_password=password&"+
					"basic_username=username&"+
					"check_rate=3600&"+
					"confirmation=3&"+
					"contact_groups%5B%5D=123&"+
					"custom_header=%7B%22Authorization%22%3A+%22Bearer+abc123%22%7D&"+
					"do_not_find=true&"+
					"enable_ssl_alert=true&"+
					"final_endpoint=https%3A%2F%2Fwww.statuscake.com%2Fredirected&"+
					"find_string=Hello%2C+world&"+
					"follow_redirects=true&"+
					"host=AWS&"+
					"include_header=true&"+
					"name=statuscake.com&"+
					"paused=true&"+
					"port=123&"+
					"post_body=%7B%22key%22%3A+%22value%22%7D&"+
					"post_raw=key%3Dvalue&"+
					"regions%5B%5D=london&"+
					"status_codes_csv=200%2C201&"+
					"tags%5B%5D=testing&"+
					"test_type=HTTP&"+
					"timeout=10&"+
					"trigger_rate=2&"+
					"use_jar=true&"+
					"user_agent=Mozilla%2F5.0+%28Macintosh%3B+Intel+Mac+OS+X+11_1%29+AppleWebKit%2F537.36+%28KHTML%2C+like+Gecko%29+Chrome%2F87.0.4280.141+Safari%2F537.36+OPR%2F73.0.3856.344&"+
					"website_url=https%3A%2F%2Fwww.statuscake.com",
			)).
			WillRespondWith(http.StatusCreated).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"data": matchers.StructMatcher{
					"new_id": Like("1"),
				},
			})

		executeTest(t, func(c *statuscake.Client) error {
			res, _ := c.CreateUptimeTest(context.Background()).
				Name("statuscake.com").
				TestType(statuscake.UptimeTestTypeHTTP).
				WebsiteURL("https://www.statuscake.com").
				CheckRate(statuscake.UptimeTestCheckRateOneHour).
				BasicPassword("password").
				BasicUsername("username").
				Confirmation(3).
				ContactGroups([]string{
					"123",
				}).
				CustomHeader(`{"Authorization": "Bearer abc123"}`).
				DoNotFind(true).
				EnableSSLAlert(true).
				FinalEndpoint("https://www.statuscake.com/redirected").
				FindString("Hello, world").
				FollowRedirects(true).
				Host("AWS").
				IncludeHeader(true).
				Paused(true).
				Port(123).
				PostBody(`{"key": "value"}`).
				PostRaw("key=value").
				Regions([]string{"london"}).
				StatusCodes([]string{
					"200",
					"201",
				}).
				Tags([]string{
					"testing",
				}).
				Timeout(10).
				TriggerRate(2).
				UseJAR(true).
				UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36 OPR/73.0.3856.344").
				Execute()

			return equal(res.Data.NewID, "1")
		})
	})

	t.Run("returns an error if the request fails", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to create an invalid uptime test").
			WithRequest(http.MethodPost, S("/v1/uptime")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
				"Content-Type":  []Matcher{S("application/x-www-form-urlencoded")},
			}).
			WithBody("application/x-www-form-urlencoded", []byte(
				"check_rate=3600&"+
					"name=statuscake.com&"+
					"test_type=HTTP&"+
					"website_url=this%2Cis%2Cnot%2Cvalid",
			)).
			WillRespondWith(http.StatusBadRequest).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"message": Like("The provided parameters are invalid. Check the errors output for detailed information."),
				"errors": matchers.StructMatcher{
					"website_url": EachLike("Website Url is not a valid URL", 1),
				},
			})

		executeTest(t, func(c *statuscake.Client) error {
			_, err := c.CreateUptimeTest(context.Background()).
				Name("statuscake.com").
				TestType(statuscake.UptimeTestTypeHTTP).
				WebsiteURL("this,is,not,valid").
				CheckRate(statuscake.UptimeTestCheckRateOneHour).
				Execute()

			return equal(err, statuscake.APIError{
				Status:  http.StatusBadRequest,
				Message: "The provided parameters are invalid. Check the errors output for detailed information.",
				Errors: map[string][]string{
					"website_url": []string{"Website Url is not a valid URL"},
				},
			})
		})
	})
}

func TestDeleteUptimeTest(t *testing.T) {
	t.Run("returns a no content status on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing uptime test",
			}).
			UponReceiving("A request to delete an uptime test").
			WithRequest(http.MethodDelete, FromProviderState("/v1/uptime/${id}", "/v1/uptime/1")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusNoContent)

		executeTest(t, func(c *statuscake.Client) error {
			return c.DeleteUptimeTest(context.Background(), "1").Execute()
		})
	})

	t.Run("returns an error when the uptime test does not exist", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to delete an uptime test").
			WithRequest(http.MethodDelete, S("/v1/uptime/2")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusNotFound).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"message": Like("No results found"),
				"errors":  matchers.StructMatcher{},
			})

		executeTest(t, func(c *statuscake.Client) error {
			err := c.DeleteUptimeTest(context.Background(), "2").Execute()
			return equal(err, statuscake.APIError{
				Status:  http.StatusNotFound,
				Message: "No results found",
				Errors:  map[string][]string{},
			})
		})
	})
}

func TestGetUptimeTest(t *testing.T) {
	t.Run("returns an uptime test on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing uptime test and contact group",
			}).
			UponReceiving("A request to get a pagespeed test").
			WithRequest(http.MethodGet, FromProviderState("/v1/pagespeed/${id}", "/v1/pagespeed/1")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusOK).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"data": matchers.StructMatcher{
					"id":               FromProviderState("${id}", "1"),
					"name":             Like("statuscake.com"),
					"test_type":        Like("HTTP"),
					"website_url":      Like("https://www.statuscake.com"),
					"check_rate":       Integer(3600),
					"confirmation":     Integer(3),
					"contact_groups":   EachLike("123", 1),
					"custom_header":    Like(`{"Authorization": "Bearer abc123"}`),
					"do_not_find":      Like(true),
					"enable_ssl_alert": Like(true),
					"final_endpoint":   Like("https://www.statuscake.com/redirected"),
					"find_string":      Like("Hello, world"),
					"follow_redirects": Like(true),
					"host":             Like("AWS"),
					"include_header":   Like(true),
					"last_tested_at":   Timestamp(),
					"paused":           Like(true),
					"port":             Integer(123),
					"post_body":        Like(`{"key": "value"}`),
					"post_raw":         Like("key=value"),
					"processing":       Like(false),
					"servers": EachLike(
						matchers.StructMatcher{
							"description": Like("United Kingdom, London - 5"),
							"ipv4":        IPAddress(),
							"ipv6":        IPv6Address(),
							"region":      Like("United Kingdom / London"),
							"region_code": Like("london"),
							"status":      Status(),
						}, 1,
					),
					"status":       Status(),
					"status_codes": EachLike("200", 1),
					"tags":         EachLike("testing", 1),
					"timeout":      Integer(10),
					"trigger_rate": Integer(2),
					"uptime":       Decimal(100),
					"use_jar":      Like(true),
					"user_agent":   Like("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.102 Safari/537.3"),
				},
			})

		executeTest(t, func(c *statuscake.Client) error {
			test, _ := c.GetUptimeTest(context.Background(), "1").Execute()
			return equal(test.Data, statuscake.UptimeTest{
				ID:           "1",
				Name:         "statuscake.com",
				TestType:     statuscake.UptimeTestTypeHTTP,
				WebsiteURL:   "https://www.statuscake.com",
				CheckRate:    statuscake.UptimeTestCheckRateOneHour,
				Confirmation: 3,
				ContactGroups: []string{
					"123",
				},
				CustomHeader:    statuscake.PtrString(`{"Authorization": "Bearer abc123"}`),
				DoNotFind:       true,
				EnableSSLAlert:  true,
				FinalEndpoint:   statuscake.PtrString("https://www.statuscake.com/redirected"),
				FindString:      statuscake.PtrString("Hello, world"),
				FollowRedirects: true,
				Host:            statuscake.PtrString("AWS"),
				IncludeHeader:   true,
				LastTested:      statuscake.PtrTime(time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)),
				Paused:          true,
				Port:            statuscake.PtrInt32(123),
				PostBody:        statuscake.PtrString(`{"key": "value"}`),
				PostRaw:         statuscake.PtrString("key=value"),
				Servers: []statuscake.MonitoringLocation{
					statuscake.MonitoringLocation{
						Description: "United Kingdom, London - 5",
						IPv4:        statuscake.PtrString("127.0.0.1"),
						IPv6:        statuscake.PtrString("::ffff:192.0.2.128"),
						Region:      "United Kingdom / London",
						RegionCode:  "london",
						Status:      statuscake.MonitoringLocationStatusUp,
					},
				},
				Status: statuscake.UptimeTestStatusUp,
				StatusCodes: []string{
					"200",
				},
				Tags: []string{
					"testing",
				},
				Timeout:     10,
				TriggerRate: 2,
				Uptime:      100,
				UseJAR:      true,
				UserAgent:   statuscake.PtrString("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.102 Safari/537.3"),
			})
		})
	})

	t.Run("returns an error when the uptime test does not exist", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to get an uptime test").
			WithRequest(http.MethodGet, S("/v1/uptime/2")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusNotFound).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"message": Like("No results found"),
				"errors":  matchers.StructMatcher{},
			})

		executeTest(t, func(c *statuscake.Client) error {
			_, err := c.GetUptimeTest(context.Background(), "2").Execute()
			return equal(err, statuscake.APIError{
				Status:  http.StatusNotFound,
				Message: "No results found",
				Errors:  map[string][]string{},
			})
		})
	})
}

func TestListUptimeTests(t *testing.T) {
	t.Run("returns a list of uptime tests on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "Existing uptime tests and contact group",
			}).
			UponReceiving("A request to get a list of uptime tests").
			WithRequest(http.MethodGet, S("/v1/uptime")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusOK).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"data": EachLike(
					matchers.StructMatcher{
						"id":             FromProviderState("${id}", "1"),
						"name":           Like("statuscake.com"),
						"test_type":      Like("HTTP"),
						"website_url":    Like("https://www.statuscake.com"),
						"check_rate":     Integer(3600),
						"contact_groups": EachLike("123", 1),
						"paused":         Like(true),
						"status":         Status(),
						"tags":           EachLike("testing", 1),
						"uptime":         Decimal(100),
					}, 1,
				),
				"metadata": matchers.StructMatcher{
					"page":        Like(1),
					"per_page":    Like(25),
					"page_count":  Like(1),
					"total_count": Like(5),
				},
			})

		executeTest(t, func(c *statuscake.Client) error {
			tests, _ := c.ListUptimeTests(context.Background()).Execute()
			return equal(tests.Data, []statuscake.UptimeTestOverview{
				statuscake.UptimeTestOverview{
					ID:         "1",
					Name:       "statuscake.com",
					TestType:   statuscake.UptimeTestTypeHTTP,
					WebsiteURL: "https://www.statuscake.com",
					CheckRate:  statuscake.UptimeTestCheckRateOneHour,
					ContactGroups: []string{
						"123",
					},
					Paused: true,
					Status: statuscake.UptimeTestStatusUp,
					Tags: []string{
						"testing",
					},
					Uptime: statuscake.PtrFloat32(100),
				},
			})
		})
	})

	t.Run("returns an empty list when there are no uptime tests", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to get a list of uptime tests").
			WithRequest(http.MethodGet, S("/v1/uptime")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusOK).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"data": Like([]interface{}{}),
				"metadata": matchers.StructMatcher{
					"page":        Like(1),
					"per_page":    Like(25),
					"page_count":  Like(1),
					"total_count": 0,
				},
			})

		executeTest(t, func(c *statuscake.Client) error {
			tests, _ := c.ListUptimeTests(context.Background()).Execute()
			return equal(tests.Data, []statuscake.UptimeTestOverview{})
		})
	})
}

func TestListUptimeTestHistory(t *testing.T) {
	t.Run("returns a list of uptime test history results on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing uptime test with history results",
			}).
			UponReceiving("A request to get a list of uptime test history results").
			WithRequest(http.MethodGet, FromProviderState("/v1/uptime/${id}/history", "/v1/uptime/1/history")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusOK).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"data": EachLike(
					matchers.StructMatcher{
						"created_at":  Timestamp(),
						"location":    Like("BR1"),
						"performance": Integer(259),
						"status_code": Integer(200),
					}, 1,
				),
			})

		executeTest(t, func(c *statuscake.Client) error {
			results, _ := c.ListUptimeTestHistory(context.Background(), "1").Execute()
			return equal(results.Data, []statuscake.UptimeTestHistoryResult{
				statuscake.UptimeTestHistoryResult{
					Created:     time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
					Location:    statuscake.PtrString("BR1"),
					Performance: statuscake.PtrInt64(259),
					StatusCode:  statuscake.PtrInt32(200),
				},
			})
		})
	})

	t.Run("returns an error when the uptime test does not exist", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to get a list of uptime test history results").
			WithRequest(http.MethodGet, S("/v1/uptime/2/history")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusNotFound).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"message": Like("No results found"),
				"errors":  matchers.StructMatcher{},
			})

		executeTest(t, func(c *statuscake.Client) error {
			_, err := c.ListUptimeTestHistory(context.Background(), "2").Execute()
			return equal(err, statuscake.APIError{
				Status:  http.StatusNotFound,
				Message: "No results found",
				Errors:  map[string][]string{},
			})
		})
	})

	t.Run("returns an empty list when there are no uptime test history results", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing uptime test",
			}).
			UponReceiving("A request to get a list of uptime test history results").
			WithRequest(http.MethodGet, FromProviderState("/v1/uptime/${id}/history", "/v1/uptime/1/history")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusOK).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"data": Like([]interface{}{}),
			})

		executeTest(t, func(c *statuscake.Client) error {
			results, _ := c.ListUptimeTestHistory(context.Background(), "1").Execute()
			return equal(results.Data, []statuscake.UptimeTestHistoryResult{})
		})
	})
}

func TestListUptimeTestPeriods(t *testing.T) {
	t.Run("returns a list of uptime test periods on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing uptime test with test periods",
			}).
			UponReceiving("A request to get a list of uptime test periods").
			WithRequest(http.MethodGet, FromProviderState("/v1/uptime/${id}/periods", "/v1/uptime/1/periods")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusOK).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"data": EachLike(
					matchers.StructMatcher{
						"created_at": Timestamp(),
						"duration":   Integer(189000),
						"ended_at":   Timestamp(),
						"status":     Status(),
					}, 1,
				),
			})

		executeTest(t, func(c *statuscake.Client) error {
			periods, _ := c.ListUptimeTestPeriods(context.Background(), "1").Execute()
			return equal(periods.Data, []statuscake.UptimeTestPeriod{
				statuscake.UptimeTestPeriod{
					Created:  time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
					Duration: statuscake.PtrInt64(189000),
					Ended:    statuscake.PtrTime(time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)),
					Status:   "up",
				},
			})
		})
	})

	t.Run("returns an error when the uptime test does not exist", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to get a list of uptime test periods").
			WithRequest(http.MethodGet, S("/v1/uptime/2/periods")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusNotFound).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"message": Like("No results found"),
				"errors":  matchers.StructMatcher{},
			})

		executeTest(t, func(c *statuscake.Client) error {
			_, err := c.ListUptimeTestPeriods(context.Background(), "2").Execute()
			return equal(err, statuscake.APIError{
				Status:  http.StatusNotFound,
				Message: "No results found",
				Errors:  map[string][]string{},
			})
		})
	})

	t.Run("returns an empty list when there are no uptime test periods", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing uptime test",
			}).
			UponReceiving("A request to get a list of uptime test periods").
			WithRequest(http.MethodGet, FromProviderState("/v1/uptime/${id}/periods", "/v1/uptime/1/periods")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusOK).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"data": Like([]interface{}{}),
			})

		executeTest(t, func(c *statuscake.Client) error {
			periods, _ := c.ListUptimeTestPeriods(context.Background(), "1").Execute()
			return equal(periods.Data, []statuscake.UptimeTestPeriod{})
		})
	})
}

func TestListUptimeTestAlerts(t *testing.T) {
	t.Run("returns a list of uptime test alerts on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing uptime test with alerts",
			}).
			UponReceiving("A request to get a list of uptime test alerts").
			WithRequest(http.MethodGet, FromProviderState("/v1/uptime/${id}/alerts", "/v1/uptime/1/alerts")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusOK).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"data": EachLike(
					matchers.StructMatcher{
						"id":           FromProviderState("${id}", "1"),
						"status":       Status(),
						"status_code":  Integer(404),
						"triggered_at": Timestamp(),
					}, 1,
				),
			})

		executeTest(t, func(c *statuscake.Client) error {
			alerts, _ := c.ListUptimeTestAlerts(context.Background(), "1").Execute()
			return equal(alerts.Data, []statuscake.UptimeTestAlert{
				statuscake.UptimeTestAlert{
					ID:         "1",
					Status:     "up",
					StatusCode: 404,
					Triggered:  statuscake.PtrTime(time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)),
				},
			})
		})
	})

	t.Run("returns an error when the uptime test does not exist", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to get a list of uptime test alerts").
			WithRequest(http.MethodGet, S("/v1/uptime/2/alerts")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusNotFound).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"message": Like("No results found"),
				"errors":  matchers.StructMatcher{},
			})

		executeTest(t, func(c *statuscake.Client) error {
			_, err := c.ListUptimeTestAlerts(context.Background(), "2").Execute()

			return equal(err, statuscake.APIError{
				Status:  http.StatusNotFound,
				Message: "No results found",
				Errors:  map[string][]string{},
			})
		})
	})

	t.Run("returns an empty list when there are no uptime test alerts", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing uptime test",
			}).
			UponReceiving("A request to get a list of uptime test alerts").
			WithRequest(http.MethodGet, FromProviderState("/v1/uptime/${id}/alerts", "/v1/uptime/1/alerts")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusOK).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"data": Like([]interface{}{}),
			})

		executeTest(t, func(c *statuscake.Client) error {
			alerts, _ := c.ListUptimeTestAlerts(context.Background(), "1").Execute()
			return equal(alerts.Data, []statuscake.UptimeTestAlert{})
		})
	})
}

func TestUpdateUptimeTest(t *testing.T) {
	t.Run("returns a no content status on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing uptime test",
			}).
			UponReceiving("A request to update a uptime test").
			WithRequest(http.MethodPut, FromProviderState("/v1/uptime/${id}", "/v1/uptime/1")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
				"Content-Type":  []Matcher{S("application/x-www-form-urlencoded")},
			}).
			WithBody("application/x-www-form-urlencoded", []byte(
				"check_rate=1800&"+
					"contact_groups%5B%5D=&"+
					"do_not_find=false&"+
					"enable_ssl_alert=false&"+
					"final_endpoint=&"+
					"follow_redirects=false&"+
					"name=example.com&"+
					"paused=false&"+
					"regions%5B%5D=london&"+
					"regions%5B%5D=paris&"+
					"status_codes_csv=100%2C200%2C400&"+
					"tags%5B%5D=example&"+
					"use_jar=false",
			)).
			WillRespondWith(http.StatusNoContent)

		executeTest(t, func(c *statuscake.Client) error {
			return c.UpdateUptimeTest(context.Background(), "1").
				Name("example.com").
				CheckRate(statuscake.UptimeTestCheckRateThirtyMinutes).
				ContactGroups([]string{}).
				DoNotFind(false).
				EnableSSLAlert(false).
				FinalEndpoint("").
				FollowRedirects(false).
				Paused(false).
				Regions([]string{
					"london",
					"paris",
				}).
				StatusCodes([]string{
					"100",
					"200",
					"400",
				}).
				Tags([]string{
					"example",
				}).
				UseJAR(false).
				Execute()
		})
	})

	t.Run("returns an error if the request fails", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing uptime test",
			}).
			UponReceiving("A request to update an invalid uptime test").
			WithRequest(http.MethodPut, FromProviderState("/v1/uptime/${id}", "/v1/uptime/1")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
				"Content-Type":  []Matcher{S("application/x-www-form-urlencoded")},
			}).
			WithBody("application/x-www-form-urlencoded", []byte(
				"regions%5B%5D=tatooine",
			)).
			WillRespondWith(http.StatusBadRequest).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"message": Like("invalid region"),
				"errors":  matchers.StructMatcher{},
			})

		executeTest(t, func(c *statuscake.Client) error {
			err := c.UpdateUptimeTest(context.Background(), "1").
				Regions([]string{
					"tatooine",
				}).
				Execute()

			return equal(err, statuscake.APIError{
				Status:  http.StatusBadRequest,
				Message: "invalid region",
				Errors:  map[string][]string{},
			})
		})
	})

	t.Run("returns an error when the uptime test does not exist", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to update a uptime test").
			WithRequest(http.MethodPut, S("/v1/uptime/2")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
				"Content-Type":  []Matcher{S("application/x-www-form-urlencoded")},
			}).
			WithBody("application/x-www-form-urlencoded", []byte(
				"paused=false",
			)).
			WillRespondWith(http.StatusNotFound).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"message": Like("No results found"),
				"errors":  matchers.StructMatcher{},
			})

		executeTest(t, func(c *statuscake.Client) error {
			err := c.UpdateUptimeTest(context.Background(), "2").
				Paused(false).
				Execute()

			return equal(err, statuscake.APIError{
				Status:  http.StatusNotFound,
				Message: "No results found",
				Errors:  map[string][]string{},
			})
		})
	})
}
