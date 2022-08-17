package user

import (
	"errors"

	"github.com/felipefadoni/boilerplate-golang/src/dto"
	"github.com/felipefadoni/boilerplate-golang/src/helpers"
	"github.com/go-playground/validator/v10"
)

type createUserValidation struct {
	Name     string `validate:"required,min=3,max=200"`
	Email    string `validate:"required,email"`
	Login    string `validate:"required,min=3,max=200"`
	Password string `validate:"required"`
	Accepted bool   `validate:"required"`
}

func ValidationCreateUser(userDTO dto.CreateUserDTO) error {
	validate := validator.New()
	userValidation := createUserValidation{}
	helpers.TransformData(userDTO, &userValidation)
	err := validate.Struct(userValidation)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Field() + " is " + err.Tag())
		}
		return err
	}
	return err
}
