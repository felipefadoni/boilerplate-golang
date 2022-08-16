package helpers

import (
	"testing"
)

func TestPagination_return_pagination(t *testing.T) {
	total := int64(100)
	page := int64(1)
	limit := int64(10)

	result := Pagination(total, page, limit)

	if result["totalPages"] != 10 {
		t.Errorf("Pagination() = %v; want 10", result["totalPages"])
		return
	}
}

func TestPagination_return_pagination_previous(t *testing.T) {
	total := int64(100)
	page := int64(2)
	limit := int64(10)

	result := Pagination(total, page, limit)

	if result["totalPages"] != 10 {
		t.Errorf("Pagination() = %v; want 10", result["totalPages"])
		return
	}

	if result["previousPage"] != 1 {
		t.Errorf("Pagination() = %v; want 1", result["previousPage"])
		return
	}
}

func TestPagination_return_pagination_next(t *testing.T) {
	total := int64(1)
	page := int64(1)
	limit := int64(10)

	result := Pagination(total, page, limit)

	if result["totalPages"] != 1 {
		t.Errorf("Pagination() = %v; want 1", result["totalPages"])
		return
	}

	if result["nextPage"] != 0 {
		t.Errorf("Pagination() = %v; want 0", result["nextPage"])
		return
	}
}
