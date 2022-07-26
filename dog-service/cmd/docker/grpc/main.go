package main

import (
	"dog-service/util/config"
	"dog-service/util/logging"
	"dog-service/util/signal"
	"fmt"
	"net"
)

func setup() {
	config.SetupConfig()
	signal.SetupSignals()
	logging.SetupLogging(config.GetString("log.level"))
}

//initialise - server and services
func initialise() {
	initGRPCServices()
	initGRPCServer()
}

func run() {
	signal.CleanupOnSignal(cleanUp)
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GetString("server.grpcPort")))
	if err != nil {
		logging.Fatal("Something went wrong with the service")
	}
	err = server.Serve(l)
	if err != nil {
		logging.Fatal("Something went wrong with the service")
	}
}

func main() {
	setup()
	initialise()
	run()
}

func cleanUp() {
}
