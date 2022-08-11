package user

import (
	"github.com/felipefadoni/boilerplate-golang/src/dto"
	"github.com/felipefadoni/boilerplate-golang/src/infra/postgres"
)

func GetAll() []dto.GetAllUserDTO {
	db := postgres.GetInstance()

	var result []dto.GetAllUserDTO

	db.Raw("SELECT * FROM users").Scan(&result)

	return result
}
