package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap"
	"github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap/flags"
	"github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap/handlers"
	"github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap/interfaces"
	"github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap/startup"
	"github.com/edgexfoundry/go-mod-bootstrap/v3/di"

	"iiot-backend/services/core/data/application"
	"iiot-backend/services/core/data/container"
	"iiot-backend/pkg/common"
)

func main() {
	startupTimer := startup.NewStartUpTimer(common.CoreDataServiceKey)

	// All common command-line flags have been moved to DefaultCommonFlags. Service specific flags can be added here,
	// by inserting service specific flag prior to call to commonFlags.Parse().
	// Example:
	// flags.FlagSet.StringVar(&myvar, "m", "", "Specify a ....")
	// flags.Parse(os.Args[1:])
	flags.Parse(os.Args[1:])

	service, ok := handlers.NewServiceWithRequestValidation(common.CoreDataServiceKey, common.ConfigStemCore+common.CoreDataServiceKey)
	if !ok {
		os.Exit(1)
	}

	lc := service.LoggingClient()
	configuration := container.ConfigurationFrom(service.DIContainer())

	lc.Info("Starting " + common.CoreDataServiceKey + " " + common.ServiceVersion)

	// wire up dependencies
	if err := container.Container.Update(service.DIContainer()); err != nil {
		lc.Error(err.Error())
		os.Exit(1)
	}

	httpServer := handlers.NewHttpServer(service.RequestTimeout(), service.MaxRequestSize())

	applicationService := application.NewDataService()
	handler := handlers.NewDataHandler(applicationService, lc, configuration.Service.MaxResultCount)

	handlers.LoadRestRoutes(service.WebRouter(), httpServer, handler, lc)

	// Bootstrap will start the web server and run until it receives a signal to stop
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		startupTimer.SinceAsString()
		// run until the service is told to stop
		service.Run(ctx, cancel, wg, startupTimer)
	}()

	// wait for interrupt or termination signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// Block until signal is received
	<-c

	// start shutdown
	cancel()
	wg.Wait()
}