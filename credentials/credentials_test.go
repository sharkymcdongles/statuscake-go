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

package credentials_test

import (
	"encoding/base64"
	"net/http"
	"testing"

	"github.com/StatusCakeDev/statuscake-go/credentials"
)

func TestBasicAuthentication(t *testing.T) {
	t.Run("sets basic authenticaton credentials on a HTTP request", func(t *testing.T) {
		r, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Fatalf("failed to create request: %+v", err)
		}

		c := credentials.NewBasicAuthentication("jean-luc", "enterprise")
		c.AddCredentials(r)

		authorisationHeader := r.Header.Get("Authorization")
		expected := "Basic " + base64.StdEncoding.EncodeToString([]byte("jean-luc:enterprise"))

		if authorisationHeader != expected {
			t.Errorf("expected: %s, got: %s", expected, authorisationHeader)
		}
	})
}

func TestBearerToken(t *testing.T) {
	t.Run("sets bearer authenticaton credentials on a HTTP request", func(t *testing.T) {
		r, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Fatalf("failed to create request: %+v", err)
		}

		c := credentials.NewBearerWithStaticToken("this-is-a-token")
		c.AddCredentials(r)

		authorisationHeader := r.Header.Get("Authorization")
		expected := "Bearer this-is-a-token"

		if authorisationHeader != expected {
			t.Errorf("expected: %s, got: %s", expected, authorisationHeader)
		}
	})

	t.Run("when a custom function is given, sets bearer authenticaton credentials on a HTTP request", func(t *testing.T) {
		r, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Fatalf("failed to create request: %+v", err)
		}

		c := credentials.NewBearer(func(r *http.Request) string {
			return "this-is-a-token"
		})
		c.AddCredentials(r)

		authorisationHeader := r.Header.Get("Authorization")
		expected := "Bearer this-is-a-token"

		if authorisationHeader != expected {
			t.Errorf("expected: %s, got: %s", expected, authorisationHeader)
		}
	})
}
