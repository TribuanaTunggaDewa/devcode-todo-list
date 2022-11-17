package util

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	err := cv.Validator.Struct(i)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		msgs := make([]string, 0, len(errs))
		for _, fieldError := range err.(validator.ValidationErrors) {
			var msg string
			if fieldError.Param() != "" {
				msg := fmt.Sprintf(`%s cannot be null`, ToSnakeCase(fieldError.Field()))
				msgs = append(msgs, msg)
				continue
			}
			msg = fmt.Sprintf(`%s cannot be null`, ToSnakeCase(fieldError.Field()))
			msgs = append(msgs, msg)
		}
		erroMsg := ""
		for i, v := range msgs {
			if len(msgs) > 1 {
				if i == len(msgs)-1 {
					erroMsg = erroMsg + v
					continue
				}
				erroMsg = erroMsg + v + " and "
			} else {
				erroMsg = erroMsg + v
			}
		}
		return errors.New(string(erroMsg))
	}

	return nil
}

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
