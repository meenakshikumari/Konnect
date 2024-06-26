package contracts

import (
	"api/pkg/errors"
	"strconv"
	"strings"
)

type GetAllApiServiceRequestParams struct {
	Page         string
	PerPage      string
	FilterOnName string
	Sort         string
}

type SortValueEnum string

const (
	Ascending  SortValueEnum = "asc"
	Descending SortValueEnum = "desc"
)

type GetAllApiServiceRequestVal struct {
	Page         int
	PerPage      int
	FilterOnName string
	SortParam    string
	SortValue    SortValueEnum
}

func (r *GetAllApiServiceRequestParams) Validate() (*GetAllApiServiceRequestVal, error) {
	req := &GetAllApiServiceRequestVal{}
	if r.PerPage == "" {
		return nil, errors.NewMissingFieldError("per_page is required")
	}
	perPage, err := strconv.Atoi(r.PerPage)
	if err != nil || perPage < 1 {
		return nil, errors.NewInvalidValueError(r.PerPage, "per_page value is not correct. It should be integer value")
	}
	req.PerPage = perPage

	if r.Sort == "" {
		req.SortParam = "updated_at"
		req.SortValue = Descending
	} else {
		s := strings.Split(r.Sort, ":")
		req.SortParam = s[0]
		req.SortValue = SortValueEnum(s[1])
	}

	if r.Page == "" && r.FilterOnName == "" {
		return nil, errors.NewMissingFieldError("Page Num or filter_on_name is missing")
	}

	if r.Page != "" {
		page, err := strconv.Atoi(r.Page)
		if err != nil || page < 1 {
			return nil, errors.NewInvalidValueError(r.Page, "page number value is not correct. It should be integer value")
		}
		req.Page = page
	} else if r.Page == "" {
		req.Page = 1
	}

	req.FilterOnName = r.FilterOnName

	return req, nil
}
