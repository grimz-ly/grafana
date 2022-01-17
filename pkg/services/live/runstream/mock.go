// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/grafana/grafana/pkg/services/live/runstream (interfaces: ChannelLocalPublisher,NumSubscribersGetter,StreamRunner,PluginContextGetter)

// Package runstream is a generated GoMock package.
package runstream

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	backend "github.com/grafana/grafana-plugin-sdk-go/backend"
	models "github.com/grafana/grafana/pkg/models"
)

// MockChannelLocalPublisher is a mock of ChannelLocalPublisher interface.
type MockChannelLocalPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockChannelLocalPublisherMockRecorder
}

// MockChannelLocalPublisherMockRecorder is the mock recorder for MockChannelLocalPublisher.
type MockChannelLocalPublisherMockRecorder struct {
	mock *MockChannelLocalPublisher
}

// NewMockChannelLocalPublisher creates a new mock instance.
func NewMockChannelLocalPublisher(ctrl *gomock.Controller) *MockChannelLocalPublisher {
	mock := &MockChannelLocalPublisher{ctrl: ctrl}
	mock.recorder = &MockChannelLocalPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannelLocalPublisher) EXPECT() *MockChannelLocalPublisherMockRecorder {
	return m.recorder
}

// PublishLocal mocks base method.
func (m *MockChannelLocalPublisher) PublishLocal(arg0 string, arg1 []byte, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishLocal", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishLocal indicates an expected call of PublishLocal.
func (mr *MockChannelLocalPublisherMockRecorder) PublishLocal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishLocal", reflect.TypeOf((*MockChannelLocalPublisher)(nil).PublishLocal), arg0, arg1, arg2)
}

// MockNumSubscribersGetter is a mock of NumSubscribersGetter interface.
type MockNumSubscribersGetter struct {
	ctrl     *gomock.Controller
	recorder *MockNumSubscribersGetterMockRecorder
}

// MockNumSubscribersGetterMockRecorder is the mock recorder for MockNumSubscribersGetter.
type MockNumSubscribersGetterMockRecorder struct {
	mock *MockNumSubscribersGetter
}

// NewMockNumSubscribersGetter creates a new mock instance.
func NewMockNumSubscribersGetter(ctrl *gomock.Controller) *MockNumSubscribersGetter {
	mock := &MockNumSubscribersGetter{ctrl: ctrl}
	mock.recorder = &MockNumSubscribersGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNumSubscribersGetter) EXPECT() *MockNumSubscribersGetterMockRecorder {
	return m.recorder
}

// GetNumLocalSubscribers mocks base method.
func (m *MockNumSubscribersGetter) GetNumLocalSubscribers(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNumLocalSubscribers", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNumLocalSubscribers indicates an expected call of GetNumLocalSubscribers.
func (mr *MockNumSubscribersGetterMockRecorder) GetNumLocalSubscribers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumLocalSubscribers", reflect.TypeOf((*MockNumSubscribersGetter)(nil).GetNumLocalSubscribers), arg0)
}

// GetNumSubscribers mocks base method.
func (m *MockNumSubscribersGetter) GetNumSubscribers(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNumSubscribers", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNumSubscribers indicates an expected call of GetNumSubscribers.
func (mr *MockNumSubscribersGetterMockRecorder) GetNumSubscribers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumSubscribers", reflect.TypeOf((*MockNumSubscribersGetter)(nil).GetNumSubscribers), arg0)
}

// MockStreamRunner is a mock of StreamRunner interface.
type MockStreamRunner struct {
	ctrl     *gomock.Controller
	recorder *MockStreamRunnerMockRecorder
}

// MockStreamRunnerMockRecorder is the mock recorder for MockStreamRunner.
type MockStreamRunnerMockRecorder struct {
	mock *MockStreamRunner
}

// NewMockStreamRunner creates a new mock instance.
func NewMockStreamRunner(ctrl *gomock.Controller) *MockStreamRunner {
	mock := &MockStreamRunner{ctrl: ctrl}
	mock.recorder = &MockStreamRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamRunner) EXPECT() *MockStreamRunnerMockRecorder {
	return m.recorder
}

// RunStream mocks base method.
func (m *MockStreamRunner) RunStream(arg0 context.Context, arg1 *backend.RunStreamRequest, arg2 *backend.StreamSender) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunStream", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunStream indicates an expected call of RunStream.
func (mr *MockStreamRunnerMockRecorder) RunStream(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunStream", reflect.TypeOf((*MockStreamRunner)(nil).RunStream), arg0, arg1, arg2)
}

// MockPluginContextGetter is a mock of PluginContextGetter interface.
type MockPluginContextGetter struct {
	ctrl     *gomock.Controller
	recorder *MockPluginContextGetterMockRecorder
}

// MockPluginContextGetterMockRecorder is the mock recorder for MockPluginContextGetter.
type MockPluginContextGetterMockRecorder struct {
	mock *MockPluginContextGetter
}

// NewMockPluginContextGetter creates a new mock instance.
func NewMockPluginContextGetter(ctrl *gomock.Controller) *MockPluginContextGetter {
	mock := &MockPluginContextGetter{ctrl: ctrl}
	mock.recorder = &MockPluginContextGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPluginContextGetter) EXPECT() *MockPluginContextGetterMockRecorder {
	return m.recorder
}

// GetPluginContext mocks base method.
func (m *MockPluginContextGetter) GetPluginContext(arg0 context.Context, arg1 *models.SignedInUser, arg2, arg3 string, arg4 bool) (backend.PluginContext, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPluginContext", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(backend.PluginContext)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPluginContext indicates an expected call of GetPluginContext.
func (mr *MockPluginContextGetterMockRecorder) GetPluginContext(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginContext", reflect.TypeOf((*MockPluginContextGetter)(nil).GetPluginContext), arg0, arg1, arg2, arg3, arg4)
}
