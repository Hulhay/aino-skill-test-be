package auth

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type LoginRequest struct {
	Email    string `jsno:"email"`
	Password string `json:"password"`
}

func (c *LoginRequest) Validate() error {

	if err := validation.Validate(c.Email, validation.Required); err != nil {
		return errors.New("email must be filled")
	}

	if err := validation.Validate(c.Email, is.Email); err != nil {
		return errors.New("invalid email format")
	}

	if err := validation.Validate(c.Password, validation.Required); err != nil {
		return errors.New("password must be filled")
	}

	if err := validation.Validate(c.Password, validation.Length(6, 0)); err != nil {
		return errors.New("password minimal 6 character")
	}

	return nil
}
