package main

import (
	"net/http"

	"github.com/koha90/project-driver/internal/config"
	"github.com/koha90/project-driver/pkg/logging"
)

func main() {
	log := logging.GetLogger()
	cfg := config.GetConfig()

	log.Print("start logger fine ", cfg.Addr)

	http.ListenAndServe(cfg.Port, nil)
}
