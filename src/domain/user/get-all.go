package user

import (
	"github.com/felipefadoni/boilerplate-golang/src/dto"
	"github.com/felipefadoni/boilerplate-golang/src/infra/postgres"
)

func GetAll() []dto.GetAllUserDTO {
	db := postgres.GetInstance()

	var result []dto.GetAllUserDTO

	db.Raw("SELECT * FROM users WHERE deleted_at IS NULL").Scan(&result)

	return result
}
