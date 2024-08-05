package service_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"golang-gin-company-api/internal/config"
	"golang-gin-company-api/internal/service"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetData1(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Dados esperados
	expectedData := map[string]interface{}{"key": "value"}
	responseBody, err := json.Marshal(expectedData)
	if err != nil {
		t.Fatalf("Failed to marshal expected data: %v", err)
	}

	// Configuração do mock
	url := "http://example.com/api/123/json/"
	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(http.StatusOK, string(responseBody)))

	// Instanciação do serviço
	cfg := &config.Config{Data1URL: "http://example.com/api/"}
	cs := service.NewCompanyService(cfg)

	// Chamada ao método
	data, err := cs.GetData1("123")

	// Verificações
	if err != nil {
		t.Fatalf("GetData1 failed: %v", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedData, data)
}
