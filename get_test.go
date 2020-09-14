package credentialhelper

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mitchellh/cli"
	"github.com/stretchr/testify/assert"
)

func TestGetCommand(t *testing.T) {
	cases := []struct {
		Hostname string
		Result   []byte
		Error    error

		Code   int
		Output string
	}{
		{
			Hostname: "app.terraform.io",
			Result:   []byte(`{"token":"foo"}`),

			Code:   0,
			Output: `{"token":"foo"}`,
		},
		{
			Hostname: "foo.terraform.io",
			Result:   []byte{},
			Error:    errors.New("The specified item could not be found in the keyring"),

			Code:   1,
			Output: "error getting credentials: The specified item could not be found in the keyring\n",
		},
	}

	for _, tc := range cases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ui := cli.NewMockUi()
		helper := NewMockHelper(ctrl)
		cmd := &getCommand{
			meta: &meta{
				UI:     ui,
				Helper: helper,
			},
		}

		helper.EXPECT().Get(tc.Hostname).Return(tc.Result, tc.Error)

		code := cmd.Run([]string{tc.Hostname})
		if code != tc.Code {
			t.Fatalf("expected command to exit with %d, got %d", tc.Code, code)
		}

		if code == 0 {
			assert.JSONEq(t, tc.Output, ui.OutputWriter.String())
			assert.Empty(t, ui.ErrorWriter.Bytes())
		} else {
			assert.Equal(t, tc.Output, string(ui.ErrorWriter.Bytes()))
		}
	}
}
