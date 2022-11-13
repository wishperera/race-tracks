package input

import (
	"io"
)

// MockInputFile : implements io.Reader interface for mocking a file read operation testing purposes
type MockInputFile struct {
	data              []byte
	expectedException error
}

// NewMockInputFile : returns a new instance of MockInputFile
func NewMockInputFile(options ...Option) *MockInputFile {
	f := &MockInputFile{
		data:              make([]byte, 0),
		expectedException: nil,
	}

	for _, opt := range options {
		opt(f)
	}

	return f
}

// Option : constructor options
type Option func(file *MockInputFile)

// WithContent : returns a file with the provided content
func WithContent(content string) Option {
	return func(file *MockInputFile) {
		file.data = []byte(content)
	}
}

// WithException : return a file with MockInputFile.Read always returning the provided exception upon invocation
func WithException(err error) Option {
	return func(file *MockInputFile) {
		file.expectedException = err
	}
}

// eof : returns true if we have reached the end of the file
func (r *MockInputFile) eof() bool {
	return len(r.data) == 0
}

// readByte : simulates reading from a file byte by byte
func (r *MockInputFile) readByte() byte {
	// this function assumes that eof() check was done before
	b := r.data[0]
	r.data = r.data[1:]
	return b
}

// Read : reads the content of mock file
func (r *MockInputFile) Read(p []byte) (n int, err error) {
	if r.expectedException != nil {
		return 0, r.expectedException
	}

	if r.eof() {
		err = io.EOF
		return
	}

	if c := cap(p); c > 0 {
		for n < c {
			p[n] = r.readByte()
			n++
			if r.eof() {
				break
			}
		}
	}
	return
}
