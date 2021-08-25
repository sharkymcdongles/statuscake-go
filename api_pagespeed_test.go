package statuscake_test

import (
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/StatusCakeDev/statuscake-go"
)

func TestCreatePagespeedTest(t *testing.T) {
	t.Run("returns a created status on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expectEqual(t, mustParse(t, r), url.Values{
				"alert_bigger":       []string{"200"},
				"alert_slower":       []string{"300"},
				"alert_smaller":      []string{"20"},
				"check_rate":         []string{"3600"},
				"contact_groups_csv": []string{"123"},
				"location_iso":       []string{"DE"},
				"name":               []string{"statuscake.com"},
				"paused":             []string{"false"},
				"website_url":        []string{"https://www.statuscake.com"},
			})

			w.WriteHeader(http.StatusCreated)
			w.Write(mustRead(t, "testdata/create-resource-success.json"))
		}))
		defer s.Close()

		res, err := c.CreatePagespeedTest(context.Background()).
			Name("statuscake.com").
			WebsiteURL("https://www.statuscake.com").
			CheckRate(statuscake.PagespeedTestCheckRateOneHour).
			ContactGroups([]string{
				"123",
			}).
			AlertBigger(200).
			AlertSlower(300).
			AlertSmaller(20).
			LocationISO(statuscake.PagespeedTestLocationISOGermany).
			Paused(false).
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

		_, err := c.CreatePagespeedTest(context.Background()).
			Name("statuscake.com").
			WebsiteURL("this,is,not,valid").
			CheckRate(statuscake.PagespeedTestCheckRateOneHour).
			ContactGroups([]string{
				"123",
			}).
			AlertBigger(200).
			AlertSlower(300).
			AlertSmaller(20).
			LocationISO(statuscake.PagespeedTestLocationISOGermany).
			Paused(false).
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

func TestDeletePagespeedTest(t *testing.T) {
	t.Run("returns a no content status on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
		}))
		defer s.Close()

		err := c.DeletePagespeedTest(context.Background(), "2").Execute()
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

		err := c.DeletePagespeedTest(context.Background(), "3").Execute()
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

func TestGetPagespeedTest(t *testing.T) {
	t.Run("returns a pagespeed test on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(mustRead(t, "testdata/get-pagespeed-test-success.json"))
		}))
		defer s.Close()

		test, err := c.GetPagespeedTest(context.Background(), "2").Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, test.Data, statuscake.PagespeedTest{
			ID:         "2",
			Name:       "statuscake.com",
			Paused:     false,
			WebsiteURL: "https://www.statuscake.com",
			CheckRate:  statuscake.PagespeedTestCheckRateOneHour,
			ContactGroups: []string{
				"123",
			},
			AlertBigger:  200,
			AlertSlower:  300,
			AlertSmaller: 20,
			Location:     "DE4.PAGESPEED.STATUSCAKE.NET",
			LocationISO:  statuscake.PagespeedTestLocationISOGermany,
			LatestStats: &statuscake.PagespeedTestStats{
				Requests:    27,
				HasIssue:    true,
				LatestIssue: statuscake.PtrString("The Total Load Time of the Page (20216/ms) is larger than the alert threshold of 300/ms"),
			},
		})
	})

	t.Run("returns an error when the request fails", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(mustRead(t, "testdata/fetch-resource-error.json"))
		}))
		defer s.Close()

		_, err := c.GetPagespeedTest(context.Background(), "3").Execute()
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

func TestListPagespeedTests(t *testing.T) {
	t.Run("returns a list of pagespeed tests on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(mustRead(t, "testdata/list-pagespeed-tests-success.json"))
		}))
		defer s.Close()

		tests, err := c.ListPagespeedTests(context.Background()).Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, tests.Data, []statuscake.PagespeedTest{
			statuscake.PagespeedTest{
				ID:            "1",
				Name:          "google.com",
				Paused:        false,
				WebsiteURL:    "https://www.google.com",
				CheckRate:     statuscake.PagespeedTestCheckRateFiveMinutes,
				ContactGroups: []string{},
				AlertBigger:   0,
				AlertSlower:   0,
				AlertSmaller:  0,
				Location:      "PAGESPD-AU2",
				LocationISO:   statuscake.PagespeedTestLocationISOAustralia,
			},
			statuscake.PagespeedTest{
				ID:         "2",
				Name:       "statuscake.com",
				Paused:     false,
				WebsiteURL: "https://www.statuscake.com",
				CheckRate:  statuscake.PagespeedTestCheckRateOneHour,
				ContactGroups: []string{
					"123",
				},
				AlertBigger:  200,
				AlertSlower:  300,
				AlertSmaller: 20,
				Location:     "DE4.PAGESPEED.STATUSCAKE.NET",
				LocationISO:  statuscake.PagespeedTestLocationISOGermany,
				LatestStats: &statuscake.PagespeedTestStats{
					Requests:    27,
					HasIssue:    true,
					LatestIssue: statuscake.PtrString("The Total Load Time of the Page (20216/ms) is larger than the alert threshold of 300/ms"),
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

		_, err := c.ListPagespeedTests(context.Background()).Execute()
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

func TestListPagespeedTestHistory(t *testing.T) {
	t.Run("returns a list of pagespeed test history results on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write(mustRead(t, "testdata/list-pagespeed-test-history-success.json"))
		}))
		defer s.Close()

		histroy, err := c.ListPagespeedTestHistory(context.Background(), "2").
			Days(10).
			Execute()
		if err != nil {
			t.Fatal(err.Error())
		}

		expectEqual(t, histroy.Data, statuscake.PagespeedTestHistoryData{
			Aggregated: statuscake.PagespeedTestHistoryDataAggregated{
				Loadtime: statuscake.PagespeedTestHistoryDataAggregatedLoadtime{
					Min: 0,
					Max: 1490,
					Avg: 1490,
				},
				Filesize: statuscake.PagespeedTestHistoryDataAggregatedFilesize{
					Min: 0,
					Max: 598.384,
					Avg: 598.384,
				},
				Requests: statuscake.PagespeedTestHistoryDataAggregatedRequests{
					Min: 0,
					Max: 4,
					Avg: 4,
				},
				Results: 1,
			},
			Results: map[string]statuscake.PagespeedTestHistoryResult{
				"1611241767": statuscake.PagespeedTestHistoryResult{
					Created:     time.Date(2021, 1, 21, 15, 9, 27, 0, time.UTC),
					Loadtime:    1490,
					Requests:    4,
					Filesize:    598.384,
					Throttling:  statuscake.PagespeedTestThrottlingNone,
					HARLocation: "https://16a0fd6b5b5bece1d29a-7aa19249e604542958e6a694f67d0bbf.ssl.cf5.rackcdn.com/53a6b075-3b93-4752-b707-93b08fe5ae44.json",
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

		_, err := c.ListPagespeedTestHistory(context.Background(), "3").
			Days(10).
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

func TestUpdatePagespeedTest(t *testing.T) {
	t.Run("returns a no content status on success", func(t *testing.T) {
		s, c := createTestEndpoint(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expectEqual(t, mustParse(t, r), url.Values{
				"alert_bigger":       []string{"10"},
				"alert_slower":       []string{"100"},
				"alert_smaller":      []string{"1"},
				"check_rate":         []string{"1800"},
				"contact_groups_csv": []string{""},
				"location_iso":       []string{"UK"},
				"name":               []string{"example.com"},
				"paused":             []string{"true"},
			})

			w.WriteHeader(http.StatusNoContent)
		}))
		defer s.Close()

		err := c.UpdatePagespeedTest(context.Background(), "2").
			Name("example.com").
			Paused(true).
			CheckRate(statuscake.PagespeedTestCheckRateThirtyMinutes).
			ContactGroups([]string{}).
			AlertBigger(10).
			AlertSlower(100).
			AlertSmaller(1).
			LocationISO(statuscake.PagespeedTestLocationISOUnitedKingdom).
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

		err := c.UpdatePagespeedTest(context.Background(), "3").
			Name("example.com").
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
