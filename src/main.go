package main

import (
	"log"

	"laches1sm/demo-parrot-service/adapters"
	"laches1sm/demo-parrot-service/infrastructure"
	"laches1sm/demo-parrot-service/services/httpserver"
)

const (
	serviceRunning = `Parrot Service HTTP server running...`
)

func main() {
	logger := log.Logger()
	parrotInfra := infrastructure.ParrotInfra
	parrotAdapter := adapters.NewParrotHTTPAdapter(logger, parrotInfra)

	server := httpserver.NewParrotServer(logger, parrotAdapter)
	server.SetupRoutes()

	logger.Println(serviceRunning)
	if err := server.Start(httpserver.ServerPort); err != nil {
		logger.Println(err.Error())
	}
}
