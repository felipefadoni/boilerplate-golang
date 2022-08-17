package user

import "github.com/felipefadoni/boilerplate-golang/src/domain/user"

func DeleteUserUseCase(id string) error {
	err := user.DeleteUser(id)
	return err
}
