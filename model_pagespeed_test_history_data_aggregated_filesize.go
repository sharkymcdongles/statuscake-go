/*
 * StatusCake API
 *
 * Copyright (c) 2021 StatusCake
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

package statuscake

import (
	"encoding/json"
)

// PagespeedTestHistoryDataAggregatedFilesize Aggregate filesize (kb)
type PagespeedTestHistoryDataAggregatedFilesize struct {
	Min float32 `json:"min"`
	Max float32 `json:"max"`
	Avg float32 `json:"avg"`
}

// NewPagespeedTestHistoryDataAggregatedFilesize instantiates a new PagespeedTestHistoryDataAggregatedFilesize object.
// This constructor will assign default values to properties that have it
// defined, and makes sure properties required by API are set, but the set of
// arguments will change when the set of required properties is changed.
func NewPagespeedTestHistoryDataAggregatedFilesize(min float32, max float32, avg float32) *PagespeedTestHistoryDataAggregatedFilesize {
	return &PagespeedTestHistoryDataAggregatedFilesize{
		Min: min,
		Max: max,
		Avg: avg,
	}
}

func (o PagespeedTestHistoryDataAggregatedFilesize) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["min"] = o.Min
	}
	if true {
		toSerialize["max"] = o.Max
	}
	if true {
		toSerialize["avg"] = o.Avg
	}
	return json.Marshal(toSerialize)
}
