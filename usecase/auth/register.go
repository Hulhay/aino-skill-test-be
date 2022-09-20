package auth

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `jsno:"email"`
	Password string `json:"password"`
}

func (c *RegisterRequest) Validate() error {

	if err := validation.Validate(c.Name, validation.Required); err != nil {
		return errors.New("name must be filled")
	}

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
