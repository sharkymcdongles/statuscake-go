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

package throttle_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/StatusCakeDev/statuscake-go/throttle"
)

// Basic handler to return a successful response.
func OK(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) }

func TestTransport(t *testing.T) {
	t.Run("returns from a request when the throttling constraints are not met", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(OK))
		defer srv.Close()

		client := &http.Client{
			Transport: throttle.NewWithDefaultTransport(throttle.LimiterFunc(func(ctx context.Context) error {
				return nil
			})),
		}

		_, err := client.Get(srv.URL)
		if err != nil {
			t.Errorf("expected: no error, got: %+v", err)
		}
	})

	t.Run("returns an error when the throttling constaints are met", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(OK))
		defer srv.Close()

		client := &http.Client{
			Transport: throttle.NewWithDefaultTransport(throttle.LimiterFunc(func(ctx context.Context) error {
				return errors.New("something went wrong")
			})),
		}

		_, err := client.Get(srv.URL)
		if err == nil {
			t.Error("expected: error, got <nil>")
		}
	})
}

func TestGroup(t *testing.T) {
	t.Run("returns nil when there are no errors returned from any nested limiters", func(t *testing.T) {
		l := throttle.NewGroup(
			throttle.LimiterFunc(func(ctx context.Context) error {
				return nil
			}),
			throttle.LimiterFunc(func(ctx context.Context) error {
				return nil
			}),
		)

		if err := l.Wait(context.Background()); err != nil {
			t.Errorf("expected: no error, got: %+v", err)
		}
	})

	t.Run("returns an error if any of the nested limiters returns an error", func(t *testing.T) {
		l := throttle.NewGroup(
			throttle.LimiterFunc(func(ctx context.Context) error {
				return nil
			}),
			throttle.LimiterFunc(func(ctx context.Context) error {
				return errors.New("something went wrong")
			}),
			throttle.LimiterFunc(func(ctx context.Context) error {
				return nil
			}),
		)

		if err := l.Wait(context.Background()); err == nil {
			t.Error("expected: error, got: <nil>")
		}
	})
}
