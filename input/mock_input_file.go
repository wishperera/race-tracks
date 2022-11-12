package input

import "io"

type MockInputFile struct {
	data []byte
}

func NewMockFile(toRead string) *MockInputFile {
	return &MockInputFile{[]byte(toRead)}
}

func (r MockInputFile) eof() bool {
	return len(r.data) == 0
}

func (r *MockInputFile) readByte() byte {
	// this function assumes that eof() check was done before
	b := r.data[0]
	r.data = r.data[1:]
	return b
}

func (r *MockInputFile) Read(p []byte) (n int, err error) {
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
