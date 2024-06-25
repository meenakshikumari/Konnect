package handlers

import (
	"api/contracts"
	"api/internal/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"net/http"
)

func NewRouter(deps Dependencies) http.Handler {
	router := mux.NewRouter()
	router.Use(withRequestedAt)

	router.HandleFunc("/", indexRouterHandler).Methods(http.MethodGet)

	router.HandleFunc("/api/api-services", GetAllAPIServicesHandler(deps.APIService)).Methods(http.MethodGet)
	router.HandleFunc("/api/api-services/{service_id}", GetAPIServiceDetailsHandler(deps.APIService)).Methods(http.MethodGet)

	return http.HandlerFunc(router.ServeHTTP)
}

type Dependencies struct {
	Db             *sqlx.DB
	APIService     services.APIService
	ServiceVersion services.ServiceVersion
}

type APISvc interface {
	FindAllAPIServices(req contracts.GetAllApiServiceRequestVal) (*contracts.ServicesDetails, error)
	FindAPIServiceDetails(id int64) (*contracts.APIService, error)
}

func indexRouterHandler(wr http.ResponseWriter, req *http.Request) {
	logrus.Infof("Server is running")
	writeResponse(wr, http.StatusOK, struct{}{})
}

func writeResponse(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if v == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(v); err != nil {
		logrus.Errorf("failed to write response: %v", err)
	}
}
