package main

import (
	"golang-gin-company-api/internal/config"
	"golang-gin-company-api/internal/controller"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetFormatter(&log.JSONFormatter{})

	cfg := config.LoadConfig()

	r := gin.Default()

	r.Use(config.Logger())

	r.GET("/conta/:id", controller.NewCompanyController(cfg).GetData1)

	log.Info("Iniciando o servidor na porta" + cfg.Port)
	if err := r.Run(cfg.Port); err != nil {
		log.Fatalf("Falha ao iniciar o servidor: %v", err)
	}
}
