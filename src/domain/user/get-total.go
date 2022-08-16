package user

import (
	"github.com/felipefadoni/boilerplate-golang/src/infra/postgres"
)

func GetTotal() int64 {
	db := postgres.GetInstance()

	var total int64

	db.Raw(`SELECT count(users.id) as total FROM users WHERE deleted_at IS NULL`).Scan(&total)

	return total
}
