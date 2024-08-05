package service

import (
	"encoding/json"
	"golang-gin-company-api/internal/config"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type CompanyService struct {
	config *config.Config
}

func NewCompanyService(cfg *config.Config) *CompanyService {
	return &CompanyService{
		config: cfg,
	}
}

func (cs *CompanyService) GetData1(id string) (map[string]interface{}, error) {
	resp, err := http.Get(cs.config.Data1URL + id + "/json/")
	if err != nil {
		log.Errorf("Failed to make GET request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Failed to read response body: %v", err)
		return nil, err
	}
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Errorf("Failed to unmarshal response body: %v", err)
		return nil, err
	}
	log.Info("GET request successful")
	return response, nil
}

func (cs *CompanyService) GetData2(id string) (map[string]interface{}, error) {
	resp, err := http.Get(cs.config.Data2URL)
	if err != nil {
		log.Errorf("Failed to make GET request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Failed to read response body: %v", err)
		return nil, err
	}
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Errorf("Failed to unmarshal response body: %v", err)
		return nil, err
	}
	log.Info("GET request successful")
	return response, nil
}

// func (cs *CompanyService) CreateCompany(data map[string]interface{}) (map[string]interface{}, error) {
// 	response, err := cs.client.Post(cs.config.CreateURL, data)
// 	return response, err
// }
