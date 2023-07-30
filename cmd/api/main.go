package main

import (
	"github.com/koha90/project-driver/internal/config"
	"github.com/koha90/project-driver/internal/server"
	"github.com/koha90/project-driver/pkg/logging"
	"github.com/koha90/project-driver/storage"
)

func main() {
	log := logging.GetLogger()
	cfg := config.GetConfig()

	storage, err := storage.NewPostgresStorage(
		cfg.StorageConfig.Username,
		cfg.StorageConfig.Database,
		cfg.StorageConfig.Password,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("start storage initilazation")
	if err := storage.Init(); err != nil {
		log.Fatal(err)
	}
	log.Debug("finish. Next step")

	server := server.NewAPIServer(cfg.Addr, *storage)

	log.Infof("start server on port: %s", cfg.Addr)

	server.Run()
}
