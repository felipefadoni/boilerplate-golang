package user

import "github.com/felipefadoni/boilerplate-golang/src/domain/user"

func DeleteUserService(id string) error {
	err := user.DeleteUserRepository(id)
	return err
}
