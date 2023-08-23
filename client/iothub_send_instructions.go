package client

import (
	"encoding/json"
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
	DeviceIMEI   string             `json:"deviceImei"`
	CmdContent   string             `json:"cmdContent"`
	ServerFlagID int64              `json:"serverFlagId"`
	ProNo        ProNumber          `json:"proNo"`
	Platform     RequestPlatform    `json:"platform"`
	RequestID    int64              `json:"requestId"`
	CmdType      RequestCommandType `json:"cmdType"`
	Language     string             `json:"language"`
	Sync         bool               `json:"sync"`
	OfflineFlag  bool               `json:"offlineFlag"`
	Timeout      int                `json:"timeOut"`
	Token        string             `json:"token"`
}

func NewDeviceInstructionRequest(imei, command string) *DeviceInstructionRequest {
	return &DeviceInstructionRequest{
		DeviceIMEI:   imei,
		CmdContent:   command,
		ProNo:        ProNoOnlineCommand,
		Platform:     RequestPlatformWeb,
		CmdType:      NormallnsCommandType,
		Token:        "123456",
		OfflineFlag:  true,
		Timeout:      30,
		Sync:         true,
		RequestID:    getRequestID(),
		ServerFlagID: getServerFlagID(),
	}
}
func (cli *IotHubClient) SendDeviceInstruction(request *DeviceInstructionRequest) (*Response, error) {
	// Convert the CommandRequest to JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	// Send the POST request with x-www-form-urlencoded data
	resp, err := cli.client.R().
		SetBody(string(jsonData)).
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
