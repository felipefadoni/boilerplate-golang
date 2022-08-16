package user

import (
	"testing"

	"github.com/felipefadoni/boilerplate-golang/src/dto"
)

func Test_ValidationCreateUser(t *testing.T) {
	userDTO := dto.CreateUserDTO{
		Name:     "Felipe Fadoni",
		Email:    "felipefadonimt@gmail.com",
		Login:    "felipefadoni",
		Password: "123456",
		Accepted: true,
	}

	err := ValidationCreateUser(userDTO)
	if err != nil {
		t.Errorf("ValidationCreateUser() = %v; want nil", err)
		return
	}
}

func Test_ValidationMinNameUser(t *testing.T) {
	userDTO := dto.CreateUserDTO{
		Name:     "Fe",
		Email:    "felipefadonimt@gmail.com",
		Login:    "felipefadoni",
		Password: "123456",
		Accepted: true,
	}

	err := ValidationCreateUser(userDTO)
	if err.Error() != "Name is min" {
		t.Errorf("ValidationCreateUser() = %v; want Name is min", err)
		return
	}
}
