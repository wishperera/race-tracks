package input

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	models2 "github.com/wishperera/race-tracks/internal/models"
	"github.com/wishperera/race-tracks/internal/pkg/log"
)

func TestProvider_ReadInput(t *testing.T) { //nolint:funlen //minor
	type fields struct {
		log *log.MockLogger
	}
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantInput     []models2.Input
		wantErrorLogs []string
		wantErr       string
	}{
		{
			name: "valid file",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				reader: NewMockFile("2\n5 5\n4 0 4 4\n1\n1 4 2 3\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1"),
			},
			wantInput: []models2.Input{
				{
					GridLength: 5,
					GridWidth:  5,
					Start: models2.Coordinate{
						X: 4,
						Y: 0,
					},
					Target: models2.Coordinate{
						X: 4,
						Y: 4,
					},
					Obstacles: []models2.Obstacles{
						{X1: 1, X2: 4, Y1: 2, Y2: 3},
					},
				},
				{
					GridLength: 3,
					GridWidth:  3,
					Start: models2.Coordinate{
						X: 0,
						Y: 0,
					},
					Target: models2.Coordinate{
						X: 2,
						Y: 2,
					},
					Obstacles: []models2.Obstacles{
						{X1: 1, X2: 1, Y1: 0, Y2: 2},
						{X1: 0, X2: 2, Y1: 1, Y2: 1},
					},
				},
			},
			wantErrorLogs: []string{},
			wantErr:       "",
		},
		{
			name: "invalid file, insufficient lines",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				reader: NewMockFile("2\n5 5\n4 0 4 4\n"),
			},
			wantInput:     nil,
			wantErrorLogs: []string{"Insufficient lines of input"},
			wantErr:       "invalid input file format",
		},
		{
			name: "invalid file, invalid value for number of testcases",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				reader: NewMockFile("x\n5 5\n4 0 4 4\n1\n1 4 2 3\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1"),
			},
			wantInput:     nil,
			wantErrorLogs: []string{"Failed to read number of test cases due: strconv.Atoi: parsing \"x\": invalid syntax"},
			wantErr:       "invalid input file format",
		},
		{
			name: "invalid file, incomplete value for grid size, width missing",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				reader: NewMockFile("2\n5 \n4 0 4 4\n1\n1 4 2 3\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1"),
			},
			wantInput:     nil,
			wantErrorLogs: []string{"Failed to read grid size for test case: 1"},
			wantErr:       "invalid input file format",
		},
		{
			name: "invalid file, invalid value for grid length",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				reader: NewMockFile("2\nx 5\n4 0 4 4\n1\n1 4 2 3\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1"),
			},
			wantInput:     nil,
			wantErrorLogs: []string{"Failed to read grid length for test case: 1 due: strconv.Atoi: parsing \"x\": invalid syntax"},
			wantErr:       "invalid input file format",
		},
		{
			name: "invalid file, invalid value for grid width",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				reader: NewMockFile("2\n5 x\n4 0 4 4\n1\n1 4 2 3\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1"),
			},
			wantInput:     nil,
			wantErrorLogs: []string{"Failed to read grid width for test case: 1 due: strconv.Atoi: parsing \"x\": invalid syntax"},
			wantErr:       "invalid input file format",
		},
		{
			name: "invalid file, incomplete start/end",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				reader: NewMockFile("2\n5 5\n4 0 4\n1\n1 4 2 3\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1"),
			},
			wantInput:     nil,
			wantErrorLogs: []string{"Failed to read start/end points for test case: 1"},
			wantErr:       "invalid input file format",
		},
		{
			name: "invalid file, invalid value for start X coordinate",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				reader: NewMockFile("2\n5 5\nx 0 4 4\n1\n1 4 2 3\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1"),
			},
			wantInput:     nil,
			wantErrorLogs: []string{"Failed to read start point x coordinate for test case: 1 due: strconv.Atoi: parsing \"x\": invalid syntax"},
			wantErr:       "invalid input file format",
		},
		{
			name: "invalid file, invalid value for start Y coordinate",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				reader: NewMockFile("2\n5 5\n4 x 4 4\n1\n1 4 2 3\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1"),
			},
			wantInput:     nil,
			wantErrorLogs: []string{"Failed to read start point y coordinate for test case: 1 due: strconv.Atoi: parsing \"x\": invalid syntax"},
			wantErr:       "invalid input file format",
		},
		{
			name: "invalid file, invalid value for end X coordinate",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				reader: NewMockFile("2\n5 5\n4 0 x 4\n1\n1 4 2 3\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1"),
			},
			wantInput:     nil,
			wantErrorLogs: []string{"Failed to read end point x coordinate for test case: 1 due: strconv.Atoi: parsing \"x\": invalid syntax"},
			wantErr:       "invalid input file format",
		},
		{
			name: "invalid file, invalid value for end Y coordinate",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				reader: NewMockFile("2\n5 5\n4 0 4 x\n1\n1 4 2 3\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1"),
			},
			wantInput:     nil,
			wantErrorLogs: []string{"Failed to read end point y coordinate for test case: 1 due: strconv.Atoi: parsing \"x\": invalid syntax"},
			wantErr:       "invalid input file format",
		},
		{
			name: "invalid file, invalid value for number of obstacles",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				reader: NewMockFile("2\n5 5\n4 0 4 4\nx\n1 4 2 3\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1"),
			},
			wantInput:     nil,
			wantErrorLogs: []string{"Failed to read number of obstacles for test case: 1 due: strconv.Atoi: parsing \"x\": invalid syntax"},
			wantErr:       "invalid input file format",
		},
		{
			name: "invalid file, incomplete value for obstacle coordinates",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				reader: NewMockFile("2\n5 5\n4 0 4 4\n1\n1 4 2 \n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1"),
			},
			wantInput:     nil,
			wantErrorLogs: []string{"Failed to read obstacle margins for test case: 1, obstacle group: 1"},
			wantErr:       "invalid input file format",
		},
		{
			name: "invalid file, invalid value for obstacle coordinate x1",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				reader: NewMockFile("2\n5 5\n4 0 4 4\n1\nx 4 2 3\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1"),
			},
			wantInput:     nil,
			wantErrorLogs: []string{"Failed to read x1 for test case: 1, obstacle group: 1, due: strconv.Atoi: parsing \"x\": invalid syntax"},
			wantErr:       "invalid input file format",
		},
		{
			name: "invalid file, invalid value for obstacle coordinate x2",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				reader: NewMockFile("2\n5 5\n4 0 4 4\n1\n1 x 2 3\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1"),
			},
			wantInput:     nil,
			wantErrorLogs: []string{"Failed to read x2 for test case: 1, obstacle group: 1, due: strconv.Atoi: parsing \"x\": invalid syntax"},
			wantErr:       "invalid input file format",
		},
		{
			name: "invalid file, invalid value for obstacle coordinate y1",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				reader: NewMockFile("2\n5 5\n4 0 4 4\n1\n1 4 x 3\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1"),
			},
			wantInput:     nil,
			wantErrorLogs: []string{"Failed to read y1 for test case: 1, obstacle group: 1, due: strconv.Atoi: parsing \"x\": invalid syntax"},
			wantErr:       "invalid input file format",
		},
		{
			name: "invalid file, invalid value for obstacle coordinate y2",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				reader: NewMockFile("2\n5 5\n4 0 4 4\n1\n1 4 2 x\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1"),
			},
			wantInput:     nil,
			wantErrorLogs: []string{"Failed to read y2 for test case: 1, obstacle group: 1, due: strconv.Atoi: parsing \"x\": invalid syntax"},
			wantErr:       "invalid input file format",
		},
	}
	for _, tt := range tests {
		temp := tt
		t.Run(temp.name, func(t *testing.T) {
			p := &Reader{
				log: temp.fields.log,
			}
			gotInput, err := p.ReadInput(temp.args.reader)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			assert.Equal(t, temp.wantErr, gotErr)
			assert.Equal(t, temp.wantErrorLogs, temp.fields.log.ErrorLogs())
			assert.Equal(t, gotInput, temp.wantInput)
		})
	}
}
