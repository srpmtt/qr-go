package main

import (
	"flag"
	"fmt"
	"os"
	"qr-go/commands"
)

var usage string = `
  A simple qr code tool

  Commands:
    create    create new qr code
    version   prints version info
  `

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprint(usage))
	}

	flag.Parse()
	if flag.NArg() < 1 {
		usageAndExit("")
	}

	var cmd *commands.Command

	switch os.Args[1] {
	case "create":
		cmd = commands.CreateCommand()
	case "version":
		cmd = commands.VersionCommand()
	default:
		usageAndExit(fmt.Sprintf("%s is not a valid command.\n", os.Args[1]))
	}

	cmd.Init(os.Args[2:])
	cmd.Run()
}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprint(os.Stderr, "\n")
	}

	flag.Usage()
	os.Exit(0)
}
