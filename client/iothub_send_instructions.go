package client

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"sync"
)

var (
	serverFlagID   int64
	flagIDMutex    sync.Mutex
	requestID      int64
	requestIDMutex sync.Mutex
)

func getServerFlagID() int64 {
	flagIDMutex.Lock()
	defer flagIDMutex.Unlock()
	serverFlagID++
	return serverFlagID
}

func getRequestID() int64 {
	requestIDMutex.Lock()
	defer requestIDMutex.Unlock()
	requestID++
	return requestID
}

type DeviceInstructionRequest struct {
	DeviceIMEI   string             `url:"deviceImei"`
	CmdContent   string             `url:"cmdContent"`
	ServerFlagID int64              `url:"serverFlagId"`
	ProNo        ProNumber          `url:"proNo"`
	Platform     RequestPlatform    `url:"platform"`
	RequestID    int64              `url:"requestId"`
	CmdType      RequestCommandType `url:"cmdType"`
	Language     string             `url:"language"`
	Sync         bool               `url:"sync"`
	OfflineFlag  bool               `url:"offlineFlag"`
	Timeout      int                `url:"timeOut"`
	Token        string             `url:"token"`
}

func (cli *IotHubClient) NewDeviceInstructionRequest(imei, command string) *DeviceInstructionRequest {
	return &DeviceInstructionRequest{
		DeviceIMEI:   imei,
		CmdContent:   command,
		ProNo:        ProNoOnlineCommand,
		Platform:     RequestPlatformWeb,
		CmdType:      NormallnsCommandType,
		Token:        cli.apiToken,
		OfflineFlag:  true,
		Timeout:      30,
		Sync:         true,
		RequestID:    getRequestID(),
		ServerFlagID: getServerFlagID(),
	}
}
func (cli *IotHubClient) SendDeviceInstruction(request *DeviceInstructionRequest) (*Response, error) {
	values, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	// Send the POST request with x-www-form-urlencoded data
	resp, err := cli.client.R().
		SetBody(values.Encode()).
		Post(cli.endPointURL.String() + "/api/device/sendInstruct")

	if err != nil {
		return nil, err
	}
	apiResponse := &Response{}
	err = json.Unmarshal(resp.Body(), apiResponse)
	if err != nil {
		return nil, err
	}
	return apiResponse, nil
}
