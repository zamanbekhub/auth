package main

import (
	_ "service/docs"
	"service/internal/app"
	"service/internal/config"
)

//go:generate go run github.com/swaggo/swag/cmd/swag init

// @title service
// @версия 1.0.0
// @description service
//
// @host 127.0.0.1:8000
// @BasePath /service
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
