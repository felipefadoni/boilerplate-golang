package user

import (
	"errors"

	"github.com/felipefadoni/boilerplate-golang/src/dto"
	"github.com/felipefadoni/boilerplate-golang/src/infra/postgres"
	"github.com/felipefadoni/boilerplate-golang/src/infra/postgres/entities"
)

func CreateUser(user entities.User) (dto.CreateUserReturnDTO, error) {
	db := postgres.GetInstance()

	userReturnDTO := dto.CreateUserReturnDTO{}

	result := db.Create(&user).Scan(&userReturnDTO)
	if result.Error != nil {
		return dto.CreateUserReturnDTO{}, errors.New("ERROR creating user")
	}
	return userReturnDTO, nil
}
