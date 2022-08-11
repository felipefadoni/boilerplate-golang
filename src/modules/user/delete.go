package user

import "github.com/felipefadoni/boilerplate-golang/src/domain/user"

func DeleteUserService(id string) {
	user.DeleteUserRepository(id)
}
