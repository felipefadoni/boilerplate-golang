package helpers

import "math"

func Pagination(total int64, page int64, limit int64) map[string]int64 {
	var result = make(map[string]int64)

	var previousPage int64
	var nextPage int64

	if page > 1 {
		previousPage = int64(page - 1)
	} else {
		previousPage = 0
	}

	if page < int64(math.Ceil(float64(total)/float64(limit))) {
		nextPage = int64(page + 1)
	} else {
		nextPage = 0
	}

	result["totalPages"] = int64(math.Ceil(float64(total) / float64(limit)))
	result["nextPage"] = nextPage
	result["previousPage"] = previousPage

	return result
}
