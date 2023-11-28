package appError

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Message(err error) string {
	var messages []string
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				messages = append(messages, fmt.Sprintf("%s failed on the '%s' tag", e.Field(), e.ActualTag()))
			}
		} else {
			messages = append(messages, err.Error())
		}
	}

	return strings.Join(messages, "\n")
}
