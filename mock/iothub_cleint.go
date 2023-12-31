// Code generated by MockGen. DO NOT EDIT.
// Source: iothub_client.go

// Package mock_client is a generated GoMock package.
package mock_client

import (
	context "context"
	url "net/url"
	reflect "reflect"

	resty "github.com/go-resty/resty/v2"
	gomock "github.com/golang/mock/gomock"
	client "github.com/openfms/jimi-iothub/client"
)

// MockJimiIotHub is a mock of JimiIotHub interface.
type MockJimiIotHub struct {
	ctrl     *gomock.Controller
	recorder *MockJimiIotHubMockRecorder
}

// MockJimiIotHubMockRecorder is the mock recorder for MockJimiIotHub.
type MockJimiIotHubMockRecorder struct {
	mock *MockJimiIotHub
}

// NewMockJimiIotHub creates a new mock instance.
func NewMockJimiIotHub(ctrl *gomock.Controller) *MockJimiIotHub {
	mock := &MockJimiIotHub{ctrl: ctrl}
	mock.recorder = &MockJimiIotHubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJimiIotHub) EXPECT() *MockJimiIotHubMockRecorder {
	return m.recorder
}

// Client mocks base method.
func (m *MockJimiIotHub) Client() *resty.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(*resty.Client)
	return ret0
}

// Client indicates an expected call of Client.
func (mr *MockJimiIotHubMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockJimiIotHub)(nil).Client))
}

// Config mocks base method.
func (m *MockJimiIotHub) Config(canModify bool) *client.IotHubConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config", canModify)
	ret0, _ := ret[0].(*client.IotHubConfig)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockJimiIotHubMockRecorder) Config(canModify interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockJimiIotHub)(nil).Config), canModify)
}

// DeviceInstructionRequest mocks base method.
func (m *MockJimiIotHub) DeviceInstructionRequest(ctx context.Context, imei, command string) (*client.InstructRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeviceInstructionRequest", ctx, imei, command)
	ret0, _ := ret[0].(*client.InstructRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeviceInstructionRequest indicates an expected call of DeviceInstructionRequest.
func (mr *MockJimiIotHubMockRecorder) DeviceInstructionRequest(ctx, imei, command interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeviceInstructionRequest", reflect.TypeOf((*MockJimiIotHub)(nil).DeviceInstructionRequest), ctx, imei, command)
}

// EndpointURL mocks base method.
func (m *MockJimiIotHub) EndpointURL() *url.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndpointURL")
	ret0, _ := ret[0].(*url.URL)
	return ret0
}

// EndpointURL indicates an expected call of EndpointURL.
func (mr *MockJimiIotHubMockRecorder) EndpointURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndpointURL", reflect.TypeOf((*MockJimiIotHub)(nil).EndpointURL))
}

// GenerateDeviceConfigLinks mocks base method.
func (m *MockJimiIotHub) GenerateDeviceConfigLinks(rtmpPrefix string) *client.DeviceConfigLinks {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateDeviceConfigLinks", rtmpPrefix)
	ret0, _ := ret[0].(*client.DeviceConfigLinks)
	return ret0
}

// GenerateDeviceConfigLinks indicates an expected call of GenerateDeviceConfigLinks.
func (mr *MockJimiIotHubMockRecorder) GenerateDeviceConfigLinks(rtmpPrefix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateDeviceConfigLinks", reflect.TypeOf((*MockJimiIotHub)(nil).GenerateDeviceConfigLinks), rtmpPrefix)
}

// GenerateFLVHistoryLink mocks base method.
func (m *MockJimiIotHub) GenerateFLVHistoryLink(secure bool, channel int, imei string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateFLVHistoryLink", secure, channel, imei)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateFLVHistoryLink indicates an expected call of GenerateFLVHistoryLink.
func (mr *MockJimiIotHubMockRecorder) GenerateFLVHistoryLink(secure, channel, imei interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateFLVHistoryLink", reflect.TypeOf((*MockJimiIotHub)(nil).GenerateFLVHistoryLink), secure, channel, imei)
}

// GenerateFLVReplayLink mocks base method.
func (m *MockJimiIotHub) GenerateFLVReplayLink(secure bool, prefix, imei string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateFLVReplayLink", secure, prefix, imei)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateFLVReplayLink indicates an expected call of GenerateFLVReplayLink.
func (mr *MockJimiIotHubMockRecorder) GenerateFLVReplayLink(secure, prefix, imei interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateFLVReplayLink", reflect.TypeOf((*MockJimiIotHub)(nil).GenerateFLVReplayLink), secure, prefix, imei)
}

// GenerateFlvLiveLink mocks base method.
func (m *MockJimiIotHub) GenerateFlvLiveLink(secure bool, prefix string, channel int, imei string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateFlvLiveLink", secure, prefix, channel, imei)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateFlvLiveLink indicates an expected call of GenerateFlvLiveLink.
func (mr *MockJimiIotHubMockRecorder) GenerateFlvLiveLink(secure, prefix, channel, imei interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateFlvLiveLink", reflect.TypeOf((*MockJimiIotHub)(nil).GenerateFlvLiveLink), secure, prefix, channel, imei)
}

// GenerateRtmpLiveLink mocks base method.
func (m *MockJimiIotHub) GenerateRtmpLiveLink(secure bool, prefix string, channel int, imei string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRtmpLiveLink", secure, prefix, channel, imei)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRtmpLiveLink indicates an expected call of GenerateRtmpLiveLink.
func (mr *MockJimiIotHubMockRecorder) GenerateRtmpLiveLink(secure, prefix, channel, imei interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRtmpLiveLink", reflect.TypeOf((*MockJimiIotHub)(nil).GenerateRtmpLiveLink), secure, prefix, channel, imei)
}

// GenerateVideoLinks mocks base method.
func (m *MockJimiIotHub) GenerateVideoLinks(secure bool, prefix string, channel int, imei string) (*client.VideoLinks, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateVideoLinks", secure, prefix, channel, imei)
	ret0, _ := ret[0].(*client.VideoLinks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateVideoLinks indicates an expected call of GenerateVideoLinks.
func (mr *MockJimiIotHubMockRecorder) GenerateVideoLinks(secure, prefix, channel, imei interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateVideoLinks", reflect.TypeOf((*MockJimiIotHub)(nil).GenerateVideoLinks), secure, prefix, channel, imei)
}

// GetEndpointHost mocks base method.
func (m *MockJimiIotHub) GetEndpointHost() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEndpointHost")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEndpointHost indicates an expected call of GetEndpointHost.
func (mr *MockJimiIotHubMockRecorder) GetEndpointHost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEndpointHost", reflect.TypeOf((*MockJimiIotHub)(nil).GetEndpointHost))
}

// HistoryPlaybackControlRequest mocks base method.
func (m *MockJimiIotHub) HistoryPlaybackControlRequest(ctx context.Context, imei string, deviceModel client.DeviceModel, cmdContent *client.PlaybackControlCmdContent) (*client.InstructRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HistoryPlaybackControlRequest", ctx, imei, deviceModel, cmdContent)
	ret0, _ := ret[0].(*client.InstructRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HistoryPlaybackControlRequest indicates an expected call of HistoryPlaybackControlRequest.
func (mr *MockJimiIotHubMockRecorder) HistoryPlaybackControlRequest(ctx, imei, deviceModel, cmdContent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HistoryPlaybackControlRequest", reflect.TypeOf((*MockJimiIotHub)(nil).HistoryPlaybackControlRequest), ctx, imei, deviceModel, cmdContent)
}

// HistoryVideoPlaybackRequest mocks base method.
func (m *MockJimiIotHub) HistoryVideoPlaybackRequest(ctx context.Context, imei string, deviceModel client.DeviceModel, cmdContent *client.PlaybackCmdContent) (*client.InstructRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HistoryVideoPlaybackRequest", ctx, imei, deviceModel, cmdContent)
	ret0, _ := ret[0].(*client.InstructRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HistoryVideoPlaybackRequest indicates an expected call of HistoryVideoPlaybackRequest.
func (mr *MockJimiIotHubMockRecorder) HistoryVideoPlaybackRequest(ctx, imei, deviceModel, cmdContent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HistoryVideoPlaybackRequest", reflect.TypeOf((*MockJimiIotHub)(nil).HistoryVideoPlaybackRequest), ctx, imei, deviceModel, cmdContent)
}

// ListAVResourcesRequest mocks base method.
func (m *MockJimiIotHub) ListAVResourcesRequest(ctx context.Context, imei string, deviceModel client.DeviceModel, cmdContent *client.AVResourceListCmdContent) (*client.InstructRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAVResourcesRequest", ctx, imei, deviceModel, cmdContent)
	ret0, _ := ret[0].(*client.InstructRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAVResourcesRequest indicates an expected call of ListAVResourcesRequest.
func (mr *MockJimiIotHubMockRecorder) ListAVResourcesRequest(ctx, imei, deviceModel, cmdContent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAVResourcesRequest", reflect.TypeOf((*MockJimiIotHub)(nil).ListAVResourcesRequest), ctx, imei, deviceModel, cmdContent)
}

// RealTimeAVControlRequest mocks base method.
func (m *MockJimiIotHub) RealTimeAVControlRequest(ctx context.Context, imei string, deviceModel client.DeviceModel, cmdContent *client.RealTimeControlCmdContent) (*client.InstructRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RealTimeAVControlRequest", ctx, imei, deviceModel, cmdContent)
	ret0, _ := ret[0].(*client.InstructRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RealTimeAVControlRequest indicates an expected call of RealTimeAVControlRequest.
func (mr *MockJimiIotHubMockRecorder) RealTimeAVControlRequest(ctx, imei, deviceModel, cmdContent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RealTimeAVControlRequest", reflect.TypeOf((*MockJimiIotHub)(nil).RealTimeAVControlRequest), ctx, imei, deviceModel, cmdContent)
}

// RealTimeAVRequest mocks base method.
func (m *MockJimiIotHub) RealTimeAVRequest(ctx context.Context, imei string, deviceModel client.DeviceModel, cmdContent *client.RealTimeCmdContent) (*client.InstructRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RealTimeAVRequest", ctx, imei, deviceModel, cmdContent)
	ret0, _ := ret[0].(*client.InstructRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RealTimeAVRequest indicates an expected call of RealTimeAVRequest.
func (mr *MockJimiIotHubMockRecorder) RealTimeAVRequest(ctx, imei, deviceModel, cmdContent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RealTimeAVRequest", reflect.TypeOf((*MockJimiIotHub)(nil).RealTimeAVRequest), ctx, imei, deviceModel, cmdContent)
}

// SendDeviceInstruction mocks base method.
func (m *MockJimiIotHub) SendDeviceInstruction(ctx context.Context, request *client.InstructRequest) (*client.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendDeviceInstruction", ctx, request)
	ret0, _ := ret[0].(*client.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendDeviceInstruction indicates an expected call of SendDeviceInstruction.
func (mr *MockJimiIotHubMockRecorder) SendDeviceInstruction(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDeviceInstruction", reflect.TypeOf((*MockJimiIotHub)(nil).SendDeviceInstruction), ctx, request)
}

// Stop mocks base method.
func (m *MockJimiIotHub) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockJimiIotHubMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockJimiIotHub)(nil).Stop))
}
