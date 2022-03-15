package utils

import (
	"errors"
	"strings"
)

func CheckSecret(secret string) error {
	// remove the blank
	trimSecret := strings.Trim(secret, " ")

	// secret has blank
	if secret != trimSecret || len(strings.Split(trimSecret, " ")) > 1 {
		return errors.New("secret has a blank space")
	}

	// empty secret
	if trimSecret == "" {
		return errors.New("please input secret")
	}

	// various additional constraints can be set

	return nil
}
