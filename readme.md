# terraform-credential-helper-sdk 
[![tests](https://github.com/bendrucker/terraform-credential-helper-sdk/workflows/tests/badge.svg?branch=master)](https://github.com/bendrucker/terraform-credential-helper-sdk/actions?query=workflow%3Atests)
[![GoDoc](https://godoc.org/github.com/bendrucker/terraform-credential-helper-sdk?status.svg)](https://godoc.org/bendrucker/terraform/credential-helper-sdk)

> Framework for creating [Terraform credentials helpers](https://www.terraform.io/docs/commands/cli-config.html#credentials-helpers)

## Usage

```go
func main () {
  cli := New("example", "dev", new(ExampleHelper), flags)
  code, err = cli.Run(os.Args[1:])
  
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }

  os.Exit(code)
}

type Helper struct{}

func (h *Helper) Get(hostname string, f *flag.FlagSet) ([]byte, error) {
	return nil, nil
}

func (h *Helper) Store(hostname string, b []byte, f *flag.FlagSet) error {
	return nil
}

func (h *Helper) Forget(hostname string, f *flag.FlagSet) error {
	return nil
}
```

## License

MIT © [Ben Drucker](http://bendrucker.me)
