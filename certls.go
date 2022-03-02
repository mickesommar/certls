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

import (
	"crypto/tls"
	"net"
	"time"
)

// Certls
type Certls struct {
	options Options
}

// NewCertls return a ner Certtls.
func NewCertls(options Options) Certls {
	return Certls{
		options: options,
	}
}

// Connect to address, read remote certificate.
func (c *Certls) Connect(h Host) ([]Certificate, error) {
	dialer := net.Dialer{
		Timeout: time.Second * time.Duration(c.options.TimeOut),
	}

	conn, err := tls.DialWithDialer(&dialer, "tcp", h.String(), &tls.Config{
		InsecureSkipVerify: c.options.SkipTLSVerify,
	})
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	certs := make([]Certificate, 0)
	for _, cert := range conn.ConnectionState().PeerCertificates {
		if cert.IsCA && !c.options.ShowAll {
			continue
		}
		certs = append(certs, NewCertificate(
			h.String(),
			cert.Subject.CommonName,
			cert.Issuer.CommonName,
			cert.NotBefore,
			cert.NotAfter,
			cert.DNSNames, c.options))
	}
	return certs, nil
}
