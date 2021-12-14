// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package user is a generated GoMock package.
package user

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/sundaytycoon/buttons-api/internal/constants/model"
)

// MockuserService is a mock of userService interface.
type MockuserService struct {
	ctrl     *gomock.Controller
	recorder *MockuserServiceMockRecorder
}

// MockuserServiceMockRecorder is the mock recorder for MockuserService.
type MockuserServiceMockRecorder struct {
	mock *MockuserService
}

// NewMockuserService creates a new mock instance.
func NewMockuserService(ctrl *gomock.Controller) *MockuserService {
	mock := &MockuserService{ctrl: ctrl}
	mock.recorder = &MockuserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockuserService) EXPECT() *MockuserServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockuserService) Create(ctx context.Context, name, state string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, name, state)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockuserServiceMockRecorder) Create(ctx, name, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockuserService)(nil).Create), ctx, name, state)
}

// Get mocks base method.
func (m *MockuserService) Get(ctx context.Context, id string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockuserServiceMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockuserService)(nil).Get), ctx, id)
}
