package main

import (
	"github.com/agadilkhan/onelab-final/internal/app"
	"github.com/agadilkhan/onelab-final/internal/config"
)

func main() {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	err = app.Run(cfg)
	if err != nil {
		panic(err)
	}
	
}