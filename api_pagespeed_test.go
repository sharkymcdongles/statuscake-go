//go:build consumer
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
 * API version: 1.0.0-beta.3
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

func TestCreatePagespeedTest(t *testing.T) {
	t.Run("returns a created status on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing contact group",
				Parameters: map[string]interface{}{
					"group_id": 123,
				},
			}).
			UponReceiving("A request to create a valid pagespeed test").
			WithRequest(http.MethodPost, S("/v1/pagespeed")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
				"Content-Type":  []Matcher{S("application/x-www-form-urlencoded")},
			}).
			WithBody("application/x-www-form-urlencoded", []byte(
				"alert_bigger=200&"+
					"alert_slower=300&"+
					"alert_smaller=20&"+
					"check_rate=3600&"+
					"contact_groups%5B%5D=123&"+
					"name=statuscake.com&"+
					"paused=true&"+
					"region=UK&"+
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
			res, _ := c.CreatePagespeedTest(context.Background()).
				Name("statuscake.com").
				WebsiteURL("https://www.statuscake.com").
				CheckRate(statuscake.PagespeedTestCheckRateOneHour).
				AlertBigger(200).
				AlertSlower(300).
				AlertSmaller(20).
				ContactGroups([]string{
					"123",
				}).
				Paused(true).
				Region(statuscake.PagespeedTestRegionUnitedKingdom).
				Execute()

			return equal(res.Data.NewID, "1")
		})
	})

	t.Run("returns an error if the request fails", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to create an invalid pagespeed test").
			WithRequest(http.MethodPost, S("/v1/pagespeed")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
				"Content-Type":  []Matcher{S("application/x-www-form-urlencoded")},
			}).
			WithBody("application/x-www-form-urlencoded", []byte(
				"check_rate=3600&"+
					"name=statuscake.com&"+
					"region=UK&"+
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
			_, err := c.CreatePagespeedTest(context.Background()).
				Name("statuscake.com").
				WebsiteURL("this,is,not,valid").
				CheckRate(statuscake.PagespeedTestCheckRateOneHour).
				Region(statuscake.PagespeedTestRegionUnitedKingdom).
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

func TestDeletePagespeedTest(t *testing.T) {
	t.Run("returns a no content status on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing pagespeed test",
			}).
			UponReceiving("A request to delete a pagespeed test").
			WithRequest(http.MethodDelete, FromProviderState("/v1/pagespeed/${id}", "/v1/pagespeed/1")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusNoContent)

		executeTest(t, func(c *statuscake.Client) error {
			return c.DeletePagespeedTest(context.Background(), "1").Execute()
		})
	})

	t.Run("returns an error when the pagespeed test does not exist", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to delete a pagespeed test").
			WithRequest(http.MethodDelete, S("/v1/pagespeed/2")).
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
			err := c.DeletePagespeedTest(context.Background(), "2").Execute()
			return equal(err, statuscake.APIError{
				Status:  http.StatusNotFound,
				Message: "No results found",
				Errors:  map[string][]string{},
			})
		})
	})
}

func TestGetPagespeedTest(t *testing.T) {
	t.Run("returns a pagespeed test on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing pagespeed test and contact group",
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
					"id":             FromProviderState("${id}", "1"),
					"name":           Like("statuscake.com"),
					"website_url":    Like("https://www.statuscake.com"),
					"check_rate":     Integer(3600),
					"alert_bigger":   Integer(200),
					"alert_slower":   Integer(300),
					"alert_smaller":  Integer(20),
					"contact_groups": EachLike("123", 1),
					"latest_stats": matchers.StructMatcher{
						"requests":     Integer(27),
						"has_issue":    Like(true),
						"latest_issue": Like("The Total Load Time of the Page (20216/ms) is larger than the alert threshold of 300/ms"),
					},
					"location": Like("PAGESPD-UK1"),
					"paused":   Like(true),
				},
			})

		executeTest(t, func(c *statuscake.Client) error {
			test, _ := c.GetPagespeedTest(context.Background(), "1").Execute()
			return equal(test.Data, statuscake.PagespeedTest{
				ID:           "1",
				Name:         "statuscake.com",
				WebsiteURL:   "https://www.statuscake.com",
				CheckRate:    statuscake.PagespeedTestCheckRateOneHour,
				AlertBigger:  200,
				AlertSlower:  300,
				AlertSmaller: 20,
				ContactGroups: []string{
					"123",
				},
				LatestStats: &statuscake.PagespeedTestStats{
					Requests:    27,
					HasIssue:    true,
					LatestIssue: statuscake.PtrString("The Total Load Time of the Page (20216/ms) is larger than the alert threshold of 300/ms"),
				},
				Location: "PAGESPD-UK1",
				Paused:   true,
			})
		})
	})

	t.Run("returns an error when the pagespeed test does not exist", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to get a pagespeed test").
			WithRequest(http.MethodGet, S("/v1/pagespeed/2")).
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
			_, err := c.GetPagespeedTest(context.Background(), "2").Execute()
			return equal(err, statuscake.APIError{
				Status:  http.StatusNotFound,
				Message: "No results found",
				Errors:  map[string][]string{},
			})
		})
	})
}

func TestListPagespeedTests(t *testing.T) {
	t.Run("returns a list of pagespeed tests on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "Existing pagespeed tests and contact group",
			}).
			UponReceiving("A request to get a list of pagespeed tests").
			WithRequest(http.MethodGet, S("/v1/pagespeed")).
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
						"website_url":    Like("https://www.statuscake.com"),
						"check_rate":     Integer(3600),
						"alert_bigger":   Integer(200),
						"alert_slower":   Integer(300),
						"alert_smaller":  Integer(20),
						"contact_groups": EachLike("123", 1),
						"latest_stats": matchers.StructMatcher{
							"requests":     Integer(27),
							"has_issue":    Like(true),
							"latest_issue": Like("The Total Load Time of the Page (20216/ms) is larger than the alert threshold of 300/ms"),
						},
						"location": Like("PAGESPD-UK1"),
						"paused":   Like(true),
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
			tests, _ := c.ListPagespeedTests(context.Background()).Execute()
			return equal(tests.Data, []statuscake.PagespeedTest{
				statuscake.PagespeedTest{
					ID:           "1",
					Name:         "statuscake.com",
					WebsiteURL:   "https://www.statuscake.com",
					CheckRate:    statuscake.PagespeedTestCheckRateOneHour,
					AlertBigger:  200,
					AlertSlower:  300,
					AlertSmaller: 20,
					ContactGroups: []string{
						"123",
					},
					LatestStats: &statuscake.PagespeedTestStats{
						Requests:    27,
						HasIssue:    true,
						LatestIssue: statuscake.PtrString("The Total Load Time of the Page (20216/ms) is larger than the alert threshold of 300/ms"),
					},
					Location: "PAGESPD-UK1",
					Paused:   true,
				},
			})
		})
	})

	t.Run("returns an empty list when there are no pagespeed tests", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to get a list of pagespeed tests").
			WithRequest(http.MethodGet, S("/v1/pagespeed")).
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
			tests, _ := c.ListPagespeedTests(context.Background()).Execute()
			return equal(tests.Data, []statuscake.PagespeedTest{})
		})
	})
}

func TestListPagespeedTestHistory(t *testing.T) {
	t.Run("returns a list of pagespeed test history results on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing pagespeed test with history results",
			}).
			UponReceiving("A request to get a list of pagespeed test history results").
			WithRequest(http.MethodGet, FromProviderState("/v1/pagespeed/${id}/history", "/v1/pagespeed/1/history")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusOK).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"data": EachLike(
					matchers.StructMatcher{
						"created_at":   Timestamp(),
						"loadtime":     Integer(1490),
						"filesize":     Decimal(598.384),
						"har_location": Like("https://16a0fd6b5b5bece1d29a-7aa19249e604542958e6a694f67d0bbf.ssl.cf5.rackcdn.com/53a6b075-3b93-4752-b707-93b08fe5ae44.json"),
						"requests":     Integer(4),
						"throttling":   Like("NONE"),
					}, 1,
				),
				"links": matchers.StructMatcher{
					"self": FromProviderState(
						"https://api.statuscake.com/v1/pagespeed/${id}/history?limit=25&before=949411800",
						"https://api.statuscake.com/v1/pagespeed/1/history?limit=25&before=949411800",
					),
				},
				"metadata": matchers.StructMatcher{
					"aggregates": matchers.StructMatcher{
						"filesize": matchers.StructMatcher{
							"min": Decimal(0),
							"max": Decimal(598.384),
							"avg": Decimal(598.384),
						},
						"loadtime": matchers.StructMatcher{
							"min": Integer(0),
							"max": Integer(1490),
							"avg": Decimal(1490),
						},
						"requests": matchers.StructMatcher{
							"min": Integer(0),
							"max": Integer(4),
							"avg": Decimal(4),
						},
					},
				},
			})

		executeTest(t, func(c *statuscake.Client) error {
			results, _ := c.ListPagespeedTestHistory(context.Background(), "1").Execute()

			return equal(results.Data, []statuscake.PagespeedTestHistoryResult{
				statuscake.PagespeedTestHistoryResult{
					Created:     time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
					Filesize:    598.384,
					HARLocation: "https://16a0fd6b5b5bece1d29a-7aa19249e604542958e6a694f67d0bbf.ssl.cf5.rackcdn.com/53a6b075-3b93-4752-b707-93b08fe5ae44.json",
					Loadtime:    1490,
					Requests:    4,
					Throttling:  statuscake.PagespeedTestThrottlingNone,
				},
			})
		})
	})

	t.Run("returns an error when the pagespeed test does not exist", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to get a list of pagespeed test history results").
			WithRequest(http.MethodGet, S("/v1/pagespeed/2/history")).
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
			_, err := c.ListPagespeedTestHistory(context.Background(), "2").Execute()

			return equal(err, statuscake.APIError{
				Status:  http.StatusNotFound,
				Message: "No results found",
				Errors:  map[string][]string{},
			})
		})
	})

	t.Run("returns an empty result set when there are no pagespeed test history results", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing pagespeed test",
			}).
			UponReceiving("A request to get a list of pagespeed test history results").
			WithRequest(http.MethodGet, FromProviderState("/v1/pagespeed/${id}/history", "/v1/pagespeed/1/history")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusOK).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"data": Like([]interface{}{}),
				"links": matchers.StructMatcher{
					"self": FromProviderState(
						"https://api.statuscake.com/v1/pagespeed/${id}/history?limit=25&before=949411800",
						"https://api.statuscake.com/v1/pagespeed/1/history?limit=25&before=949411800",
					),
				},
			})

		executeTest(t, func(c *statuscake.Client) error {
			results, _ := c.ListPagespeedTestHistory(context.Background(), "1").Execute()

			return equal(results.Data, []statuscake.PagespeedTestHistoryResult{})
		})
	})
}

func TestUpdatePagespeedTest(t *testing.T) {
	t.Run("returns a no content status on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing pagespeed test",
			}).
			UponReceiving("A request to update a pagespeed test").
			WithRequest(http.MethodPut, FromProviderState("/v1/pagespeed/${id}", "/v1/pagespeed/1")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
				"Content-Type":  []Matcher{S("application/x-www-form-urlencoded")},
			}).
			WithBody("application/x-www-form-urlencoded", []byte(
				"alert_bigger=10&"+
					"alert_slower=100&"+
					"alert_smaller=1&"+
					"check_rate=1800&"+
					"contact_groups%5B%5D=&"+
					"name=example.com&"+
					"paused=false",
			)).
			WillRespondWith(http.StatusNoContent)

		executeTest(t, func(c *statuscake.Client) error {
			return c.UpdatePagespeedTest(context.Background(), "1").
				Name("example.com").
				CheckRate(statuscake.PagespeedTestCheckRateThirtyMinutes).
				AlertBigger(10).
				AlertSlower(100).
				AlertSmaller(1).
				ContactGroups([]string{}).
				Paused(false).
				Execute()
		})
	})

	t.Run("returns an error if the request fails", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "An existing pagespeed test",
			}).
			UponReceiving("A request to update an invalid pagespeed test").
			WithRequest(http.MethodPut, FromProviderState("/v1/pagespeed/${id}", "/v1/pagespeed/1")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
				"Content-Type":  []Matcher{S("application/x-www-form-urlencoded")},
			}).
			WithBody("application/x-www-form-urlencoded", []byte(
				"region=DE",
			)).
			WillRespondWith(http.StatusBadRequest).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"message": Like("No server available for ISO 'DE'. Please select a different ISO or contact support."),
				"errors":  matchers.StructMatcher{},
			})

		executeTest(t, func(c *statuscake.Client) error {
			err := c.UpdatePagespeedTest(context.Background(), "1").
				Region(statuscake.PagespeedTestRegionGermany).
				Execute()

			return equal(err, statuscake.APIError{
				Status:  http.StatusBadRequest,
				Message: "No server available for ISO 'DE'. Please select a different ISO or contact support.",
				Errors:  map[string][]string{},
			})
		})
	})

	t.Run("returns an error when the pagespeed test does not exist", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to update a pagespeed test").
			WithRequest(http.MethodPut, S("/v1/pagespeed/2")).
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
			err := c.UpdatePagespeedTest(context.Background(), "2").
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
