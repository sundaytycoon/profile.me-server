// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package user is a generated GoMock package.
package user

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/sundaytycoon/profile.me-server/internal/constants/model"
)

// MockuserRepository is a mock of userRepository interface.
type MockuserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockuserRepositoryMockRecorder
}

// MockuserRepositoryMockRecorder is the mock recorder for MockuserRepository.
type MockuserRepositoryMockRecorder struct {
	mock *MockuserRepository
}

// NewMockuserRepository creates a new mock instance.
func NewMockuserRepository(ctrl *gomock.Controller) *MockuserRepository {
	mock := &MockuserRepository{ctrl: ctrl}
	mock.recorder = &MockuserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockuserRepository) EXPECT() *MockuserRepositoryMockRecorder {
	return m.recorder
}

// GetUser mocks base method.
func (m *MockuserRepository) GetUser(ctx context.Context, id string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockuserRepositoryMockRecorder) GetUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockuserRepository)(nil).GetUser), ctx, id)
}

// Save mocks base method.
func (m *MockuserRepository) Save(ctx context.Context, u *model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, u)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockuserRepositoryMockRecorder) Save(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockuserRepository)(nil).Save), ctx, u)
}
