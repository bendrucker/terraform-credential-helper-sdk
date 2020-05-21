package main

import (
	"flag"
	"os"

	credentialhelper "github.com/bendrucker/terraform-credential-helper-sdk"
)

func main() {
	flags := flag.NewFlagSet("foo", flag.ContinueOnError)
	flags.String("boop", "beep", "beep boop")
	cli := credentialhelper.New("foo", "dev", new(FakeHelper), flags)
	cli.Run(os.Args[1:])
}
