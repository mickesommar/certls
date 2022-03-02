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
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Config store configuration for certls.
type Config struct {
	Hosts []Host `json:"hosts"`
	path  string
}

// NewConfig return a Config with path.
func NewConfig(path string) Config {
	return Config{
		path: path,
	}
}

// Exists check if config file exists.
func (c *Config) Exists() bool {
	if _, err := os.Stat(c.path); os.IsNotExist(err) {
		return false
	}
	return true
}

// Read config from disk, try to read both JSON or YAML.
func (c *Config) Read() error {
	if _, err := os.Stat(c.path); os.IsNotExist(err) {
		return fmt.Errorf("file is missing: %s", c.path)
	}

	buf, err := os.ReadFile(c.path)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	switch strings.ToLower(filepath.Ext(c.path)) {
	case ".json":
		if err := json.Unmarshal(buf, &c); err != nil {
			return fmt.Errorf("error unmarshal JSON: %v", err)
		}
	case ".yaml":
		if err := yaml.Unmarshal(buf, &c); err != nil {
			return fmt.Errorf("error unmarshal YAML: %v", err)
		}
	default:
		return fmt.Errorf("wrong file format, need to be JSON or YAML")
	}
	return nil
}
