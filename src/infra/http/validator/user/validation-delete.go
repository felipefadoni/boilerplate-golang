package user

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidationDeleteUser(id string) error {

	validate := validator.New()

	err := validate.Var(id, "required,uuid4")
	if err != nil {
		return errors.New("ID is required or invalid")
	}
	return err
}
