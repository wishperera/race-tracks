package output

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/wishperera/race-tracks/internal/pkg/log"
)

type Writer struct {
	log log.Logger
}

func NewWriter(logger log.Logger) (w *Writer, err error) {
	if logger == nil {
		return nil, errors.New("param log cannot be empty")
	}

	return &Writer{
		log: logger,
	}, nil
}

func (w *Writer) WriteOutput(results []int, writer io.Writer) (err error) {
	outputWriter := bufio.NewWriter(writer)
	resultToWrite := ""
	for _, v := range results {
		if v < 0 {
			resultToWrite = "No solution.\n"
		} else {
			resultToWrite = fmt.Sprintf("Optimal solution takes %d hops.\n", v)
		}

		_, err = outputWriter.WriteString(resultToWrite)
		if err != nil {
			return fmt.Errorf("failed to write output due: %w", err)
		}
	}

	err = outputWriter.Flush()
	if err != nil {
		return fmt.Errorf("failed to write output due: %w", err)
	}

	return nil
}
