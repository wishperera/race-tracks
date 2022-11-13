package util

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	readOnlyPermissionLevel          = 0o444
	createWriteAppendPermissionLevel = 0o644
)

// OpenInputFile : constructs absolute file path from relative path
// opens the input file and returns reference
func OpenInputFile(path string) (f *os.File, err error) {
	var absolutPath string
	absolutPath, err = filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read input file path due: %w", err)
	}

	f, err = os.OpenFile(absolutPath, os.O_RDONLY, readOnlyPermissionLevel)
	if err != nil {
		return nil, fmt.Errorf("failed to open input file due: %w", err)
	}

	return f, nil
}

// OpenOutputFile : constructs absolute file path from relative path
// opens/creates the output file and returns reference
func OpenOutputFile(path string) (f *os.File, err error) {
	var absolutPath string
	absolutPath, err = filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read output file path due: %w", err)
	}

	f, err = os.OpenFile(absolutPath, os.O_CREATE|os.O_WRONLY, createWriteAppendPermissionLevel)
	if err != nil {
		return nil, fmt.Errorf("failed to open output file due: %w", err)
	}

	return f, nil
}
