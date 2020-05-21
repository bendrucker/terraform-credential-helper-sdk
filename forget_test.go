package credentialhelper

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mitchellh/cli"
	"github.com/stretchr/testify/assert"
)

func TestForgetCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ui := cli.NewMockUi()
	helper := NewMockHelper(ctrl)
	cmd := &forgetCommand{
		meta: &meta{
			UI:     ui,
			Helper: helper,
		},
	}

	helper.EXPECT().
		Forget("app.terraform.io", gomock.Any()).
		Return(nil)

	status := cmd.Run([]string{"app.terraform.io"})
	if status != 0 {
		t.Fatalf("expected command to exit with 0, got %d", status)
	}

	assert.Empty(t, ui.OutputWriter.Bytes())
	assert.Empty(t, ui.ErrorWriter.Bytes())
}
