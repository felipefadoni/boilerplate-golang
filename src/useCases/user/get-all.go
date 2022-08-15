package user

import (
	"github.com/felipefadoni/boilerplate-golang/src/domain/user"
	"github.com/felipefadoni/boilerplate-golang/src/dto"
)

func GetAllUserModule() []dto.GetAllUserDTO {
	users := user.GetAll()
	return users
}
