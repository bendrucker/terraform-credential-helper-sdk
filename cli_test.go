package credentialhelper

import (
	"flag"
	"fmt"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCLIRun(t *testing.T) {
	cases := []struct {
		Name   string
		Args   []string
		Expect func(*MockHelper)
		Code   int
	}{
		{
			Name: "success",
			Args: []string{"get", "app.terraform.io"},
			Expect: func(helper *MockHelper) {
				helper.EXPECT().
					Get("app.terraform.io", gomock.Any()).
					Return([]byte("{}"), nil)
			},
			Code: 0,
		},
		{
			Name: "undefined flags",
			Args: []string{"-foo=bar"},
			Code: 1,
		},
		{
			Name: "no args",
			Code: 127,
		},
		{
			Name: "too many args",
			Args: []string{"get", "app.terraform.io", "foo"},
			Code: 1,
		},
		{
			Name: "no hostname",
			Args: []string{"get"},
			Code: 1,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			helper := NewMockHelper(ctrl)
			if tc.Expect != nil {
				tc.Expect(helper)
			}
			cli := New("helper", "", helper, nil)

			code, err := cli.Run(tc.Args)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if code != tc.Code {
				t.Errorf("unexpected exit code: wanted %d, got %d", tc.Code, code)
			}
		})
	}
}

func ExampleCLI_Get() {
	cli := New("example", "dev", new(ExampleHelper), nil)

	_, _ = cli.Run([]string{"get", "app.terraform.io"})
	// Output:
	// Getting credentials for app.terraform.io
	// {"token":"secret"}
}

func ExampleCLI_Store() {
	cli := New("example", "dev", new(ExampleHelper), nil)

	cli.Stdin = strings.NewReader(`{"token":"secret"}`)
	_, _ = cli.Run([]string{"store", "app.terraform.io"})
	// Output:
	// Storing credentials for app.terraform.io: {"token":"secret"}
}

func ExampleCLI_Forget() {
	cli := New("example", "dev", new(ExampleHelper), nil)

	_, _ = cli.Run([]string{"forget", "app.terraform.io"})
	// Output:
	// Forgetting credentials for app.terraform.io
}

func ExampleCLI_Flags() {
	flags := flag.NewFlagSet("example", flag.ContinueOnError)
	flags.Bool("insecure", false, "")

	cli := New("example", "dev", new(ExampleHelper), flags)

	_, _ = cli.Run([]string{"--insecure", "get", "app.terraform.io"})
	// Output:
	// Getting credentials for app.terraform.io
	// with insecure = true
	// {"token":"secret"}
}

type ExampleHelper struct{}

func (h *ExampleHelper) Get(hostname string, f *flag.FlagSet) ([]byte, error) {
	fmt.Println("Getting credentials for", hostname)

	if insecure := f.Lookup("insecure"); insecure != nil {
		fmt.Println("with insecure =", insecure.Value.(flag.Getter).Get().(bool))
	}

	return []byte(`{"token":"secret"}`), nil
}

func (h *ExampleHelper) Store(hostname string, b []byte, f *flag.FlagSet) error {
	fmt.Printf("Storing credentials for %s: %s", hostname, string(b))
	return nil
}

func (h *ExampleHelper) Forget(hostname string, f *flag.FlagSet) error {
	fmt.Println("Forgetting credentials for", hostname)
	return nil
}
