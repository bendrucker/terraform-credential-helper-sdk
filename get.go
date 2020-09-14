package credentialhelper

import (
	"fmt"
)

type getCommand struct {
	*meta
}

func (c *getCommand) Run(args []string) int {
	creds, err := c.Helper.Get(args[0])
	if err != nil {
		c.UI.Error(fmt.Sprintf("error getting credentials: %v", err))
		return 1
	}

	c.UI.Output(string(creds))
	return 0
}

func (c *getCommand) Synopsis() string {
	return "Retrieve the credentials for the given hostname"
}

func (c *getCommand) Help() string {
	return c.help("get", `
To retrieve credentials, Terraform will run the "get" command with any configured arguments,
plus the hostname for which credentials should be retrieved.
	`)
}
