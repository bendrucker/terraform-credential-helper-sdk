package credentialhelper

import (
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
					Get("app.terraform.io").
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
