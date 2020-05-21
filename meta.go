package credentialhelper

import (
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/mitchellh/cli"
)

type meta struct {
	Program string
	Stdin   io.Reader
	UI      cli.Ui
	Flags   *flag.FlagSet
	Helper  Helper
}

func (m *meta) help(command string, text string) string {
	return fmt.Sprintf(
		"%s %s\n\n%s\n\n%s",
		m.Program,
		command,
		strings.TrimSpace(text),
		m.options(),
	)
}

func (m *meta) options() string {
	var out string

	m.Flags.VisitAll(func(f *flag.Flag) {
		s := fmt.Sprintf("  -%s", f.Name) // Two spaces before -; see next two comments.
		name, usage := flag.UnquoteUsage(f)
		if len(name) > 0 {
			s += " " + name
		}
		s += "\n    \t"
		s += strings.ReplaceAll(usage, "\n", "\n    \t")

		out += s + "\n\n"
	})

	if out != "" {
		out = "Options:\n" + out
	}

	return strings.TrimSpace(out)
}
