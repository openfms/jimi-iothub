package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/openfms/jimi-iothub/utils"
	"net/url"
	"sync"
)

type IotHubClient struct {
	client      *resty.Client
	apiToken    string
	endPointURL *url.URL
	wg          *sync.WaitGroup
}

type JimiIotHub interface {
	Stop()
	EndpointURL() *url.URL
	SendDeviceInstruction(request *InstructRequest) (*Response, error)
}

func NewIotHubClient(endPoint, proxy, token string) (*IotHubClient, error) {
	endPointURL, err := utils.GetEndpointURL(endPoint)
	if err != nil {
		return nil, err
	}
	client := resty.New().
		SetBaseURL(endPointURL.String()).
		SetHeaders(map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		})
	if len(proxy) > 0 {
		client.SetProxy(proxy)
	}
	return &IotHubClient{
		client:      client,
		wg:          &sync.WaitGroup{},
		endPointURL: endPointURL,
		apiToken:    token,
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
