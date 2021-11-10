// The MIT License (MIT)
//
// Copyright (c) 2021 Micke Sommar <me@mickesommar.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package certls
package certls

// Options for reading and showing certificates.
type Options struct {
	// SkipTLSVerify do not verify TLS certificate.
	SkipTLSVerify bool `json:"skiptlsverify"`
	// ShowCA also show CA certificates.
	ShowCA bool `json:"showca"`
	// ShowIssuer show issuer in output/fields.
	ShowIssuer bool `json:"showissuer"`
	// ShowDNSNames show DNS Names in output/fields.
	ShowDNSNames bool `json:"showdnsnames"`
	// TimeOut timeout for TCP connection.
	TimeOut int `json:"timeout"`
}