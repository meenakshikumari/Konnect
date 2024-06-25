package handlers

import (
	"api/contracts"
	"api/pkg/errors"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func GetAPIServiceDetailsHandler(svc APISvc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		serviceID := params["service_id"]
		id, err := strconv.Atoi(serviceID)
		if err != nil {
			logrus.Errorf("[GetAPIServiceDetailsHandler]Failed to validate serviceID: %s", err.Error())
			handleError(w, errors.NewMalformedRequestError("Failed to Parse the query params"))
			return
		}

		logrus.Infof("[GetAPIServiceDetailsHandler]REQUEST INITATED")

		resp, err := svc.FindAPIServiceDetails(int64(id))
		if err != nil {
			logrus.Errorf("[GetAPIServiceDetailsHandler][FindAPIServiceDetails]Failed with serviceID params: %v err: %s", id, err.Error())
			handleError(w, err)
		}

		writeResponse(w, http.StatusOK, contracts.GetAPIServiceResponse{
			Success: true,
			Details: *resp,
		})
	}
}
