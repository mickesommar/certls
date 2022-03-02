// Package cmd application root file.
package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"gitea.mickesommar.com/golang/certls"
	cli "github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

// Consts
const (
	formatTable string = "table"
	formatCSV   string = "csv"
	formatJSON  string = "json"
	formatYAML  string = "yaml"
)

// app global cli.App.
var app = &cli.App{}

// Execute
func Execute(version string) {
	// App
	app.Name = "certls"
	app.Usage = "Read SSL Certificates from remote hosts."
	app.Authors = []*cli.Author{
		{
			Name:  "Micke Sommar",
			Email: "me@mickesommar.com",
		},
	}
	app.Copyright = "Micke Sommar 2022"
	app.Version = version

	// Description
	app.Description = `
Load a list for addresses and ports (hosts), loop throught all hosts and resolve
SSL certificates.

OUTPUT
It is possible to format the output:
* TEXT
* CSV
* JSON
* YAML

INPUT FILE:
It is possible to use JSON or YAML as input file/host-file.

JSON:
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

YAML:
---
hosts:
  - address: www.google.com
    port: 443
  - address: www.bing.com
    port: 443

SHOW ALL
As default, only host, common name, create date and expire date, will 
be displayed (if using JSON or YAML, all will be displayed).

It is possible to use: --show-all, to display CA issuer and Domain Names.
`

	// Flags
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "host-file",
			Aliases: []string{"f"},
			Usage:   "Path to JSON or YAML file, containg hosts", Required: true,
			EnvVars: []string{"CERTLS_HOST_FILE"},
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   fmt.Sprintf("Format output (%s, %s, %s or %s)", formatTable, formatCSV, formatJSON, formatYAML),
			Value:   formatTable,
		},
		&cli.IntFlag{
			Name:    "timeout",
			Aliases: []string{"t"},
			Usage:   "Timeout before closing connection, if host does not respond (in seconds)",
			Value:   1,
			EnvVars: []string{"CERTLS_TIMEOUT"},
		},
		&cli.BoolFlag{
			Name:    "skip-tls-verify",
			Aliases: []string{"i"},
			Usage:   "Skip verification of TLS certificates (INSECURE)",
			EnvVars: []string{"CERTLS_SKIP_TLS_VERIFY"},
		},
		&cli.BoolFlag{
			Name:    "show-all",
			Aliases: []string{"a"},
			Usage:   "Show all",
			EnvVars: []string{"CERTLS_SHOW_ALL"},
		},
	}

	// Action
	app.Action = func(c *cli.Context) error {
		// Load host file.
		config := certls.NewConfig(c.String("host-file"))
		if !config.Exists() {
			return fmt.Errorf("config file does not exists: %s", c.String("host-file"))
		}
		if err := config.Read(); err != nil {
			return fmt.Errorf("error reading config file: %v", err)
		}

		// Options
		options := certls.Options{
			SkipTLSVerify: c.Bool("skip-tls-verify"),
			ShowAll:       c.Bool("show-all"),
			TimeOut:       c.Int("timeout"),
		}

		// Scan remote hosts for SSL/TLS certificates.
		allCerts := make([]certls.Certificate, 0)
		certLS := certls.NewCertls(options)
		for _, host := range config.Hosts {
			certs, err := certLS.Connect(host)
			if err != nil {
				log.Printf("error: %s %v\n", host.String(), err)
				continue
			}
			allCerts = append(allCerts, certs...)
		}

		// Format output/print to terminal.
		switch strings.ToLower(c.String("output")) {
		case formatTable:
			tabWriter := tabwriter.NewWriter(os.Stdout, 20, 0, 4, ' ', 0)
			fmt.Fprintln(tabWriter, strings.Join(certls.CertificateFieldsNames(options), "\t"))
			for _, cert := range allCerts {
				fmt.Fprintf(tabWriter, "%s\n", strings.Join(cert.Fields(), "\t"))
			}
			if err := tabWriter.Flush(); err != nil {
				return fmt.Errorf("flush tabwriter: %v", err)
			}
		case formatCSV:
			csvWriter := csv.NewWriter(os.Stdout)
			csvWriter.UseCRLF = true
			csvWriter.Comma = ';'
			if err := csvWriter.Write(certls.CertificateFieldsNames(options)); err != nil {
				return fmt.Errorf("error writing CSV header: %v", err)
			}
			for _, cert := range allCerts {
				csvWriter.Write(cert.Fields())
			}
			csvWriter.Flush()
			if err := csvWriter.Error(); err != nil {
				return fmt.Errorf("error flushing csv writer: %v", err)
			}
		case formatJSON:
			buf, err := json.MarshalIndent(allCerts, " ", " ")
			if err != nil {
				return fmt.Errorf("error marshal certs to JSON: %v", err)
			}
			fmt.Printf("%s\n", buf)
		case formatYAML:
			buf, err := yaml.Marshal(allCerts)
			if err != nil {
				return fmt.Errorf("error marshal certs to YAML: %v", err)
			}
			fmt.Printf("%s\n", buf)
		default:
			return fmt.Errorf("bad output format: %s\n", c.String("output"))
		}
		return nil
	}

	// Update AppHelpTemplate-message.
	cli.AppHelpTemplate = fmt.Sprintf("%s\nREPOSITORY\n", cli.AppHelpTemplate)
	cli.AppHelpTemplate = fmt.Sprintf("%s   https://github.com/mickesommar/certls.git\n", cli.AppHelpTemplate)

	// Run
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
