package client

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/openfms/jimi-iothub/utils"
	"github.com/redis/go-redis/v9"
	"net/url"
	"sync"
)

type IotHubClient struct {
	client      *resty.Client
	config      *IotHubConfig
	redis       *redis.Client
	endPointURL *url.URL
	wg          *sync.WaitGroup
}

type JimiIotHub interface {
	Stop()
	EndpointURL() *url.URL
	SendDeviceInstruction(ctx context.Context, request *InstructRequest) (*Response, error)

	DeviceInstructionRequest(ctx context.Context, imei string, command string) (*InstructRequest, error)
	RealTimeAVRequest(ctx context.Context, imei string, deviceModel DeviceModel, cmdContent *RealTimeCmdContent) (*InstructRequest, error)
	RealTimeAVControlRequest(ctx context.Context, imei string, deviceModel DeviceModel, cmdContent *RealTimeControlCmdContent) (*InstructRequest, error)
	ListAVResourcesRequest(ctx context.Context, imei string, deviceModel DeviceModel, cmdContent *AVResourceListCmdContent) (*InstructRequest, error)
	HistoryVideoPlaybackRequest(ctx context.Context, imei string, deviceModel DeviceModel, cmdContent *PlaybackCmdContent) (*InstructRequest, error)
	HistoryPlaybackControlRequest(ctx context.Context, imei string, deviceModel DeviceModel, cmdContent *PlaybackControlCmdContent) (*InstructRequest, error)
}

var (
	_ JimiIotHub = &IotHubClient{}
)

func NewIotHubClient(config *IotHubConfig, redisCli *redis.Client) (*IotHubClient, error) {
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
		redis:       redisCli,
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
