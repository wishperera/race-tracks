package input

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/wishperera/race-tracks/log"
	"github.com/wishperera/race-tracks/models"
	"os"
	"path/filepath"
)

type Provider struct {
	log log.Logger
}

func NewProvider(log log.Logger) (*Provider, error) {
	if log == nil {
		return nil, errors.New("param log cannot be empty")
	}

	return &Provider{
		log: log,
	}, nil
}

// ReadInput : reads the input from a given filepath
func (p *Provider) ReadInput(file string) (input models.Input, err error) {
	path, err := filepath.Abs(file)
	if err != nil {
		return input, fmt.Errorf("failed to read input file path due: %w", err)
	}

	readFile, err := os.Open(path)
	if err != nil {
		return input, fmt.Errorf("failed to read input file due: %w", err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	readFile.Close()

	return input, nil
}
