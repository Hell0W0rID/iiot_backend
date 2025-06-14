/*******************************************************************************
 *******************************************************************************/

package logger

// MockLogger is a type that can be used for mocking the LoggerClient interface during unit tests
type MockLogger struct {
}

// NewMockClient creates a mock instance of LoggerClient
func NewMockClient() LoggerClient {
	return MockLogger{}
}

// SetLogLevel simulates setting a log severity level
func (lc MockLogger) SetLogLevel(_ string) error {
	return nil
}

// LogLevel returns the current log level setting
func (lc MockLogger) LogLevel() string {
	return ""
}

// Info simulates logging an entry at the INFO severity level
func (lc MockLogger) Info(_ string, _ ...interface{}) {
}

// Debug simulates logging an entry at the DEBUG severity level
func (lc MockLogger) Debug(_ string, _ ...interface{}) {
}

// Error simulates logging an entry at the ERROR severity level
func (lc MockLogger) Error(_ string, _ ...interface{}) {
}

// Trace simulates logging an entry at the TRACE severity level
func (lc MockLogger) Trace(_ string, _ ...interface{}) {
}

// Warn simulates logging an entry at the WARN severity level
func (lc MockLogger) Warn(_ string, _ ...interface{}) {
}

// Infof simulates logging an formatted message at the INFO severity level
func (lc MockLogger) Infof(_ string, _ ...interface{}) {
}

// Debugf simulates logging an formatted message at the DEBUG severity level
func (lc MockLogger) Debugf(_ string, _ ...interface{}) {
}

// Errorf simulates logging an formatted message at the ERROR severity level
func (lc MockLogger) Errorf(_ string, _ ...interface{}) {
}

// Tracef simulates logging an formatted message at the TRACE severity level
func (lc MockLogger) Tracef(_ string, _ ...interface{}) {
}

// Warnf simulates logging an formatted message at the WARN severity level
func (lc MockLogger) Warnf(_ string, _ ...interface{}) {
}
