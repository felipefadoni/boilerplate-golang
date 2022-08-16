package user

import (
	"github.com/felipefadoni/boilerplate-golang/src/domain/user"
	"github.com/felipefadoni/boilerplate-golang/src/dto"
	"github.com/felipefadoni/boilerplate-golang/src/infra/postgres/entities"
)

func CreateUserUseCase(userDTO dto.CreateUserDTO) (dto.CreateUserReturnDTO, error) {
	var userEntity = entities.User{
		Name:     userDTO.Name,
		Email:    userDTO.Email,
		Password: userDTO.Password,
		Login:    userDTO.Login,
		Accepted: userDTO.Accepted,
	}

	user, err := user.CreateUser(userEntity)
	if err != nil {
		return user, err
	}

	return user, nil
}
