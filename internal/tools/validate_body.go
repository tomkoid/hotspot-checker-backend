package tools

import (
	"bytes"
	"errors"

	"github.com/labstack/echo/v4"
)

func ValidateBody(c echo.Context, secretPassword string) error {
	if c.Request().Body == nil {
		return errors.New("Please send a request body")
	}

	body := new(bytes.Buffer)

	_, err := body.ReadFrom(c.Request().Body)
	if err != nil {
		return errors.New("Please send a request body")
	}

	if body.String() == secretPassword {
		return nil
	} else {
		return errors.New("Invalid password")
	}
}
