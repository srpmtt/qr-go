package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/skip2/go-qrcode"
)

var createCommandUsage = `
  create new qr code
   
  Options:
    -u, --url   qr code url
    -o, --out   output path       
  `

var (
	url string
	out string
)

var createFunc = func(cmd *Command, args []string) {
	if len(url) == 0 {
		errorAndExit("missing url")
	}

	if len(out) == 0 {
		errorAndExit("missing output path")
	}

	err := qrcode.WriteFile(url, qrcode.Medium, 256, out)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("qr code created")
	os.Exit(0)
}

func CreateCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("create", flag.ExitOnError),
		Execute: createFunc,
	}

	cmd.flags.StringVar(&url, "u", "", "")
	cmd.flags.StringVar(&url, "url", "", "")
	cmd.flags.StringVar(&out, "o", "", "")
	cmd.flags.StringVar(&out, "out", "", "")

	cmd.flags.Usage = func() {
		fmt.Fprint(os.Stderr, createCommandUsage)
	}

	return cmd
}
