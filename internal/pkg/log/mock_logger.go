package log

import (
	"fmt"
	"sync"
)

type MockLogger struct {
	lock  *sync.RWMutex
	trace []string
	debug []string
	info  []string
	error []string
	warn  []string
	fatal []string
}

func NewMockLogger() *MockLogger {
	return &MockLogger{
		new(sync.RWMutex),
		make([]string, 0),
		make([]string, 0),
		make([]string, 0),
		make([]string, 0),
		make([]string, 0),
		make([]string, 0),
	}
}

func (m *MockLogger) Trace(message string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.trace = append(m.trace, message)
}

func (m *MockLogger) Debug(message string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.debug = append(m.debug, message)
}

func (m *MockLogger) Info(message string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.info = append(m.info, message)
}

func (m *MockLogger) Warn(message string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.warn = append(m.warn, message)
}

func (m *MockLogger) Error(message string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.error = append(m.error, message)
}

func (m *MockLogger) Fatal(message string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.fatal = append(m.fatal, message)
}

func (m *MockLogger) TraceF(formatString string, params ...interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.trace = append(m.trace, fmt.Sprintf(formatString, params...))
}

func (m *MockLogger) DebugF(formatString string, params ...interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.debug = append(m.debug, fmt.Sprintf(formatString, params...))
}

func (m *MockLogger) InfoF(formatString string, params ...interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.info = append(m.info, fmt.Sprintf(formatString, params...))
}

func (m *MockLogger) WarnF(formatString string, params ...interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.warn = append(m.warn, fmt.Sprintf(formatString, params...))
}

func (m *MockLogger) ErrorF(formatString string, params ...interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.error = append(m.error, fmt.Sprintf(formatString, params...))
}

func (m *MockLogger) FatalF(formatString string, params ...interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.fatal = append(m.fatal, fmt.Sprintf(formatString, params...))
}

func (m *MockLogger) TraceLogs() []string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.trace
}

func (m *MockLogger) InfoLogs() []string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.info
}

func (m *MockLogger) ErrorLogs() []string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.error
}

func (m *MockLogger) WarnLogs() []string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.warn
}

func (m *MockLogger) FatalLogs() []string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.fatal
}

func (m *MockLogger) DebugLogs() []string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.debug
}
