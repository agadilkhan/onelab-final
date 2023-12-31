package app

import (
	"log"
	"os"
	"os/signal"

	"github.com/agadilkhan/onelab-final/internal/config"
	"github.com/agadilkhan/onelab-final/internal/handler"
	"github.com/agadilkhan/onelab-final/internal/repository/pgrepo"
	"github.com/agadilkhan/onelab-final/internal/service"
	"github.com/agadilkhan/onelab-final/pkg/httpserver"
	"github.com/agadilkhan/onelab-final/pkg/jwttoken"
)

func Run(cfg *config.Config) error {
	db, err := pgrepo.New(
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

	token := jwttoken.New(cfg.Token.SecretKey)
	repo := pgrepo.NewPostgresRepository(db)
	srvs := service.New(*repo, cfg, token)
	hndlr := handler.New(srvs)
	server := httpserver.New(
		hndlr.InitRouter(),
		httpserver.WithPort(cfg.HTTP.Port),
		httpserver.WithReadTimeout(cfg.HTTP.ReadTimeout),
		httpserver.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		httpserver.WithShutdownTimeout(cfg.HTTP.ShutdownTimeout),
	)

	log.Println("server started")
	server.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err = <-server.Notify():
		log.Printf("server notify: %s", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown err: %s", err)
	}

	return nil
}
