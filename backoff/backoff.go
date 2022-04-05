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

package backoff

import (
	"math/rand"
	"time"
)

// Config defines the configuration options for backoff.
type Config struct {
	BaseDelay  time.Duration
	Multiplier float64
	Jitter     float64
	MaxDelay   time.Duration
}

// DefaultConfig is a backoff configuration with default values. This should be
// useful for callers who want to configure backoff with non-default values
// only for a subset of the options.
var DefaultConfig = Config{
	BaseDelay:  1.0 * time.Second,
	Multiplier: 1.0,
	Jitter:     0.2,
	MaxDelay:   120 * time.Second,
}

func addJitter(backoff, jitter float64) float64 {
	backoff *= 1 + jitter*(rand.Float64()*2-1)
	if backoff < 0 {
		return 0
	}
	return backoff
}
