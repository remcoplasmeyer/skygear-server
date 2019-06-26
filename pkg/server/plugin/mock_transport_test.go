// Code generated by MockGen. DO NOT EDIT.
// Source: transport.go

package plugin

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	router "github.com/skygeario/skygear-server/pkg/server/router"
	skyconfig "github.com/skygeario/skygear-server/pkg/server/skyconfig"
	skydb "github.com/skygeario/skygear-server/pkg/server/skydb"
	reflect "reflect"
)

// MockTransport is a mock of Transport interface
type MockTransport struct {
	ctrl     *gomock.Controller
	recorder *MockTransportMockRecorder
}

// MockTransportMockRecorder is the mock recorder for MockTransport
type MockTransportMockRecorder struct {
	mock *MockTransport
}

// NewMockTransport creates a new mock instance
func NewMockTransport(ctrl *gomock.Controller) *MockTransport {
	mock := &MockTransport{ctrl: ctrl}
	mock.recorder = &MockTransportMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockTransport) EXPECT() *MockTransportMockRecorder {
	return _m.recorder
}

// State mocks base method
func (_m *MockTransport) State() TransportState {
	ret := _m.ctrl.Call(_m, "State")
	ret0, _ := ret[0].(TransportState)
	return ret0
}

// State indicates an expected call of State
func (_mr *MockTransportMockRecorder) State() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "State", reflect.TypeOf((*MockTransport)(nil).State))
}

// SetState mocks base method
func (_m *MockTransport) SetState(_param0 TransportState) {
	_m.ctrl.Call(_m, "SetState", _param0)
}

// SetState indicates an expected call of SetState
func (_mr *MockTransportMockRecorder) SetState(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetState", reflect.TypeOf((*MockTransport)(nil).SetState), arg0)
}

// SendEvent mocks base method
func (_m *MockTransport) SendEvent(name string, in []byte) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "SendEvent", name, in)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendEvent indicates an expected call of SendEvent
func (_mr *MockTransportMockRecorder) SendEvent(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SendEvent", reflect.TypeOf((*MockTransport)(nil).SendEvent), arg0, arg1)
}

// RunLambda mocks base method
func (_m *MockTransport) RunLambda(ctx context.Context, name string, in []byte) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "RunLambda", ctx, name, in)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunLambda indicates an expected call of RunLambda
func (_mr *MockTransportMockRecorder) RunLambda(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RunLambda", reflect.TypeOf((*MockTransport)(nil).RunLambda), arg0, arg1, arg2)
}

// RunHandler mocks base method
func (_m *MockTransport) RunHandler(ctx context.Context, name string, in []byte) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "RunHandler", ctx, name, in)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunHandler indicates an expected call of RunHandler
func (_mr *MockTransportMockRecorder) RunHandler(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RunHandler", reflect.TypeOf((*MockTransport)(nil).RunHandler), arg0, arg1, arg2)
}

// RunHook mocks base method
func (_m *MockTransport) RunHook(ctx context.Context, hookName string, record *skydb.Record, oldRecord *skydb.Record, async bool) (*skydb.Record, error) {
	ret := _m.ctrl.Call(_m, "RunHook", ctx, hookName, record, oldRecord, async)
	ret0, _ := ret[0].(*skydb.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunHook indicates an expected call of RunHook
func (_mr *MockTransportMockRecorder) RunHook(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RunHook", reflect.TypeOf((*MockTransport)(nil).RunHook), arg0, arg1, arg2, arg3, arg4)
}

// RunTimer mocks base method
func (_m *MockTransport) RunTimer(name string, in []byte) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "RunTimer", name, in)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunTimer indicates an expected call of RunTimer
func (_mr *MockTransportMockRecorder) RunTimer(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RunTimer", reflect.TypeOf((*MockTransport)(nil).RunTimer), arg0, arg1)
}

// RunProvider mocks base method
func (_m *MockTransport) RunProvider(context context.Context, request *AuthRequest) (*AuthResponse, error) {
	ret := _m.ctrl.Call(_m, "RunProvider", context, request)
	ret0, _ := ret[0].(*AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunProvider indicates an expected call of RunProvider
func (_mr *MockTransportMockRecorder) RunProvider(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "RunProvider", reflect.TypeOf((*MockTransport)(nil).RunProvider), arg0, arg1)
}

// MockTransportFactory is a mock of TransportFactory interface
type MockTransportFactory struct {
	ctrl     *gomock.Controller
	recorder *MockTransportFactoryMockRecorder
}

// MockTransportFactoryMockRecorder is the mock recorder for MockTransportFactory
type MockTransportFactoryMockRecorder struct {
	mock *MockTransportFactory
}

// NewMockTransportFactory creates a new mock instance
func NewMockTransportFactory(ctrl *gomock.Controller) *MockTransportFactory {
	mock := &MockTransportFactory{ctrl: ctrl}
	mock.recorder = &MockTransportFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockTransportFactory) EXPECT() *MockTransportFactoryMockRecorder {
	return _m.recorder
}

// Open mocks base method
func (_m *MockTransportFactory) Open(path string, args []string, config skyconfig.Configuration) Transport {
	ret := _m.ctrl.Call(_m, "Open", path, args, config)
	ret0, _ := ret[0].(Transport)
	return ret0
}

// Open indicates an expected call of Open
func (_mr *MockTransportFactoryMockRecorder) Open(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Open", reflect.TypeOf((*MockTransportFactory)(nil).Open), arg0, arg1, arg2)
}

// MockBidirectionalTransport is a mock of BidirectionalTransport interface
type MockBidirectionalTransport struct {
	ctrl     *gomock.Controller
	recorder *MockBidirectionalTransportMockRecorder
}

// MockBidirectionalTransportMockRecorder is the mock recorder for MockBidirectionalTransport
type MockBidirectionalTransportMockRecorder struct {
	mock *MockBidirectionalTransport
}

// NewMockBidirectionalTransport creates a new mock instance
func NewMockBidirectionalTransport(ctrl *gomock.Controller) *MockBidirectionalTransport {
	mock := &MockBidirectionalTransport{ctrl: ctrl}
	mock.recorder = &MockBidirectionalTransportMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockBidirectionalTransport) EXPECT() *MockBidirectionalTransportMockRecorder {
	return _m.recorder
}

// SetRouter mocks base method
func (_m *MockBidirectionalTransport) SetRouter(_param0 *router.Router) {
	_m.ctrl.Call(_m, "SetRouter", _param0)
}

// SetRouter indicates an expected call of SetRouter
func (_mr *MockBidirectionalTransportMockRecorder) SetRouter(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetRouter", reflect.TypeOf((*MockBidirectionalTransport)(nil).SetRouter), arg0)
}