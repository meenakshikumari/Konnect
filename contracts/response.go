package contracts

import "api/pkg/errors"

type ErrorResponse struct {
	Data    interface{}           `json:"data"`
	Success bool                  `json:"success"`
	Errors  []errors.GenericError `json:"errors"`
}

type GetAllAPIServicesResponse struct {
	Details ServicesDetails       `json:"data"`
	Success bool                  `json:"success"`
	Errors  []errors.GenericError `json:"errors"`
}

type ServicesDetails struct {
	ServiceDetail []APIService `json:"service_detail"`
}

type APIService struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Published    bool   `json:"published"`
	VersionCount string `json:"version_count"`
}

type GetAPIServiceResponse struct {
	Details APIService `json:"data"`
	Success bool       `json:"success"`
}
