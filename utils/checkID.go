package utils

import (
	"errors"
	"strings"
)

func CheckID(id string) error {
	// remove the blank
	trimID := strings.Trim(id, " ")

	// id has a blank space
	if id != trimID {
		return errors.New("id has a blank space")
	}

	// empty id
	if trimID == "" {
		return errors.New("please input id")
	}

	// various additional constraints can be set

	return nil
}
