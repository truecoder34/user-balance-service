/*
	Package to format error messages in readable style
*/
package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "nickname") {
		return errors.New("Nickname is already taken")
	}
	if strings.Contains(err, "email") {
		return errors.New("Email is already taken")
	}
	if strings.Contains(err, "phone_number") {
		return errors.New("Phone number is already taken")
	}
	return errors.New("Incorrect details")
}
