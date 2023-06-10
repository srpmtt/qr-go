package commands

import (
	"flag"
	"fmt"
	"os"
)

var versionCommandUsage = `print version info`

var version = "0.0.1"

var versionFunc = func(cmd *Command, args []string) {
	fmt.Printf("v%s", version)
	os.Exit(0)
}

func VersionCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("version", flag.ExitOnError),
		Execute: versionFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprint(os.Stderr, versionCommandUsage)
	}

	return cmd
}
