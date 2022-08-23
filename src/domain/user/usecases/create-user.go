package user

import (
	user "github.com/felipefadoni/boilerplate-golang/src/domain/user/repositories"
	"github.com/felipefadoni/boilerplate-golang/src/dto"
	"github.com/felipefadoni/boilerplate-golang/src/helpers"
	"github.com/felipefadoni/boilerplate-golang/src/infra/postgres/entities"
)

func CreateUserUseCase(userDTO dto.CreateUserDTO) (dto.CreateUserReturnDTO, error) {
	userEntity := entities.User{}
	helpers.TransformData(userDTO, &userEntity)
	user, err := user.CreateUser(userEntity)
	if err != nil {
		return user, err
	}
	return user, nil
}
