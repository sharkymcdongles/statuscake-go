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

package statuscake_test

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/StatusCakeDev/statuscake-go"
)

func TestCreateContactGroup(t *testing.T) {
	t.Run("returns a created status on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expectEqual(t, mustParse(t, r), url.Values{
				"email_addresses[]": []string{"johnsmith@example.com", "janesmith@example.com"},
				"integrations[]":    []string{"1", "2", "3"},
				"mobile_numbers[]":  []string{"447712345678", "447987462344"},
				"name":              []string{"Operations Team"},
				"ping_url":          []string{"https://ping.example.com"},
			})

			w.WriteHeader(http.StatusCreated)
			w.Write(mustRead(t, "testdata/create-resource-success.json"))
		}))
		defer s.Close()

		res, err := c.CreateContactGroup(context.Background()).
			Name("Operations Team").
			PingURL("https://ping.example.com").
			EmailAddresses([]string{
				"johnsmith@example.com",
				"janesmith@example.com",
			}).
			MobileNumbers([]string{
				"447712345678",
				"447987462344",
			}).
			Integrations([]string{
				"1",
				"2",
				"3",
			}).
			Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, res.Data.NewID, "2")
	})

	t.Run("returns an error if the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(mustRead(t, "testdata/invalid-ping-url-error.json"))
		}))
		defer s.Close()

		_, err := c.CreateContactGroup(context.Background()).
			Name("Operations Team").
			PingURL("this,is,not,valid").
			EmailAddresses([]string{
				"johnsmith@example.com",
				"janesmith@example.com",
			}).
			MobileNumbers([]string{
				"+447712345678",
				"+447987462344",
			}).
			Integrations([]string{
				"1",
				"2",
				"3",
			}).
			Execute()
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectEqual(t, err, statuscake.APIError{
			Status:  http.StatusBadRequest,
			Message: "The provided parameters are invalid. Check the errors output for details information.",
			Errors: map[string][]string{
				"ping_url": []string{"Ping Url is not a valid URL"},
			},
		})
	})
}

func TestDeleteContactGroup(t *testing.T) {
	t.Run("returns a no content status on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
		}))
		defer s.Close()

		err := c.DeleteContactGroup(context.Background(), "2").Execute()
		if err != nil {
			t.Fatal(err.Error())
		}
	})

	t.Run("returns an error when the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(mustRead(t, "testdata/fetch-resource-error.json"))
		}))
		defer s.Close()

		err := c.DeleteContactGroup(context.Background(), "3").Execute()
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectEqual(t, err, statuscake.APIError{
			Status:  http.StatusNotFound,
			Message: "No results found",
			Errors:  map[string][]string{},
		})
	})
}

func TestGetContactGroup(t *testing.T) {
	t.Run("returns a contact group on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(mustRead(t, "testdata/get-contact-group-success.json"))
		}))
		defer s.Close()

		test, err := c.GetContactGroup(context.Background(), "2").Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, test.Data, statuscake.ContactGroup{
			ID:   "2",
			Name: "Marketing Team",
			EmailAddresses: []string{
				"johnappleseed@example.com",
				"janeappleseed@example.com",
			},
			MobileNumbers: []string{
				"447891998195",
				"447112887498",
			},
			Integrations: []string{
				"4",
				"5",
				"6",
			},
		})
	})

	t.Run("returns an error when the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(mustRead(t, "testdata/fetch-resource-error.json"))
		}))
		defer s.Close()

		_, err := c.GetContactGroup(context.Background(), "3").Execute()
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectEqual(t, err, statuscake.APIError{
			Status:  http.StatusNotFound,
			Message: "No results found",
			Errors:  map[string][]string{},
		})
	})
}

func TestListContactGroups(t *testing.T) {
	t.Run("returns a list of contact groups on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(mustRead(t, "testdata/list-contact-groups-success.json"))
		}))
		defer s.Close()

		groups, err := c.ListContactGroups(context.Background()).Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, groups.Data, []statuscake.ContactGroup{
			statuscake.ContactGroup{
				ID:      "1",
				Name:    "Operations Team",
				PingURL: statuscake.PtrString("https://ping.example.com"),
				EmailAddresses: []string{
					"johnsmith@example.com",
					"janesmith@example.com",
				},
				MobileNumbers: []string{
					"447712345678",
					"447987462344",
				},
				Integrations: []string{
					"1",
					"2",
					"3",
				},
			},
			statuscake.ContactGroup{
				ID:   "2",
				Name: "Marketing Team",
				EmailAddresses: []string{
					"johnappleseed@example.com",
					"janeappleseed@example.com",
				},
				MobileNumbers: []string{
					"447891998195",
					"447112887498",
				},
				Integrations: []string{
					"4",
					"5",
					"6",
				},
			},
		})
	})

	t.Run("returns an error when the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(mustRead(t, "testdata/fetch-resource-error.json"))
		}))
		defer s.Close()

		_, err := c.ListContactGroups(context.Background()).Execute()
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectEqual(t, err, statuscake.APIError{
			Status:  http.StatusNotFound,
			Message: "No results found",
			Errors:  map[string][]string{},
		})
	})
}

func TestUpdateContactGroup(t *testing.T) {
	t.Run("returns a no content status on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expectEqual(t, mustParse(t, r), url.Values{
				"email_addresses[]": []string{""},
				"integrations[]":    []string{"4", "5", "6"},
				"mobile_numbers[]":  []string{"447891998195"},
				"name":              []string{"Development Team"},
				"ping_url":          []string{"https://ping.example.com/groups"},
			})

			w.WriteHeader(http.StatusNoContent)
		}))
		defer s.Close()

		err := c.UpdateContactGroup(context.Background(), "2").
			Name("Development Team").
			PingURL("https://ping.example.com/groups").
			EmailAddresses([]string{}).
			MobileNumbers([]string{
				"447891998195",
			}).
			Integrations([]string{
				"4",
				"5",
				"6",
			}).
			Execute()
		if err != nil {
			t.Fatal(err.Error())
		}
	})

	t.Run("returns an error if the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(mustRead(t, "testdata/fetch-resource-error.json"))
		}))
		defer s.Close()

		err := c.UpdateContactGroup(context.Background(), "3").
			Name("Development Team").
			Execute()
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectEqual(t, err, statuscake.APIError{
			Status:  http.StatusNotFound,
			Message: "No results found",
			Errors:  map[string][]string{},
		})
	})
}
