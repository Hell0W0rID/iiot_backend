//
//
//
//
// Unless required by applicable law or agreed to in writing, software

package interfaces

import (
	"context"
	"sync"
	"time"

	gometrics "github.com/rcrowley/go-metrics"
)

// MetricsManager manages a services metrics
type MetricsManager interface {
	// ResetInterval resets the interval between reporting the current metrics
	ResetInterval(interval time.Duration)
	// Register registers a go-metrics metric item such as a Counter
	Register(name string, item interface{}, tags map[string]string) error
	// IsRegistered checks whether a metric has been registered
	IsRegistered(name string) bool
	// Unregister unregisters a go-metrics metric item such as a Counter
	Unregister(name string)
	// Run starts the collection of metrics
	Run(ctx context.Context, wg *sync.WaitGroup)
	// GetCounter retrieves the specified registered Counter
	// Returns nil if named item not registered or not a Counter
	GetCounter(name string) gometrics.Counter
	// GetGauge retrieves the specified registered Gauge
	// Returns nil if named item not registered or not a Gauge
	GetGauge(name string) gometrics.Gauge
	// GetGaugeFloat64 retrieves the specified registered GaugeFloat64
	// Returns nil if named item not registered or not a GaugeFloat64
	GetGaugeFloat64(name string) gometrics.GaugeFloat64
	// GetTimer retrieves the specified registered Timer
	// Returns nil if named item not registered or not a Timer
	GetTimer(name string) gometrics.Timer
}

// MetricsReporter reports the metrics
type MetricsReporter interface {
	Report(registry gometrics.Registry, metricTags map[string]map[string]string) error
}
