package statuscake_test

import (
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/StatusCakeDev/statuscake-go"
)

func TestCreateSSLTest(t *testing.T) {
	t.Run("returns a created status on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expectEqual(t, mustParse(t, r), url.Values{
				"alert_at_csv":       []string{"1,7,30"},
				"alert_broken":       []string{"true"},
				"alert_expiry":       []string{"true"},
				"alert_mixed":        []string{"true"},
				"alert_reminder":     []string{"true"},
				"check_rate":         []string{"3600"},
				"contact_groups_csv": []string{"123"},
				"follow_redirects":   []string{"true"},
				"hostname":           []string{"AWS"},
				"paused":             []string{"true"},
				"user_agent":         []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36 OPR/73.0.3856.344"},
				"website_url":        []string{"https://www.statuscake.com"},
			})

			w.WriteHeader(http.StatusCreated)
			w.Write(mustRead(t, "testdata/create-resource-success.json"))
		}))
		defer s.Close()

		res, err := c.CreateSslTest(context.Background()).
			Paused(true).
			WebsiteURL("https://www.statuscake.com").
			CheckRate(statuscake.SSLTestCheckRateOneHour).
			ContactGroups([]string{
				"123",
			}).
			AlertAt([]string{"1", "7", "30"}).
			AlertBroken(true).
			AlertExpiry(true).
			AlertMixed(true).
			AlertReminder(true).
			FollowRedirects(true).
			Hostname("AWS").
			UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36 OPR/73.0.3856.344").
			Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, res.Data.NewID, "2")
	})

	t.Run("returns an error if the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(mustRead(t, "testdata/invalid-website-url-error.json"))
		}))
		defer s.Close()

		_, err := c.CreateSslTest(context.Background()).
			Paused(true).
			WebsiteURL("this,is,not,valid").
			CheckRate(statuscake.SSLTestCheckRateOneHour).
			ContactGroups([]string{
				"123",
			}).
			AlertAt([]string{"1", "7", "30"}).
			AlertBroken(true).
			AlertExpiry(true).
			AlertMixed(true).
			AlertReminder(true).
			FollowRedirects(true).
			Execute()
		if err == nil {
			t.Fatal("expected error, got nil")
		}

		expectEqual(t, err, statuscake.APIError{
			Status:  http.StatusBadRequest,
			Message: "The provided parameters are invalid. Check the errors output for details information.",
			Errors: map[string][]string{
				"website_url": []string{"Website Url is not a valid URL"},
			},
		})
	})
}

func TestDeleteSSLTest(t *testing.T) {
	t.Run("returns a no content status on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
		}))
		defer s.Close()

		err := c.DeleteSslTest(context.Background(), "2").Execute()
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

		err := c.DeleteSslTest(context.Background(), "3").Execute()
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

func TestGetSSLTest(t *testing.T) {
	t.Run("returns a ssl test on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(mustRead(t, "testdata/get-ssl-test-success.json"))
		}))
		defer s.Close()

		test, err := c.GetSslTest(context.Background(), "2").Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, test.Data, statuscake.SSLTest{
			ID:         "2",
			Paused:     true,
			WebsiteURL: "https://www.statuscake.com",
			CheckRate:  statuscake.SSLTestCheckRateOneHour,
			ContactGroups: []string{
				"123",
			},
			AlertAt:       []int32{1, 7, 30},
			AlertBroken:   true,
			AlertExpiry:   true,
			AlertMixed:    true,
			AlertReminder: true,
			Flags: &statuscake.SSLTestFlags{
				FollowRedirects: true,
			},
			FollowRedirects: true,
			MixedContent:    []statuscake.SSLTestMixedContent{},
			Hostname:        statuscake.PtrString("svc.example.com"),
			LastReminder:    statuscake.PtrInt32(0),
			Updated:         statuscake.PtrTime(time.Date(2021, 1, 6, 12, 7, 35, 0, time.UTC)),
			UserAgent:       statuscake.PtrString("Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36 OPR/73.0.3856.344"),
		})
	})

	t.Run("returns an error when the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(mustRead(t, "testdata/fetch-resource-error.json"))
		}))
		defer s.Close()

		_, err := c.GetSslTest(context.Background(), "3").Execute()
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

func TestListSSLTests(t *testing.T) {
	t.Run("returns a list of ssl tests on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(mustRead(t, "testdata/list-ssl-tests-success.json"))
		}))
		defer s.Close()

		tests, err := c.ListSslTests(context.Background()).Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, tests.Data, []statuscake.SSLTest{
			statuscake.SSLTest{
				ID:                "1",
				Paused:            false,
				WebsiteURL:        "https://www.statuscake.com",
				CheckRate:         statuscake.SSLTestCheckRateThirtyMinutes,
				ContactGroups:     []string{},
				AlertAt:           []int32{1, 7, 30},
				AlertBroken:       true,
				AlertExpiry:       true,
				AlertMixed:        true,
				AlertReminder:     true,
				CertificateScore:  statuscake.PtrInt32(95),
				CertificateStatus: statuscake.PtrString("CERT_OK"),
				Cipher:            statuscake.PtrString("TLS_CHACHA20_POLY1305_SHA256"),
				CipherScore:       statuscake.PtrInt32(100),
				Flags: &statuscake.SSLTestFlags{
					HasMixed: true,
					HasPFS:   true,
				},
				Hostname:         statuscake.PtrString("svc.example.com"),
				IssuerCommonName: statuscake.PtrString("Sectigo RSA Domain Validation Secure Server CA"),
				MixedContent: []statuscake.SSLTestMixedContent{
					statuscake.SSLTestMixedContent{
						Type: "script",
						URL:  "http://maps.google.com/maps?file=api&amp;v=2&amp;key=ABQIAAAAoxfgxU9nJyr-dWvHmeWo-BQcyx7nWAFtqJDMM7hWmXgT6wvpkRRDd53JONagHOWmkMUdNSh7L6C0Td",
					},
					statuscake.SSLTestMixedContent{
						Type: "img",
						URL:  "http://www.statuscake.com/wp-content/uploads/2020/07/thumbnail.gif",
					},
				},
				Updated:    statuscake.PtrTime(time.Date(2021, 1, 6, 12, 7, 35, 0, time.UTC)),
				ValidFrom:  statuscake.PtrTime(time.Date(2020, 5, 30, 0, 0, 0, 0, time.UTC)),
				ValidUntil: statuscake.PtrTime(time.Date(2022, 5, 30, 23, 59, 0, 0, time.UTC)),
			},
			statuscake.SSLTest{
				ID:         "2",
				Paused:     true,
				WebsiteURL: "https://www.statuscake.com",
				CheckRate:  statuscake.SSLTestCheckRateOneHour,
				ContactGroups: []string{
					"123",
				},
				AlertAt:       []int32{1, 7, 30},
				AlertBroken:   true,
				AlertExpiry:   true,
				AlertMixed:    true,
				AlertReminder: true,
				Flags: &statuscake.SSLTestFlags{
					FollowRedirects: true,
				},
				FollowRedirects: true,
				Hostname:        statuscake.PtrString("svc.example.com"),
				LastReminder:    statuscake.PtrInt32(0),
				MixedContent:    []statuscake.SSLTestMixedContent{},
				Updated:         statuscake.PtrTime(time.Date(2021, 1, 6, 12, 7, 35, 0, time.UTC)),
				UserAgent:       statuscake.PtrString("Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36 OPR/73.0.3856.344"),
			},
		})
	})

	t.Run("returns an error when the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(mustRead(t, "testdata/fetch-resource-error.json"))
		}))
		defer s.Close()

		_, err := c.ListSslTests(context.Background()).Execute()
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

func TestUpdateSSLTest(t *testing.T) {
	t.Run("returns a no content status on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expectEqual(t, mustParse(t, r), url.Values{
				"alert_at_csv":       []string{"1,14,25"},
				"alert_broken":       []string{"false"},
				"alert_expiry":       []string{"false"},
				"alert_mixed":        []string{"false"},
				"alert_reminder":     []string{"false"},
				"check_rate":         []string{"1800"},
				"contact_groups_csv": []string{""},
				"follow_redirects":   []string{"false"},
				"paused":             []string{"true"},
			})

			w.WriteHeader(http.StatusNoContent)
		}))
		defer s.Close()

		err := c.UpdateSslTest(context.Background(), "2").
			Paused(true).
			CheckRate(statuscake.SSLTestCheckRateThirtyMinutes).
			ContactGroups([]string{}).
			AlertAt([]string{"1", "14", "25"}).
			AlertBroken(false).
			AlertExpiry(false).
			AlertMixed(false).
			AlertReminder(false).
			FollowRedirects(false).
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

		err := c.UpdateSslTest(context.Background(), "3").
			Paused(true).
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
