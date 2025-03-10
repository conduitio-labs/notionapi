// Code generated by MockGen. DO NOT EDIT.
// Source: block.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	notionapi "github.com/conduitio-labs/notionapi"
	gomock "github.com/golang/mock/gomock"
)

// MockBlockService is a mock of BlockService interface.
type MockBlockService struct {
	ctrl     *gomock.Controller
	recorder *MockBlockServiceMockRecorder
}

// MockBlockServiceMockRecorder is the mock recorder for MockBlockService.
type MockBlockServiceMockRecorder struct {
	mock *MockBlockService
}

// NewMockBlockService creates a new mock instance.
func NewMockBlockService(ctrl *gomock.Controller) *MockBlockService {
	mock := &MockBlockService{ctrl: ctrl}
	mock.recorder = &MockBlockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlockService) EXPECT() *MockBlockServiceMockRecorder {
	return m.recorder
}

// AppendChildren mocks base method.
func (m *MockBlockService) AppendChildren(arg0 context.Context, arg1 notionapi.BlockID, arg2 *notionapi.AppendBlockChildrenRequest) (*notionapi.AppendBlockChildrenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendChildren", arg0, arg1, arg2)
	ret0, _ := ret[0].(*notionapi.AppendBlockChildrenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppendChildren indicates an expected call of AppendChildren.
func (mr *MockBlockServiceMockRecorder) AppendChildren(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendChildren", reflect.TypeOf((*MockBlockService)(nil).AppendChildren), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockBlockService) Delete(arg0 context.Context, arg1 notionapi.BlockID) (notionapi.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(notionapi.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockBlockServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBlockService)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockBlockService) Get(arg0 context.Context, arg1 notionapi.BlockID) (notionapi.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(notionapi.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBlockServiceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBlockService)(nil).Get), arg0, arg1)
}

// GetChildren mocks base method.
func (m *MockBlockService) GetChildren(arg0 context.Context, arg1 notionapi.BlockID, arg2 *notionapi.Pagination) (*notionapi.GetChildrenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChildren", arg0, arg1, arg2)
	ret0, _ := ret[0].(*notionapi.GetChildrenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChildren indicates an expected call of GetChildren.
func (mr *MockBlockServiceMockRecorder) GetChildren(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChildren", reflect.TypeOf((*MockBlockService)(nil).GetChildren), arg0, arg1, arg2)
}

// Update mocks base method.
func (m *MockBlockService) Update(ctx context.Context, id notionapi.BlockID, request *notionapi.BlockUpdateRequest) (notionapi.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, request)
	ret0, _ := ret[0].(notionapi.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBlockServiceMockRecorder) Update(ctx, id, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBlockService)(nil).Update), ctx, id, request)
}

// MockBlock is a mock of Block interface.
type MockBlock struct {
	ctrl     *gomock.Controller
	recorder *MockBlockMockRecorder
}

// MockBlockMockRecorder is the mock recorder for MockBlock.
type MockBlockMockRecorder struct {
	mock *MockBlock
}

// NewMockBlock creates a new mock instance.
func NewMockBlock(ctrl *gomock.Controller) *MockBlock {
	mock := &MockBlock{ctrl: ctrl}
	mock.recorder = &MockBlockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlock) EXPECT() *MockBlockMockRecorder {
	return m.recorder
}

// GetArchived mocks base method.
func (m *MockBlock) GetArchived() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArchived")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetArchived indicates an expected call of GetArchived.
func (mr *MockBlockMockRecorder) GetArchived() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArchived", reflect.TypeOf((*MockBlock)(nil).GetArchived))
}

// GetCreatedBy mocks base method.
func (m *MockBlock) GetCreatedBy() *notionapi.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreatedBy")
	ret0, _ := ret[0].(*notionapi.User)
	return ret0
}

// GetCreatedBy indicates an expected call of GetCreatedBy.
func (mr *MockBlockMockRecorder) GetCreatedBy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreatedBy", reflect.TypeOf((*MockBlock)(nil).GetCreatedBy))
}

// GetCreatedTime mocks base method.
func (m *MockBlock) GetCreatedTime() *time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreatedTime")
	ret0, _ := ret[0].(*time.Time)
	return ret0
}

// GetCreatedTime indicates an expected call of GetCreatedTime.
func (mr *MockBlockMockRecorder) GetCreatedTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreatedTime", reflect.TypeOf((*MockBlock)(nil).GetCreatedTime))
}

// GetHasChildren mocks base method.
func (m *MockBlock) GetHasChildren() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHasChildren")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetHasChildren indicates an expected call of GetHasChildren.
func (mr *MockBlockMockRecorder) GetHasChildren() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHasChildren", reflect.TypeOf((*MockBlock)(nil).GetHasChildren))
}

// GetID mocks base method.
func (m *MockBlock) GetID() notionapi.BlockID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(notionapi.BlockID)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockBlockMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockBlock)(nil).GetID))
}

// GetLastEditedBy mocks base method.
func (m *MockBlock) GetLastEditedBy() *notionapi.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastEditedBy")
	ret0, _ := ret[0].(*notionapi.User)
	return ret0
}

// GetLastEditedBy indicates an expected call of GetLastEditedBy.
func (mr *MockBlockMockRecorder) GetLastEditedBy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastEditedBy", reflect.TypeOf((*MockBlock)(nil).GetLastEditedBy))
}

// GetLastEditedTime mocks base method.
func (m *MockBlock) GetLastEditedTime() *time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastEditedTime")
	ret0, _ := ret[0].(*time.Time)
	return ret0
}

// GetLastEditedTime indicates an expected call of GetLastEditedTime.
func (mr *MockBlockMockRecorder) GetLastEditedTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastEditedTime", reflect.TypeOf((*MockBlock)(nil).GetLastEditedTime))
}

// GetObject mocks base method.
func (m *MockBlock) GetObject() notionapi.ObjectType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObject")
	ret0, _ := ret[0].(notionapi.ObjectType)
	return ret0
}

// GetObject indicates an expected call of GetObject.
func (mr *MockBlockMockRecorder) GetObject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockBlock)(nil).GetObject))
}

// GetType mocks base method.
func (m *MockBlock) GetType() notionapi.BlockType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetType")
	ret0, _ := ret[0].(notionapi.BlockType)
	return ret0
}

// GetType indicates an expected call of GetType.
func (mr *MockBlockMockRecorder) GetType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockBlock)(nil).GetType))
}
