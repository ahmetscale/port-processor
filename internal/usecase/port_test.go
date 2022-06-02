package usecase

import (
	"context"
	"io"
	"port-processor/internal/entity"
	"port-processor/internal/usecase/repo"
	"port-processor/pkg/db"
	"strings"
	"testing"
)

const validJson1 = `
{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman-test",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  },
  "AEAUH": {
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ],
    "code": "52001"
  }
}
`

const inValidJson1 = `
[
  {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman-test",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  },
  "AEAUH": {
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ],
    "code": "52001"
  }
]
`

func TestPortUseCase_Process(t *testing.T) {

	dbConn := db.New("sqlite", "file::memory:?cache=shared", true)

	err := db.Migrate(dbConn, &entity.Port{})
	if err != nil {
		t.Fatal(err)
	}

	testRepo := repo.NewPortRepo(dbConn)

	validReader := strings.NewReader(validJson1)
	invalidReader := strings.NewReader(inValidJson1)

	type args struct {
		ctx         context.Context
		reader      io.Reader
		workerCount int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test valid json",
			args: args{
				ctx:         context.Background(),
				reader:      validReader,
				workerCount: 2,
			},
			wantErr: false,
		},
		{
			name: "Test invalid json",
			args: args{
				ctx:         context.Background(),
				reader:      invalidReader,
				workerCount: 2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PortUseCase{
				repo: testRepo,
			}
			if err = p.Process(tt.args.ctx, tt.args.reader, tt.args.workerCount); (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
