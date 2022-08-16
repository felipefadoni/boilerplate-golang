package user

import "testing"

func Test_ValidationGetAllUser(t *testing.T) {
	_, _, err := ValidationGetAllUser("1", "1")
	if err != nil {
		t.Errorf("ValidationGetAllUser() = %v; want nil", err)
		return
	}
}

func Test_ErrorParsePage(t *testing.T) {
	_, _, err := ValidationGetAllUser("a", "1")
	if err.Error() != "PAGE must be a number" {
		t.Errorf("ValidationGetAllUser() = %v; want PAGE must be a number", err)
		return
	}
}

func Test_ErrorParseLimit(t *testing.T) {
	_, _, err := ValidationGetAllUser("1", "a")
	if err.Error() != "LIMIT must be a number" {
		t.Errorf("ValidationGetAllUser() = %v; want LIMIT must be a number", err)
		return
	}
}

func Test_ErrorRequiredPage(t *testing.T) {
	_, _, err := ValidationGetAllUser("", "1")
	if err.Error() != "Page is required" {
		t.Errorf("ValidationGetAllUser() = %v; want Page is required", err)
		return
	}
}

func Test_ErrorRequiredLimit(t *testing.T) {
	_, _, err := ValidationGetAllUser("1", "")
	if err.Error() != "Limit is required" {
		t.Errorf("ValidationGetAllUser() = %v; want Limit is required", err)
		return
	}
}
