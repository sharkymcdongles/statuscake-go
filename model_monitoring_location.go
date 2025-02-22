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

package statuscake

import (
	"encoding/json"
)

// MonitoringLocation struct for MonitoringLocation
type MonitoringLocation struct {
	// Server description
	Description string `json:"description"`
	// Server IPv4 address
	IPv4 *string `json:"ipv4,omitempty"`
	// Server IPv6 address
	IPv6 *string `json:"ipv6,omitempty"`
	// Server region
	Region string `json:"region"`
	// Server region code
	RegionCode string                   `json:"region_code"`
	Status     MonitoringLocationStatus `json:"status"`
}

// NewMonitoringLocation instantiates a new MonitoringLocation object.
// This constructor will assign default values to properties that have it
// defined, and makes sure properties required by API are set, but the set of
// arguments will change when the set of required properties is changed.
func NewMonitoringLocation(description string, region string, regionCode string, status MonitoringLocationStatus) *MonitoringLocation {
	return &MonitoringLocation{
		Description: description,
		Region:      region,
		RegionCode:  regionCode,
		Status:      status,
	}
}

// MarshalJSON serialises data in the struct to JSON.
func (o MonitoringLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.IPv4 != nil {
		toSerialize["ipv4"] = o.IPv4
	}
	if o.IPv6 != nil {
		toSerialize["ipv6"] = o.IPv6
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["region_code"] = o.RegionCode
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}
