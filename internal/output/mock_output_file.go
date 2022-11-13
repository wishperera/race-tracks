package output

// MockOutputFile : implements the io.Writer interface for mocking a file write operation for testing purposes
type MockOutputFile struct {
	expectedException error
	data              []byte
}

// NewMockOutputFile : returns a new instance of MockOutputFile
func NewMockOutputFile(options ...Option) *MockOutputFile {
	mof := &MockOutputFile{
		data: make([]byte, 0),
	}

	for _, opt := range options {
		opt(mof)
	}

	return mof
}

// Option : optional parameters for constructor
type Option func(file *MockOutputFile)

// WithException : use this option when the MockOutputFile must return an exception on Write
func WithException(exception error) Option {
	return func(file *MockOutputFile) {
		file.expectedException = exception
	}
}

// Write : simulates a file write operation
func (m *MockOutputFile) Write(p []byte) (n int, err error) {
	if m.expectedException != nil {
		return 0, m.expectedException
	}

	n = 0
	if c := len(p); c > 0 {
		for n < c {
			m.writeByte(p[n])
			n++
		}
	}

	return n, nil
}

// Data : returns the data written to the MockOutputFile
func (m *MockOutputFile) Data() []byte {
	return m.data
}

func (m *MockOutputFile) writeByte(in byte) {
	m.data = append(m.data, in)
}
