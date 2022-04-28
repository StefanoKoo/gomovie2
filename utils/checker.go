package utils

import "net/mail"

func IsEmailValid(s string) bool {
	_, err := mail.ParseAddress(s)

	return err == nil
}
