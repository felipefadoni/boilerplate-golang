package user

import (
	user "github.com/felipefadoni/boilerplate-golang/src/domain/user/repositories"
	"github.com/felipefadoni/boilerplate-golang/src/dto"
	"github.com/felipefadoni/boilerplate-golang/src/helpers"
	"github.com/felipefadoni/boilerplate-golang/src/infra/postgres/entities"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserUseCase(userDTO dto.CreateUserDTO) (dto.CreateUserReturnDTO, error) {
	var userEntity = entities.User{}

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.CreateUserReturnDTO{}, err
	}

	helpers.TransformData(userDTO, &userEntity)

	userEntity.Permission = "[\"user::default\"]"
	userEntity.Password = string(hashedPassword)

	user, err := user.CreateUser(userEntity)

	if err != nil {
		return user, err
	}

	return user, nil
}
