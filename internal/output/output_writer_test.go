package output

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wishperera/race-tracks/internal/pkg/log"
)

func TestWriter_WriteOutput(t *testing.T) {
	type fields struct {
		log *log.MockLogger
	}
	type args struct {
		writer  *MockOutputFile
		results []int
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		wantWrittenOutput string
		wantErr           string
	}{
		{
			name: "successful write",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				writer:  NewMockOutputFile(),
				results: []int{7, -1},
			},
			wantWrittenOutput: "Optimal solution takes 7 hops.\nNo solution.\n",
			wantErr:           "",
		},
		{
			name: "failure to write",
			fields: fields{
				log: log.NewMockLogger(),
			},
			args: args{
				writer:  NewMockOutputFile(WithException(errors.New("something went wrong"))),
				results: []int{7, -1},
			},
			wantWrittenOutput: "",
			wantErr:           "failed to write output due: something went wrong",
		},
	}
	for _, tt := range tests {
		temp := tt
		t.Run(temp.name, func(t *testing.T) {
			w := &Writer{
				log: temp.fields.log,
			}

			err := w.WriteOutput(temp.args.results, temp.args.writer)
			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}

			assert.Equal(t, temp.wantErr, gotErr)
			assert.Equal(t, temp.wantWrittenOutput, string(temp.args.writer.Data()))
		})
	}
}
