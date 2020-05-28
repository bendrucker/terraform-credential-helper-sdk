# terraform-credential-helper-sdk 
[![tests](https://github.com/bendrucker/terraform-credential-helper-sdk/workflows/tests/badge.svg?branch=master)](https://github.com/bendrucker/terraform-credential-helper-sdk/actions?query=workflow%3Atests)
[![GoDoc](https://godoc.org/github.com/bendrucker/terraform-credential-helper-sdk?status.svg)](https://godoc.org/github.com/bendrucker/terraform-credential-helper-sdk)

> Framework for creating [Terraform credentials helpers](https://www.terraform.io/docs/commands/cli-config.html#credentials-helpers)

## Usage

```go
func main () {
  cli := New("example", "dev", new(ExampleHelper))
  code, err = cli.Run(os.Args[1:])
  
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }

  os.Exit(code)
}

type Helper struct{
  MyFlag string
}

func (h *Helper) Flags() *flag.FlagSet {
  flags := flag.NewFlagSet("example", flag.ContinueOnError)
  flags.StringVar(&h.MyFlag, "my-flag", "default", "usage")
  return flags
}

func (h *Helper) Get(hostname string) ([]byte, error) {
	return nil, nil
}

func (h *Helper) Store(hostname string, b []byte) error {
	return nil
}

func (h *Helper) Forget(hostname string) error {
	return nil
}
```

## License

MIT © [Ben Drucker](http://bendrucker.me)
