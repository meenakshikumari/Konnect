package handlers

import (
	"api/contracts"
	"api/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetAllAPIServicesHandler(svc APISvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		page := query.Get("page")
		perPage := query.Get("per_page")
		filterNameContains := query.Get("filter[name][contains]")
		sort := query.Get("sort") // default will be updated_at: desc

		params := contracts.GetAllApiServiceRequestParams{
			Page:         page,
			PerPage:      perPage,
			FilterOnName: filterNameContains,
			Sort:         sort,
		}
		req, err := params.Validate()
		if err != nil {
			logrus.Errorf("[GetAllAPIServicesHandler]Failed to validate params: %s", err.Error())
			handleError(w, errors.NewMalformedRequestError("Failed to Parse the query params"))
			return
		}

		logrus.Infof("[GetAllAPIServicesHandler]REQUEST INITATED")

		resp := &contracts.ServicesDetails{}
		resp, err = svc.FindAllAPIServices(*req)
		if err != nil {
			logrus.Errorf("[GetAllAPIServicesHandler][FindAllAPIServices]Failed with request params: %v err: %s", req, err.Error())
			handleError(w, err)
		}

		writeResponse(w, http.StatusOK, contracts.GetAllAPIServicesResponse{
			Success: true,
			Details: *resp,
		})
	}
}
