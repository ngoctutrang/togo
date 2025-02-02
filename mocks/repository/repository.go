// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/manabie-com/backend/repository (interfaces: I_Repository)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/manabie-com/backend/entity"
	utils "github.com/manabie-com/backend/utils"
)

// MockI_Repository is a mock of I_Repository interface.
type MockI_Repository struct {
	ctrl     *gomock.Controller
	recorder *MockI_RepositoryMockRecorder
}

// MockI_RepositoryMockRecorder is the mock recorder for MockI_Repository.
type MockI_RepositoryMockRecorder struct {
	mock *MockI_Repository
}

// NewMockI_Repository creates a new mock instance.
func NewMockI_Repository(ctrl *gomock.Controller) *MockI_Repository {
	mock := &MockI_Repository{ctrl: ctrl}
	mock.recorder = &MockI_RepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockI_Repository) EXPECT() *MockI_RepositoryMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MockI_Repository) CreateTask(arg0 *entity.Task) *utils.ErrorRest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", arg0)
	ret0, _ := ret[0].(*utils.ErrorRest)
	return ret0
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockI_RepositoryMockRecorder) CreateTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockI_Repository)(nil).CreateTask), arg0)
}

// DeleteTask mocks base method.
func (m *MockI_Repository) DeleteTask(arg0 string) *utils.ErrorRest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", arg0)
	ret0, _ := ret[0].(*utils.ErrorRest)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockI_RepositoryMockRecorder) DeleteTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockI_Repository)(nil).DeleteTask), arg0)
}

// FindTaskByContent mocks base method.
func (m *MockI_Repository) FindTaskByContent(arg0 string) (*entity.Task, *utils.ErrorRest) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTaskByContent", arg0)
	ret0, _ := ret[0].(*entity.Task)
	ret1, _ := ret[1].(*utils.ErrorRest)
	return ret0, ret1
}

// FindTaskByContent indicates an expected call of FindTaskByContent.
func (mr *MockI_RepositoryMockRecorder) FindTaskByContent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTaskByContent", reflect.TypeOf((*MockI_Repository)(nil).FindTaskByContent), arg0)
}

// FindTaskByUserIdAndDay mocks base method.
func (m *MockI_Repository) FindTaskByUserIdAndDay(arg0, arg1 string) ([]entity.Task, *utils.ErrorRest) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTaskByUserIdAndDay", arg0, arg1)
	ret0, _ := ret[0].([]entity.Task)
	ret1, _ := ret[1].(*utils.ErrorRest)
	return ret0, ret1
}

// FindTaskByUserIdAndDay indicates an expected call of FindTaskByUserIdAndDay.
func (mr *MockI_RepositoryMockRecorder) FindTaskByUserIdAndDay(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTaskByUserIdAndDay", reflect.TypeOf((*MockI_Repository)(nil).FindTaskByUserIdAndDay), arg0, arg1)
}

// GetTask mocks base method.
func (m *MockI_Repository) GetTask(arg0 string) (*entity.Task, *utils.ErrorRest) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask", arg0)
	ret0, _ := ret[0].(*entity.Task)
	ret1, _ := ret[1].(*utils.ErrorRest)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask.
func (mr *MockI_RepositoryMockRecorder) GetTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockI_Repository)(nil).GetTask), arg0)
}

// GetTaskAll mocks base method.
func (m *MockI_Repository) GetTaskAll() ([]entity.Task, *utils.ErrorRest) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskAll")
	ret0, _ := ret[0].([]entity.Task)
	ret1, _ := ret[1].(*utils.ErrorRest)
	return ret0, ret1
}

// GetTaskAll indicates an expected call of GetTaskAll.
func (mr *MockI_RepositoryMockRecorder) GetTaskAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskAll", reflect.TypeOf((*MockI_Repository)(nil).GetTaskAll))
}

// UpdateTask mocks base method.
func (m *MockI_Repository) UpdateTask(arg0 *entity.Task) *utils.ErrorRest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", arg0)
	ret0, _ := ret[0].(*utils.ErrorRest)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockI_RepositoryMockRecorder) UpdateTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockI_Repository)(nil).UpdateTask), arg0)
}
