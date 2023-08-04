package main

import (
	"github.com/agadilkhan/onelab-final/internal/app"
	"github.com/agadilkhan/onelab-final/internal/config"
	"fmt"
)

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