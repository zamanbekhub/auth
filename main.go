package main

import (
	_ "auth/docs"
	"auth/internal/app"
	"auth/internal/config"
)

//go:generate go run github.com/swaggo/swag/cmd/swag init

// @title auth
// @версия 1.0.0
// @description auth
//
// @host 127.0.0.1:8000
// @BasePath /auth
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
