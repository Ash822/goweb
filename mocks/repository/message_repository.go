// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ash822/goweb/repository (interfaces: MessageRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	entity "github.com/ash822/goweb/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockMessageRepository is a mock of MessageRepository interface.
type MockMessageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMessageRepositoryMockRecorder
}

// MockMessageRepositoryMockRecorder is the mock recorder for MockMessageRepository.
type MockMessageRepositoryMockRecorder struct {
	mock *MockMessageRepository
}

// NewMockMessageRepository creates a new mock instance.
func NewMockMessageRepository(ctrl *gomock.Controller) *MockMessageRepository {
	mock := &MockMessageRepository{ctrl: ctrl}
	mock.recorder = &MockMessageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageRepository) EXPECT() *MockMessageRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMessageRepository) Create(arg0 *entity.MessageResponse) (*entity.MessageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*entity.MessageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMessageRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMessageRepository)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockMessageRepository) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMessageRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMessageRepository)(nil).Delete), arg0)
}

// FindAll mocks base method.
func (m *MockMessageRepository) FindAll() ([]entity.MessageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]entity.MessageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockMessageRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockMessageRepository)(nil).FindAll))
}

// FindById mocks base method.
func (m *MockMessageRepository) FindById(arg0 string) (*entity.MessageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(*entity.MessageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockMessageRepositoryMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockMessageRepository)(nil).FindById), arg0)
}

// Update mocks base method.
func (m *MockMessageRepository) Update(arg0 *entity.MessageResponse) (*entity.MessageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*entity.MessageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMessageRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMessageRepository)(nil).Update), arg0)
}
