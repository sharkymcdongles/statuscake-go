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
 * API version: 1.0.0-beta.2
 * Contact: support@statuscake.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package statuscake_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/pact-foundation/pact-go/v2/matchers"
	. "github.com/pact-foundation/pact-go/v2/sugar"

	"github.com/StatusCakeDev/statuscake-go"
)

func TestListPagespeedMonitoringLocations(t *testing.T) {
	t.Run("returns a list of monitoring locations on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "Existing pagespeed monitoring locations",
			}).
			UponReceiving("A request to get a list of pagespeed monitoring locations").
			WithRequest(http.MethodGet, S("/v1/pagespeed-locations")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusOK).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"data": EachLike(
					matchers.StructMatcher{
						"description": Like("Google Chrome 72.0.3626.121"),
						"ipv4":        IPAddress(),
						"ipv6":        IPv6Address(),
						"region":      Like("United Kingdom"),
						"region_code": Like("United Kingdom"),
						"status":      Status(),
					}, 1,
				),
			})

		executeTest(t, func(c *statuscake.Client) error {
			locations, _ := c.ListPagespeedMonitoringLocations(context.Background()).Execute()
			return equal(locations.Data, []statuscake.MonitoringLocation{
				statuscake.MonitoringLocation{
					Description: "Google Chrome 72.0.3626.121",
					IPv4:        statuscake.PtrString("127.0.0.1"),
					IPv6:        statuscake.PtrString("::ffff:192.0.2.128"),
					Region:      "United Kingdom",
					RegionCode:  "United Kingdom",
					Status:      statuscake.MonitoringLocationStatusUp,
				},
			})
		})
	})

	t.Run("returns an empty list when there are no monitoring locations", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to get a list of pagespeed monitoring locations").
			WithRequest(http.MethodGet, S("/v1/pagespeed-locations")).
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
			locations, _ := c.ListPagespeedMonitoringLocations(context.Background()).Execute()
			return equal(locations.Data, []statuscake.MonitoringLocation{})
		})
	})
}

func TestListUptimeMonitoringLocations(t *testing.T) {
	t.Run("returns a list of monitoring locations on success", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			Given(ProviderStateV3{
				Name: "Existing uptime monitoring locations and associated regions",
			}).
			UponReceiving("A request to get a list of uptime monitoring locations").
			WithRequest(http.MethodGet, S("/v1/uptime-locations")).
			WithHeaders(matchers.HeadersMatcher{
				"Accept":        []Matcher{S("application/json")},
				"Authorization": []Matcher{S("Bearer 123456789")},
			}).
			WillRespondWith(http.StatusOK).
			WithHeader("Content-Type", S("application/json")).
			WithJSONBody(Map{
				"data": EachLike(
					matchers.StructMatcher{
						"description": Like("United Kingdom, London - 5"),
						"ipv4":        IPAddress(),
						"ipv6":        IPv6Address(),
						"region":      Like("United Kingdom / London"),
						"region_code": Like("london"),
						"status":      Status(),
					}, 1,
				),
			})

		executeTest(t, func(c *statuscake.Client) error {
			locations, _ := c.ListUptimeMonitoringLocations(context.Background()).Execute()
			return equal(locations.Data, []statuscake.MonitoringLocation{
				statuscake.MonitoringLocation{
					Description: "United Kingdom, London - 5",
					IPv4:        statuscake.PtrString("127.0.0.1"),
					IPv6:        statuscake.PtrString("::ffff:192.0.2.128"),
					Region:      "United Kingdom / London",
					RegionCode:  "london",
					Status:      statuscake.MonitoringLocationStatusUp,
				},
			})
		})
	})

	t.Run("returns an empty list when there are no monitoring locations", func(t *testing.T) {
		mockProvider.
			AddInteraction().
			UponReceiving("A request to get a list of uptime monitoring locations").
			WithRequest(http.MethodGet, S("/v1/uptime-locations")).
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
			locations, _ := c.ListUptimeMonitoringLocations(context.Background()).Execute()
			return equal(locations.Data, []statuscake.MonitoringLocation{})
		})
	})
}
