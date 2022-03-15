package utils

import (
	"errors"
	"regexp"
	"strings"
)

func CheckID(id string) error {
	// remove the blank
	trimID := strings.Trim(id, " ")

	// id has a blank space
	if id != trimID || len(strings.Split(trimID, " ")) > 1 {
		return errors.New("id has a blank space")
	}

	// empty id
	if trimID == "" {
		return errors.New("please input id")
	}

	// various additional constraints can be set
	ok, _ := regexp.MatchString("^[-_A-Za-z0-9]+$", trimID)
	if !ok {
		return errors.New("invalid id")
	}
	return nil
}
