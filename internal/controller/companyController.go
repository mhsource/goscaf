package controller

import (
	"golang-gin-company-api/internal/config"
	"golang-gin-company-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyController struct {
	service *service.CompanyService
}

func NewCompanyController(cfg *config.Config) *CompanyController {
	return &CompanyController{
		service: service.NewCompanyService(cfg),
	}
}

func (cc *CompanyController) GetData1(c *gin.Context) {
	id := c.Param("id")
	response, err := cc.service.GetData1(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve data"})
		return
	}
	c.JSON(http.StatusOK, response)
}
