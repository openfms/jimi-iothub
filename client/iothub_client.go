package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/openfms/jimi-iothub/utils"
	"net/url"
	"sync"
)

type IotHubClient struct {
	client      *resty.Client
	config      *IotHubConfig
	endPointURL *url.URL
	wg          *sync.WaitGroup
}

type JimiIotHub interface {
	Stop()
	EndpointURL() *url.URL
	SendDeviceInstruction(request *InstructRequest) (*Response, error)

	DeviceInstructionRequest(imei string, deviceModel DeviceModel, command string) (*InstructRequest, error)
	RealTimeAVRequest(imei string, deviceModel DeviceModel, cmdContent *RealTimeCmdContent) (*InstructRequest, error)
	RealTimeAVControlRequest(imei string, deviceModel DeviceModel, cmdContent *RealTimeControlCmdContent) (*InstructRequest, error)
	ListAVResourcesRequest(imei string, deviceModel DeviceModel, cmdContent *AVResourceListCmdContent) (*InstructRequest, error)
	HistoryVideoPlaybackRequest(imei string, deviceModel DeviceModel, cmdContent *PlaybackCmdContent) (*InstructRequest, error)
	HistoryPlaybackControlRequest(imei string, deviceModel DeviceModel, cmdContent *PlaybackControlCmdContent) (*InstructRequest, error)
}

func NewIotHubClient(config *IotHubConfig) (*IotHubClient, error) {
	endPointURL, err := utils.GetEndpointURL(config.EndPoint)
	if err != nil {
		return nil, err
	}
	client := resty.New().
		SetBaseURL(endPointURL.String()).
		SetHeaders(map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		})
	if len(config.Proxy) > 0 {
		client.SetProxy(config.Proxy)
	}
	return &IotHubClient{
		client:      client,
		wg:          &sync.WaitGroup{},
		endPointURL: endPointURL,
		config:      config,
	}, nil
}

func (cli *IotHubClient) Stop() {
	cli.wg.Wait()
}

// EndpointURL returns the URL of the S3 endpoint.
func (cli *IotHubClient) EndpointURL() *url.URL {
	endpoint := *cli.endPointURL // copy to prevent callers from modifying internal state
	return &endpoint
}
