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
package certls_test

import (
	"testing"

	"gitea.mickesommar.com/golang/certls"
)

func TestReadJSONConfig(t *testing.T) {
	c := certls.NewConfig("test_data/hosts.json")
	if err := c.Read(); err != nil {
		t.Errorf("could not read config: %v", err)
	}
	if len(c.Hosts) != 3 {
		t.Errorf("want: 3, got: %d", len(c.Hosts))
	}
}

func TestReadYAMLConfig(t *testing.T) {
	c := certls.NewConfig("test_data/hosts.yaml")
	if err := c.Read(); err != nil {
		t.Errorf("could not read config: %v", err)
	}
	if len(c.Hosts) != 3 {
		t.Errorf("want: 3, got: %d", len(c.Hosts))
	}
}
