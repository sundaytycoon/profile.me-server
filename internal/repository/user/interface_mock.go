// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package user is a generated GoMock package.
package user

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mysql "github.com/sundaytycoon/buttons-api/infrastructure/mysql"
	model "github.com/sundaytycoon/buttons-api/internal/constants/model"
)

// MockuserStore is a mock of userStore interface.
type MockuserStore struct {
	ctrl     *gomock.Controller
	recorder *MockuserStoreMockRecorder
}

// MockuserStoreMockRecorder is the mock recorder for MockuserStore.
type MockuserStoreMockRecorder struct {
	mock *MockuserStore
}

// NewMockuserStore creates a new mock instance.
func NewMockuserStore(ctrl *gomock.Controller) *MockuserStore {
	mock := &MockuserStore{ctrl: ctrl}
	mock.recorder = &MockuserStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockuserStore) EXPECT() *MockuserStoreMockRecorder {
	return m.recorder
}

// GetUser mocks base method.
func (m *MockuserStore) GetUser(ctx context.Context, tx mysql.ContextExecutor, id string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, tx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockuserStoreMockRecorder) GetUser(ctx, tx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockuserStore)(nil).GetUser), ctx, tx, id)
}

// MockmysqlClient is a mock of mysqlClient interface.
type MockmysqlClient struct {
	ctrl     *gomock.Controller
	recorder *MockmysqlClientMockRecorder
}

// MockmysqlClientMockRecorder is the mock recorder for MockmysqlClient.
type MockmysqlClientMockRecorder struct {
	mock *MockmysqlClient
}

// NewMockmysqlClient creates a new mock instance.
func NewMockmysqlClient(ctrl *gomock.Controller) *MockmysqlClient {
	mock := &MockmysqlClient{ctrl: ctrl}
	mock.recorder = &MockmysqlClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockmysqlClient) EXPECT() *MockmysqlClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockmysqlClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockmysqlClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockmysqlClient)(nil).Close))
}

// Conn mocks base method.
func (m *MockmysqlClient) Conn(ctx context.Context) (*sql.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Conn", ctx)
	ret0, _ := ret[0].(*sql.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Conn indicates an expected call of Conn.
func (mr *MockmysqlClientMockRecorder) Conn(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Conn", reflect.TypeOf((*MockmysqlClient)(nil).Conn), ctx)
}
