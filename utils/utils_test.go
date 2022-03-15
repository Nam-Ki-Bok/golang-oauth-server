package utils

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGenerateSHA256(t *testing.T) {
	assert.New(t)
	type testCase struct {
		str    string
		output string
	}

	cases := []testCase{
		{str: "hello", output: "2CF24DBA5FB0A30E26E83B2AC5B9E29E1B161E5C1FA7425E73043362938B9824"},
		{str: "world", output: "486EA46224D1BB4FB680F34F7C9AD96A8F24EC88BE73EA8E5A6C65260E9CB8A7"},
	}

	for _, tc := range cases {
		output := GenerateSHA256(tc.str)
		assert.Equal(t, tc.output, output)
	}
}

func TestCheckID(t *testing.T) {
	assert.New(t)
	type testCase struct {
		id  string
		err error
	}

	cases := []testCase{
		{id: "user_client", err: nil},
		{id: "middle blank space", err: errors.New("id has a blank space")},
		{id: " front_blank_space", err: errors.New("id has a blank space")},
		{id: "back_blank_space ", err: errors.New("id has a blank space")},
		{id: "", err: errors.New("please input id")},
	}

	for _, tc := range cases {
		err := CheckID(tc.id)
		assert.Equal(t, tc.err, err)
	}
}

func TestCheckSecret(t *testing.T) {
	assert.New(t)
	type testCase struct {
		secret string
		err    error
	}

	cases := []testCase{
		{secret: "user_secret", err: nil},
		{secret: "middle blank space", err: errors.New("secret has a blank space")},
		{secret: " front_blank_space", err: errors.New("secret has a blank space")},
		{secret: "back_blank_space ", err: errors.New("secret has a blank space")},
		{secret: "", err: errors.New("please input secret")},
	}

	for _, tc := range cases {
		err := CheckSecret(tc.secret)
		assert.Equal(t, tc.err, err)
	}
}

func TestReturnError(t *testing.T) {
	assert.New(t)
	assert.Panics(t, func() { ReturnError(http.StatusUnauthorized, errors.New("invalid scope")) })
}
