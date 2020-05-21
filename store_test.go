package credentialhelper

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mitchellh/cli"
	"github.com/stretchr/testify/assert"
)

func TestStoreCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ui := cli.NewMockUi()
	helper := NewMockHelper(ctrl)
	cmd := &storeCommand{
		meta: &meta{
			Stdin:  strings.NewReader(`{"token":"foo"}`),
			UI:     ui,
			Helper: helper,
		},
	}

	helper.EXPECT().
		Store("app.terraform.io", []byte(`{"token":"foo"}`), gomock.Any()).
		Return(nil)

	status := cmd.Run([]string{"app.terraform.io"})
	if status != 0 {
		t.Fatalf("expected command to exit with 0, got %d", status)
	}

	assert.Empty(t, ui.OutputWriter.Bytes())
	assert.Empty(t, ui.ErrorWriter.Bytes())
}
