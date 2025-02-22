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
	"fmt"
)

// UptimeTestType Uptime check type
type UptimeTestType string

const (
	// UptimeTestTypeDNS a DNS uptime check.
	UptimeTestTypeDNS UptimeTestType = "DNS"
	// UptimeTestTypeHEAD an HTTP uptime check (HEAD verb).
	UptimeTestTypeHEAD UptimeTestType = "HEAD"
	// UptimeTestTypeHTTP an HTTP uptime check (GET or POST verbs).
	UptimeTestTypeHTTP UptimeTestType = "HTTP"
	// UptimeTestTypePING an PING uptime check.
	UptimeTestTypePING UptimeTestType = "PING"
	// UptimeTestTypeSMTP an SMTP uptime check.
	UptimeTestTypeSMTP UptimeTestType = "SMTP"
	// UptimeTestTypeSSH an SSH uptime check.
	UptimeTestTypeSSH UptimeTestType = "SSH"
	// UptimeTestTypeTCP a TCP uptime check.
	UptimeTestTypeTCP UptimeTestType = "TCP"
)

// Unmarshal JSON data into any of the pointers in the type.
func (v *UptimeTestType) UnmarshalJSON(src []byte) error {
	var value string
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}

	ev := UptimeTestType(value)
	if !ev.Valid() {
		return fmt.Errorf("%+v is not a valid UptimeTestType", value)
	}

	*v = ev
	return nil
}

// Valid determines if the value is valid.
func (v UptimeTestType) Valid() bool {
	return v == UptimeTestTypeDNS || v == UptimeTestTypeHEAD || v == UptimeTestTypeHTTP || v == UptimeTestTypePING || v == UptimeTestTypeSMTP || v == UptimeTestTypeSSH || v == UptimeTestTypeTCP
}

// UptimeTestTypeValues returns the values of UptimeTestType.
func UptimeTestTypeValues() []string {
	return []string{
		"DNS",
		"HEAD",
		"HTTP",
		"PING",
		"SMTP",
		"SSH",
		"TCP",
	}
}
