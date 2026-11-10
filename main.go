package main

import (
	"github.com/pbolt-dev/go-pbolt-cli/internal/cmd"
)

func main() {
	cmd.SetVersionInfo(version, commit, date, dirty)
	cmd.Execute()
}
