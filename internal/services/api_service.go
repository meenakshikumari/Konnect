package services

import (
	"api/contracts"
	"api/internal/models"
	"github.com/sirupsen/logrus"
	"strconv"
)

type APIServiceRepository interface {
	FindByID(id int64) (*models.Service, error)
	FindAll(limit, offset int, sortOn, sortBy, nameContains string) ([]models.Service, error)
}

type APIService struct {
	ServiceRepo APIServiceRepository
}

func (a APIService) FindAllAPIServices(req contracts.GetAllApiServiceRequestVal) (*contracts.ServicesDetails, error) {
	services, err := a.ServiceRepo.FindAll(req.PerPage, req.Page, "updated_at", "desc", req.FilterOnName) // added sort val as default
	if err != nil {
		logrus.Errorf("[FindAllAPIServices][FindAll] Failed with error : %s", err.Error())
		return nil, err
	}

	var svc []contracts.APIService
	for _, service := range services {
		s := contracts.APIService{
			ID:           strconv.Itoa(int(service.ID)),
			Name:         service.Name,
			Description:  service.Description,
			Published:    service.Published,
			VersionCount: strconv.Itoa(service.VersionCount),
		}
		logrus.Infof("[FindAllAPIServices][FindAll] ServiceID:%s : %v", s.ID, s)
		svc = append(svc, s)
	}
	return &contracts.ServicesDetails{ServiceDetail: svc}, nil
}
func (a APIService) FindAPIServiceDetails(id int64) (*contracts.APIService, error) {
	s, err := a.ServiceRepo.FindByID(id)
	if err != nil {
		logrus.Errorf("[FindAllAPIServices][FindAll] Failed with error : %s", err.Error())
		return nil, err
	}

	resp := &contracts.APIService{
		ID:           strconv.Itoa(int(s.ID)),
		Name:         s.Name,
		Description:  s.Description,
		Published:    s.Published,
		VersionCount: strconv.Itoa(s.VersionCount),
	}
	return resp, nil
}
