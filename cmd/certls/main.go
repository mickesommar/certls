// Generated using urfaveCLI scafolding utility.
package main

import (
	"fmt"
	"runtime"

	"gitea.mickesommar.com/golang/certls/cmd/certls/cmd"
)

var (
	version    string = "development"
	commitHash string
	buildDate  string
	buildTime  string
)

// main
func main() {
	cmd.Execute(fmt.Sprintf("%s %s %s %s %s", version, buildDate, buildTime, runtime.Version(), commitHash))
}
