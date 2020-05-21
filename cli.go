package credentialhelper

import (
	"flag"
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

// New creates a new credential helper CLI
func New(name string, version string, helper Helper, flags *flag.FlagSet) *CLI {
	c := cli.NewCLI(name, version)

	if flags == nil {
		flags = flag.NewFlagSet(name, flag.ContinueOnError)
	}

	m := &meta{
		Stdin: os.Stdin,
		UI: &cli.BasicUi{
			Reader:      os.Stdin,
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
		},
		Flags:  flags,
		Helper: helper,
	}

	c.Commands = map[string]cli.CommandFactory{
		"get": func() (cli.Command, error) {
			return &getCommand{m}, nil
		},
		"store": func() (cli.Command, error) {
			return &storeCommand{m}, nil
		},
		"forget": func() (cli.Command, error) {
			return &forgetCommand{m}, nil
		},
	}

	return &CLI{
		cli:  c,
		meta: m,
	}
}

// CLI is a Terraform credential helper CLI
type CLI struct {
	*meta
	cli *cli.CLI
}

// Run runs the CLI with the provided arguments
func (c *CLI) Run(args []string) (int, error) {
	if err := c.Flags.Parse(args); err != nil {
		c.UI.Error(err.Error())
		return 1, nil
	}

	if n := c.Flags.NArg() - 1; n > 1 {
		c.UI.Error(fmt.Sprintf("expected a hostname, got %d args: %v", n, c.Flags.Args()))
		return 1, nil
	} else if n == 0 {
		c.UI.Error(c.cli.HelpFunc(c.cli.Commands))
		return 1, nil
	}

	c.cli.Args = c.Flags.Args()
	return c.cli.Run()
}
