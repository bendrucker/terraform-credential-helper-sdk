package credentialhelper

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mitchellh/cli"
	"github.com/stretchr/testify/assert"
)

func TestGetCommand(t *testing.T) {
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

	helper.EXPECT().
		Get("app.terraform.io").
		Return([]byte(`{"token":"foo"}`), nil)

	status := cmd.Run([]string{"app.terraform.io"})
	if status != 0 {
		t.Fatalf("expected command to exit with 0, got %d", status)
	}

	assert.JSONEq(t, `{"token":"foo"}`, ui.OutputWriter.String())
	assert.Empty(t, ui.ErrorWriter.Bytes())
}
