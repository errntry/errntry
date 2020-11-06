package i18n

import (
	"errors"
)

func Error(msg string) error {
	return errors.New(msg)
}
