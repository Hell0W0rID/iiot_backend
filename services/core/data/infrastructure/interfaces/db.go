package interfaces

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/models"
)

// DBClient defines the interface for database operations
type DBClient interface {
	// Event operations
	AddEvent(event models.Event) (string, errors.EdgeX)
	AllEvents(offset int, limit int) ([]models.Event, uint32, errors.EdgeX)
	EventById(id string) (models.Event, errors.EdgeX)
	EventsByDeviceName(offset int, limit int, deviceName string) ([]models.Event, uint32, errors.EdgeX)
	EventsByTimeRange(start int64, end int64, offset int, limit int) ([]models.Event, uint32, errors.EdgeX)
	DeleteEventById(id string) errors.EdgeX
	DeleteEventsByDeviceName(deviceName string) errors.EdgeX
	DeleteEventsByAge(age int64) errors.EdgeX

	// Reading operations
	AddReading(reading models.Reading) (string, errors.EdgeX)
	AllReadings(offset int, limit int) ([]models.Reading, uint32, errors.EdgeX)
	ReadingById(id string) (models.Reading, errors.EdgeX)
	ReadingsByDeviceName(offset int, limit int, deviceName string) ([]models.Reading, uint32, errors.EdgeX)
	ReadingsByResourceName(offset int, limit int, resourceName string) ([]models.Reading, uint32, errors.EdgeX)
	ReadingsByTimeRange(start int64, end int64, offset int, limit int) ([]models.Reading, uint32, errors.EdgeX)
	DeleteReadingById(id string) errors.EdgeX
	DeleteReadingsByDeviceName(deviceName string) errors.EdgeX
	DeleteReadingsByResourceName(resourceName string) errors.EdgeX
	DeleteReadingsByAge(age int64) errors.EdgeX
}