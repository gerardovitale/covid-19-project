package mockdb

import (
	"reflect"

	"github.com/covid-19-project/api/src/pkg/services"
	gomock "github.com/golang/mock/gomock"
)

// MockRepo is a mock of Repo interface
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// GetNewCasesPerLocation mocks base method
func (m *MockRepo) GetNewCasesPerLocation(arg0 string) ([]services.NewCasesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNewCasesPerLocation", arg0)
	ret0, _ := ret[0].([]services.NewCasesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNewCasesPerLocation indicates an expected call of GetNewCasesPerLocation
func (mr *MockRepoMockRecorder) GetNewCasesPerLocation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock, "GetNewCasesPerLocation", reflect.TypeOf((*MockRepo)(nil).GetNewCasesPerLocation), arg0)
}
