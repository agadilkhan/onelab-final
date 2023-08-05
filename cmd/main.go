package main

import (
	"github.com/agadilkhan/onelab-final/internal/app"
	"github.com/agadilkhan/onelab-final/internal/config"
	"fmt"
)

// @title           Blog API
// @version         0.0.1
// @description     API for Blog application

// @host      localhost:8080
// @BasePath  /

// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%#v", cfg))

	err = app.Run(cfg)
	if err != nil {
		panic(err)
	}
	
}