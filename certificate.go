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

import "strings"

// Certificate
type Certificate struct {
	ConnectionString string   `json:"connectionstring" yaml:"connectionstring"`
	CommonName       string   `json:"commonname" yaml:"commonname"`
	Created          string   `json:"created" yaml:"created"`
	Expire           string   `json:"expire" yaml:"expire"`
	Issuer           string   `json:"issuer" yaml:"issuer"`
	DNSNames         []string `json:"dnsnames" yaml:"dnsnames"`
	Options          Options  `json:"-" yaml:"-"`
}

// NewCertificate return a new certificate.
func NewCertificate(connectionString, commonName, created, expire, issuer string, dnsNames []string, options Options) Certificate {
	return Certificate{
		ConnectionString: connectionString,
		CommonName:       commonName,
		Created:          created,
		Expire:           expire,
		Issuer:           issuer,
		DNSNames:         dnsNames,
		Options:          options,
	}
}

// Fields return a slice of strings with Certificate fields.
func (c *Certificate) Fields() []string {
	f := make([]string, 0)
	f = append(f, c.ConnectionString, c.CommonName, c.Created, c.Expire)
	if c.Options.ShowAll {
		f = append(f, c.Issuer, strings.Join(c.DNSNames, ", "))
	}
	return f
}

// CertificateFieldsNames return a slice of string with fileds names.
func CertificateFieldsNames(options Options) []string {
	f := make([]string, 0)
	f = append(f, "Host", "Common Name", "Created", "Expire")
	if options.ShowAll {
		f = append(f, "Issuer", "DNS Names")
	}
	return f
}
