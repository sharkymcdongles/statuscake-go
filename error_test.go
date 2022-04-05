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
	"errors"
	"reflect"
	"testing"

	"github.com/StatusCakeDev/statuscake-go"
)

func TestAPIError(t *testing.T) {
	t.Run("it conforms to the error interface", func(t *testing.T) {
		err := statuscake.NewAPIError(
			"error message",
			errors.New("parent error message"),
		)

		expected := "parent error message: error message"
		got := err.Error()

		if expected != got {
			t.Errorf("expected: %s, got: %s", expected, got)
		}
	})
}

func TestUnwrap(t *testing.T) {
	t.Run("it returns the wrapped error if exists", func(t *testing.T) {
		parent := errors.New("parent error message")

		err := statuscake.NewAPIError(
			"error message",
			parent,
		)

		got := err.Unwrap()
		if got != parent {
			t.Errorf("expected: %+v, got: %+v", parent, got)
		}
	})

	t.Run("it returns nil if no wrapped error exists", func(t *testing.T) {
		err := statuscake.NewAPIError(
			"error message",
			nil,
		)

		got := err.Unwrap()
		if got != nil {
			t.Errorf("expected: <nil>, got: %+v", got)
		}
	})
}

func TestErrors(t *testing.T) {
	t.Run("returns error messages contained within the error", func(t *testing.T) {
		errors := map[string][]string{
			"field": []string{
				"is required",
				"should be numeric",
			},
		}

		err := statuscake.APIError{
			Errors: errors,
		}

		expected := errors
		got := statuscake.Errors(err)

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected: %+v, got: %+v", expected, got)
		}
	})

	t.Run("returns an empty map if the error is of an unexpected type", func(t *testing.T) {
		got := statuscake.Errors(errors.New("unexpected error"))
		expected := map[string][]string{}

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected: %+v, got: %+v", expected, got)
		}
	})
}
