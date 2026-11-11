package main

import (
	"github.com/haspr-dev/go-haspr-cli/internal/cmd"
)

func main() {
	cmd.SetVersionInfo(version, commit, date, dirty)
	cmd.Execute()
}
