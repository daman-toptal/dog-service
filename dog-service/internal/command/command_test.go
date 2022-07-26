package command

import (
	"dog-service/util/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() {
	config.SetString("grpc.name", "localhost")
	config.SetString("grpc.port", "3001")
}

func TestCommandInvalidCommand(t *testing.T) {
	setup()
	_, err := NewCommand("get photos hound")
	assert.NotNil(t, err, "err should not be nil")
	assert.Equal(t, "invalid command", err.Error())
}

func TestCommandGetPhoto(t *testing.T) {
	setup()
	command, err := NewCommand("get photo hound")
	assert.Nil(t, err, "err should be nil")

	resp, err := command.Execute()
	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, "success", resp)
}

func TestCommandGetPhotoBlankName(t *testing.T) {
	setup()
	command, err := NewCommand("get photo ")
	assert.Nil(t, err, "err should be nil")

	_, err = command.Execute()
	assert.NotNil(t, err, "err should not be nil")
	assert.Equal(t, "breed name is required", err.Error())
}

func TestCommandGetPhotoInvalidCount(t *testing.T) {
	setup()
	command, err := NewCommand("get photo hound -1")
	assert.Nil(t, err, "err should be nil")

	_, err = command.Execute()
	assert.NotNil(t, err, "err should not be nil")
	assert.Equal(t,  "invalid count", err.Error())
}

func TestCommandGetPhotoInvalidParam(t *testing.T) {
	setup()
	command, err := NewCommand("get photo hound -1 2")
	assert.Nil(t, err, "err should be nil")

	_, err = command.Execute()
	assert.NotNil(t, err, "err should not be nil")
	assert.Equal(t, "too many params", err.Error())
}
