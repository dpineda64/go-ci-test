package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReturnsMessage(t *testing.T) {
	username, err := ReturnsMessage("dp27664")

	if err != nil {
		t.Fatal(err)
	}
	print(username)
}

func TestReturnsMessageError(t *testing.T) {
	_, err := ReturnsMessage("")

	expectedError := errors.New("missing name")
	require.Equal(t, err.Error(), expectedError.Error())
}
