package user

import (
	"github.com/felipefadoni/boilerplate-golang/src/domain/user"
	"github.com/felipefadoni/boilerplate-golang/src/dto"
	"github.com/felipefadoni/boilerplate-golang/src/helpers"
)

type GetAllUserModuleReturn struct {
	Rows         []dto.GetAllUserDTO `json:"rows"`
	Total        int64               `json:"total"`
	CurrentPage  int64               `json:"currentPage"`
	TotalPages   int64               `json:"totalPages"`
	NextPage     int64               `json:"nextPage"`
	PreviousPage int64               `json:"previousPage"`
}

func GetAllUserModule(page int64, limit int64) GetAllUserModuleReturn {

	var result GetAllUserModuleReturn

	result.Rows = user.GetAll(page, limit)
	result.Total = user.GetTotal()

	var pagination = helpers.Pagination(result.Total, page, limit)

	result.CurrentPage = page
	result.TotalPages = pagination["totalPages"]
	result.NextPage = pagination["nextPage"]
	result.PreviousPage = pagination["previousPage"]

	return result
}
