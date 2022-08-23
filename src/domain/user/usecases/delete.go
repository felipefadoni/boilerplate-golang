package user

import user "github.com/felipefadoni/boilerplate-golang/src/domain/user/repositories"

func DeleteUserUseCase(id string) error {
	err := user.DeleteUser(id)
	return err
}
