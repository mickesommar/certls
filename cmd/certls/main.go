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

// Package main
package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/tabwriter"

	"gitea.mickesommar.com/golang/certls"
)

var (
	version    string = "development"
	commitHash string
	buildDate  string
	buildTime  string
)

func main() {
	// Flags.
	hostFile := flag.String("host-file", "hosts.json", "Path to file containing hosts (JSON)")
	skipTLSVerify := flag.Bool("skip-tls-verify", false, "Skip verification of TLS certificates (INSECURE)")
	showCA := flag.Bool("show-ca", false, "Show CA Certificate")
	showIssuer := flag.Bool("show-issuer", false, "Show certificate issuer")
	showDNSNames := flag.Bool("show-dns-names", false, "Show DNS names")
	showAll := flag.Bool("show-all", false, "Show all columns")
	out := flag.String("out", "text", "Format output, TEXT, CSV or JSON")
	showVersion := flag.Bool("v", false, "Show version")
	showHelp := flag.Bool("h", false, "Show help message")
	timeOut := flag.Int("timeout", 5, "Timeout before closing connection, if host does not respond")
	flag.Parse()

	// Custom help message and version.
	appName := filepath.Base(os.Args[0])
	version := fmt.Sprintf("%s %s %s %s %s", version, buildDate, buildTime, commitHash, runtime.Version())
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
NAME:
  %s - Read SSL Certificates from remote hosts.

USAGE:
  %s [COMMAND OPTIONS]

VERSION:
  %s

DESCRIPTION:
  Load a list for addresses and ports (hosts), loop throught all hosts and resolve
  SSL certificates.

  OUTPUT
  It is possible to format the output:
  * TEXT
  * CSV
  * JSON (then formating in JSON, Issuer and DNS Names will be vissible, not hidden)

  INPUT FILE:
  A JSON file is used to read hosts:
  ---
  {
    "hosts": [
        {
            "address": "www.google.com",
            "port": "443"
        },
        {
            "address": "www.bing.com",
            "port": "443"
        }
    ]
  }
  ---

AUTHOR:
  Micke Sommar <me@mickesommar.com>

LICENSE:
  MIT

REPOSITORY:
  https://github.com/mickesommar/certls.git

OPTIONS:
`, appName, appName, version)
		flag.PrintDefaults()
	}

	// Show version and help message.
	if *showVersion {
		fmt.Printf("%s - %s\n", appName, version)
		os.Exit(0)
	}
	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}
	// If no flags exists, show help message.
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	// Config.
	config := certls.NewConfig(*hostFile)
	if !config.Exists() {
		fmt.Printf("config file does not exists: %s\n", *hostFile)
		os.Exit(1)
	}
	if err := config.Read(); err != nil {
		log.Printf("error reading config file: %v\n", err)
		os.Exit(1)
	}

	// Options
	options := certls.Options{
		SkipTLSVerify: *skipTLSVerify,
		ShowCA:        *showCA,
		ShowIssuer:    *showIssuer,
		ShowDNSNames:  *showDNSNames,
		TimeOut:       *timeOut,
	}
	if *showAll {
		options = certls.Options{
			SkipTLSVerify: *skipTLSVerify,
			ShowCA:        true,
			ShowIssuer:    true,
			ShowDNSNames:  true,
			TimeOut:       *timeOut,
		}
	}

	// Create writers for formating output.
	var tabWriter *tabwriter.Writer
	var csvWriter *csv.Writer
	var jsonWriter *json.Encoder
	var jsonCerts []certls.Certificate
	switch strings.ToLower(*out) {
	case "text":
		tabWriter = tabwriter.NewWriter(os.Stdout, 20, 0, 4, ' ', 0)
		fmt.Fprintln(tabWriter, strings.Join(certls.CertificateFieldsNames(options), "\t"))
	case "csv":
		csvWriter = csv.NewWriter(os.Stdout)
		csvWriter.UseCRLF = true
		csvWriter.Comma = ';'
		if err := csvWriter.Write(certls.CertificateFieldsNames(options)); err != nil {
			log.Printf("error writing csv header: %v", err)
		}
	case "json":
		jsonWriter = json.NewEncoder(os.Stdout)
		jsonWriter.SetIndent("  ", "  ")
		jsonCerts = make([]certls.Certificate, 0)
	default:
		fmt.Printf("bad output format: %s", *out)
		os.Exit(1)
	}

	// Create a Certls type, start reading remote certificates.
	certLS := certls.NewCertls(options)
	for _, address := range config.Hosts {
		certs, err := certLS.Connect(address)
		if err != nil {
			log.Printf("ERROR: %s %v\n", address.String(), err)
			continue
		}
		for _, cert := range certs {
			switch strings.ToLower(*out) {
			case "text":
				fmt.Fprintln(tabWriter, strings.Join(cert.Fields(), "\t"))
			case "csv":
				if err := csvWriter.Write(cert.Fields()); err != nil {
					log.Printf("error writing certificates to csv output: %s", err)
				}
			case "json":
				jsonCerts = append(jsonCerts, cert)
			}
		}
	}

	// Flush readers (and encode JSON).
	switch strings.ToLower(*out) {
	case "text":
		if err := tabWriter.Flush(); err != nil {
			fmt.Printf("could not flush tabwriter: %v", err)
			os.Exit(1)
		}
	case "csv":
		csvWriter.Flush()
		if err := csvWriter.Error(); err != nil {
			log.Printf("could not flush csvwriter: %v", err)
			os.Exit(1)
		}
	case "json":
		if err := jsonWriter.Encode(jsonCerts); err != nil {
			log.Printf("error encoding certificates to JSON:%v", err)
			os.Exit(1)
		}
	}
}
