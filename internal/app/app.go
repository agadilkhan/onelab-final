package app

import (
	"log"

	"github.com/agadilkhan/onelab-final/internal/config"
	"github.com/agadilkhan/onelab-final/internal/repository/pgrepo"
)

func Run(cfg *config.Config) error {
	_, err := pgrepo.New(
		pgrepo.WithHost(cfg.DB.Host),
		pgrepo.WithPort(cfg.DB.Port),
		pgrepo.WithDBName(cfg.DB.DBName),
		pgrepo.WithUsername(cfg.DB.Username),
		pgrepo.WithPassword(cfg.DB.Password),
	)
	if err != nil {
		log.Printf("connection to DB err: %s", err.Error())
		return err
	}
	log.Println("connection success")

	return nil
}