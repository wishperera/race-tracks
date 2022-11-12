package input

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"

	models2 "github.com/wishperera/race-tracks/internal/models"
	"github.com/wishperera/race-tracks/internal/pkg/log"
)

var errInvalidFileFormat = errors.New("invalid input file format")

type Reader struct {
	log log.Logger
}

func NewReader(logger log.Logger) (*Reader, error) {
	if logger == nil {
		return nil, errors.New("param log cannot be empty")
	}

	return &Reader{
		log: logger,
	}, nil
}

// ReadInput : reads the input from a given io.Reader
func (p *Reader) ReadInput(reader io.Reader) (input []models2.Input, err error) {
	fileScanner := bufio.NewScanner(reader)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return p.sanitizeInput(fileLines)
}

func (p *Reader) sanitizeInput(lines []string) (input []models2.Input, err error) {
	input = make([]models2.Input, 0)

	if len(lines) < 4 { //nolint:gomnd //minor
		p.log.Error("Insufficient lines of input")
		return nil, errInvalidFileFormat
	}

	// trim spaces
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	numberOfTestCases, err := strconv.Atoi(lines[0])
	if err != nil {
		p.log.ErrorF("Failed to read number of test cases due: %s", err)
		return nil, errInvalidFileFormat
	}

	currentIndex := 1
	for n := 0; n < numberOfTestCases; n++ {
		gridLength, gridWidth, err := p.parseGridSize(n+1, lines[currentIndex])
		if err != nil {
			return nil, err
		}

		start, end, err := p.parseStartEnd(n+1, lines[currentIndex+1])
		if err != nil {
			return nil, err
		}

		numberOfObstacles, err := strconv.Atoi(lines[currentIndex+2])
		if err != nil {
			p.log.ErrorF("Failed to read number of obstacles for test case: %d due: %s", n+1, err)
			return nil, errInvalidFileFormat
		}

		obstacles := make([]models2.Obstacles, 0)

		var j int
		for j = currentIndex + 3; j < currentIndex+3+numberOfObstacles; j++ {
			obs, err := p.parseObstacles(n+1, j-currentIndex-2, lines[j])
			if err != nil {
				return nil, err
			}

			obstacles = append(obstacles, obs)
		}

		input = append(input, models2.Input{
			GridLength: gridLength,
			GridWidth:  gridWidth,
			Start:      start,
			Target:     end,
			Obstacles:  obstacles,
		})

		currentIndex = j
	}

	return input, nil
}

func (p *Reader) parseGridSize(testCase int, in string) (length, width int, err error) {
	gridSizeStr := strings.Split(in, " ")
	if len(gridSizeStr) != 2 {
		p.log.ErrorF("Failed to read grid size for test case: %d", testCase)
		return length, width, errInvalidFileFormat
	}

	length, err = strconv.Atoi(gridSizeStr[0])
	if err != nil {
		p.log.ErrorF("Failed to read grid length for test case: %d due: %s", testCase, err)
		return length, width, errInvalidFileFormat
	}

	width, err = strconv.Atoi(gridSizeStr[1])
	if err != nil {
		p.log.ErrorF("Failed to read grid width for test case: %d due: %s", testCase, err)
		return length, width, errInvalidFileFormat
	}

	return length, width, nil
}

func (p *Reader) parseStartEnd(testCase int, in string) (start, end models2.Coordinate, err error) {
	points := strings.Split(in, " ")
	if len(points) != 4 { //nolint:gomnd //minor
		p.log.ErrorF("Failed to read start/end points for test case: %d", testCase)
		return start, end, errInvalidFileFormat
	}

	startX, err := strconv.Atoi(points[0])
	if err != nil {
		p.log.ErrorF("Failed to read start point x coordinate for test case: %d due: %s", testCase, err)
		return start, end, errInvalidFileFormat
	}

	startY, err := strconv.Atoi(points[1])
	if err != nil {
		p.log.ErrorF("Failed to read start point y coordinate for test case: %d due: %s", testCase, err)
		return start, end, errInvalidFileFormat
	}

	endX, err := strconv.Atoi(points[2])
	if err != nil {
		p.log.ErrorF("Failed to read end point x coordinate for test case: %d due: %s", testCase, err)
		return start, end, errInvalidFileFormat
	}

	endY, err := strconv.Atoi(points[3])
	if err != nil {
		p.log.ErrorF("Failed to read end point y coordinate for test case: %d due: %s", testCase, err)
		return start, end, errInvalidFileFormat
	}

	start.X = startX
	start.Y = startY
	end.X = endX
	end.Y = endY

	return start, end, nil
}

func (p *Reader) parseObstacles(testCase, obstacleGroup int, in string) (obs models2.Obstacles, err error) {
	points := strings.Split(in, " ")
	if len(points) != 4 { //nolint:gomnd //minor
		p.log.ErrorF("Failed to read obstacle margins for test case: %d, obstacle group: %d", testCase, obstacleGroup)
		return obs, errInvalidFileFormat
	}

	x1, err := strconv.Atoi(points[0])
	if err != nil {
		p.log.ErrorF("Failed to read x1 for test case: %d, obstacle group: %d, due: %s", testCase, obstacleGroup, err)
		return obs, errInvalidFileFormat
	}

	x2, err := strconv.Atoi(points[1])
	if err != nil {
		p.log.ErrorF("Failed to read x2 for test case: %d, obstacle group: %d, due: %s", testCase, obstacleGroup, err)
		return obs, errInvalidFileFormat
	}

	y1, err := strconv.Atoi(points[2])
	if err != nil {
		p.log.ErrorF("Failed to read y1 for test case: %d, obstacle group: %d, due: %s", testCase, obstacleGroup, err)
		return obs, errInvalidFileFormat
	}

	y2, err := strconv.Atoi(points[3])
	if err != nil {
		p.log.ErrorF("Failed to read y2 for test case: %d, obstacle group: %d, due: %s", testCase, obstacleGroup, err)
		return obs, errInvalidFileFormat
	}

	obs.X1 = x1
	obs.X2 = x2
	obs.Y1 = y1
	obs.Y2 = y2

	return obs, nil
}
