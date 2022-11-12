package output

import (
	"bufio"
	"fmt"
	"github.com/wishperera/race-tracks/log"
	"io"
)

type Writer struct {
	log log.Logger
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
