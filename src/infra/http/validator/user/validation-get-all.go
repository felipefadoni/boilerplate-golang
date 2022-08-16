package user

import (
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type getAllUserValidation struct {
	Page  string `validate:"required"`
	Limit string `validate:"required"`
}

func ValidationGetAllUser(pageQ string, limitQ string) (page int64, limit int64, e error) {

	validationParams := getAllUserValidation{
		Page:  pageQ,
		Limit: limitQ,
	}

	validate := validator.New()
	err := validate.Struct(validationParams)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return 0, 0, err
		}
		for _, err := range err.(validator.ValidationErrors) {
			return 0, 0, errors.New(err.Field() + " is " + err.Tag())
		}
		return 0, 0, err
	}

	page, erP := strconv.ParseInt(pageQ, 10, 64)
	if erP != nil {
		return 0, 0, errors.New("PAGE must be a number")
	}

	limit, erL := strconv.ParseInt(limitQ, 10, 64)
	if erL != nil {
		return 0, 0, errors.New("LIMIT must be a number")
	}

	return page, limit, err
}
