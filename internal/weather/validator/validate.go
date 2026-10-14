package validator

import (
	"errors"
	"strings"
)

func ValidateCity(city string) error {
	if strings.TrimSpace(city) == "" {
		return errors.New("city is not specified")
	}

	return nil
}
