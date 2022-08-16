package user

import (
	"testing"
)

func Test_ValidationDeleteUser(t *testing.T) {
	id := "c34b6aaf-69f2-4ddb-ae7e-ec15a74f4e08"

	err := ValidationDeleteUser(id)
	if err != nil {
		t.Errorf("ValidationDeleteUser() = %v; want nil", err)
		return
	}
}

func Test_ErrorValidationUUID(t *testing.T) {
	id := "c34b6aaf-69f2-4ddb-ae7e-ec15a74f4e0"

	err := ValidationDeleteUser(id)
	if err.Error() != "ID is required or invalid" {
		t.Errorf("ValidationDeleteUser() = %v; want ID is required or invalid", err)
		return
	}
}
