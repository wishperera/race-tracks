package output

type MockOutputFile struct {
	expectedException error
	data              []byte
}

func NewMockOutputFile(options ...Option) *MockOutputFile {
	mof := &MockOutputFile{
		data: make([]byte, 0),
	}

	for _, opt := range options {
		opt(mof)
	}

	return mof
}

type Option func(file *MockOutputFile)

func WithException(exception error) Option {
	return func(file *MockOutputFile) {
		file.expectedException = exception
	}
}

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

func (m *MockOutputFile) Data() []byte {
	return m.data
}

func (m *MockOutputFile) writeByte(in byte) {
	m.data = append(m.data, in)
}
