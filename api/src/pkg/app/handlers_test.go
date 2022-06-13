package app

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	mockdb "github.com/covid-19-project/api/src/pkg/repo/mock"
	"github.com/covid-19-project/api/src/pkg/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestApiStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockdb.NewMockRepo(ctrl)
	server := setupTestServer(t, repo)
	recorder := httptest.NewRecorder()

	url := "/v1/api/status"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetNewCasesPerLocation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockdb.NewMockRepo(ctrl)
	repo.EXPECT().
		GetNewCasesPerLocation(gomock.Any()).
		Times(1)

	server := setupTestServer(t, repo)
	recorder := httptest.NewRecorder()

	url := "/v1/api/new_cases"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	log.Println(reflect.TypeOf(recorder.Body))
	log.Println(recorder.Body)
}
