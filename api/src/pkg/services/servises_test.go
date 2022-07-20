package services_test

import (
	"testing"

	mockdb "github.com/covid-19-project/api/src/pkg/repo/mock"
	"github.com/covid-19-project/api/src/pkg/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var testCases = []struct {
	caseName string
	location string
	// expected []services.NewCasesResponse
	isError  bool
}{
	{"string goes well", "TestLocation", false},
	{"empty string fail", "", true},
	// {"empty space fail", " ", true},
}

func TestGetNewCasesPerLocation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mockdb.NewMockRepo(ctrl)

	service := services.NewService(repo)
	repo.EXPECT().
		GetNewCasesPerLocation(gomock.Any()).
		Times(1)

	for _, tt := range testCases {
		actual, err := service.GetNewCasesPerLocation(tt.location)
		if tt.isError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			require.IsType(t, []services.NewCasesResponse{}, actual)
		}
		
		// if actual != tt.expected {
		// 	t.Errorf("Expected %f but actual %f", tt.expected, actual)
		// }
	}
}
