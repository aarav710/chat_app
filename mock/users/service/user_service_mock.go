// Code generated by MockGen. DO NOT EDIT.
// Source: ./users/service/user_service.go

// Package mock is a generated GoMock package.
package mock

import (
	ent "chatapp/backend/ent"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// GetUserById mocks base method.
func (m *MockUserService) GetUserById(userId int) *ent.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", userId)
	ret0, _ := ret[0].(*ent.User)
	return ret0
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserServiceMockRecorder) GetUserById(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserService)(nil).GetUserById), userId)
}