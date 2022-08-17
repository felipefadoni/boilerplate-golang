package user

import (
	"github.com/felipefadoni/boilerplate-golang/src/infra/postgres"
	"github.com/felipefadoni/boilerplate-golang/src/infra/postgres/entities"
)

func DeleteUser(id string) error {
	db := postgres.GetInstance()
	var User entities.User
	err := db.Where("id = ?", id).Delete(&User).Error
	return err
}
