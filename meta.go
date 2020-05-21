package credentialhelper

import (
	"flag"
	"io"

	"github.com/mitchellh/cli"
)

type meta struct {
	Stdin  io.Reader
	UI     cli.Ui
	Flags  *flag.FlagSet
	Helper Helper
}
