package user

import (
	"github.com/felipefadoni/boilerplate-golang/src/dto"
	"github.com/felipefadoni/boilerplate-golang/src/infra/postgres"
)

func GetAll(page int64, limit int64) []dto.GetAllUserDTO {
	db := postgres.GetInstance()

	var result []dto.GetAllUserDTO

	var offset int64
	if page <= 0 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	db.Raw("SELECT * FROM users WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT ? OFFSET ?", limit, offset).Scan(&result)

	return result
}
