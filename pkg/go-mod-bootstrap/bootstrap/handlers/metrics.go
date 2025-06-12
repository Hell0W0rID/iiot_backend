//
//
//
//
// Unless required by applicable law or agreed to in writing, software

package handlers

import (
	"context"
	"math"
	"sync"
	"time"

	"iiot-backend/pkg/go-mod-core-contracts/clients/logger"

	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/container"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/interfaces"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/metrics"
	"iiot-backend/pkg/go-mod-bootstrap/bootstrap/startup"
	"iiot-backend/pkg/go-mod-bootstrap/config"
	"iiot-backend/pkg/go-mod-bootstrap/di"
)

type RegisterTelemetryFunc func(logger.LoggerClient, *config.TelemetryInfo, interfaces.MetricsManager)

type ServiceMetrics struct {
	serviceName string
}

func NewServiceMetrics(serviceName string) *ServiceMetrics {
	return &ServiceMetrics{
		serviceName: serviceName,
	}
}

// BootstrapHandler fulfills the BootstrapHandler contract and performs initialization of service metrics.
func (s *ServiceMetrics) BootstrapHandler(ctx context.Context, wg *sync.WaitGroup, _ startup.Timer, dic *di.Container) bool {
	lc := container.LoggerClientFrom(dic.Get)
	serviceConfig := container.ConfigurationFrom(dic.Get)

	telemetryConfig := serviceConfig.GetTelemetryInfo()

	if telemetryConfig.Interval == "" {
		telemetryConfig.Interval = "0s"
	}

	interval, err := time.ParseDuration(telemetryConfig.Interval)
	if err != nil {
		lc.Errorf("Telemetry interval is invalid time duration: %s", err.Error())
		return false
	}

	if interval == 0 {
		lc.Infof("0 specified for metrics reporting interval. Setting to max duration to effectively disable reporting.")
		interval = math.MaxInt64
	}

	baseTopic := serviceConfig.GetBootstrap().MessageBus.GetBaseTopicPrefix()
	reporter := metrics.NewMessageBusReporter(lc, baseTopic, s.serviceName, dic, telemetryConfig)
	manager := metrics.NewManager(lc, interval, reporter)

	manager.Run(ctx, wg)

	dic.Update(di.ServiceConstructorMap{
		container.MetricsManagerInterfaceName: func(get di.Get) interface{} {
			return manager
		},
	})

	return true
}
