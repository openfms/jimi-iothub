package client

import (
	"fmt"
	"math/rand"
	"time"
)

type ProNumber int

const (
	ProNoOnlineCommand                 ProNumber = 128
	ProNoRealTimeAudioVideoRequest     ProNumber = 37121
	ProNoAudioVideoTransmissionControl ProNumber = 37122
	ProNoRemoteVideoPlaybackRequest    ProNumber = 37377
	ProNoRemoteVideoPlaybackControl    ProNumber = 37378
	ProNoQueryAudioVideoResourceList   ProNumber = 37381
	ProNoFTPFileUploadCommand          ProNumber = 37382
	ProNoFTPFileUploadControl          ProNumber = 37383
	ProNoSetTerminalParameters         ProNumber = 33027
	ProNoQueryTerminalParameters       ProNumber = 33028
	ProNoQuerySpecificParameters       ProNumber = 33030
	ProNoCameraShootImmediately        ProNumber = 34817
	ProNoMultimediaDataRetrieval       ProNumber = 34818
	ProNoMultimediaDataUpload          ProNumber = 34819
	ProNoSingleMultimediaDataUpload    ProNumber = 34821
)

type RequestPlatform string

const (
	RequestPlatformWeb RequestPlatform = "web"
	RequestPlatformApp RequestPlatform = "app"
)

type RequestCommandType string

const (
	NormallnsCommandType RequestCommandType = "normallns"
	GeneralCommandType   RequestCommandType = "general"
)

type ResponseCode string

const (
	// Success The command sent successfully and received a return code indicating the device status, such as busy or error.
	Success ResponseCode = "100"

	// InvalidParameter Invalid parameter.
	InvalidParameter ResponseCode = "200"

	// DeviceOffline Device offline.
	DeviceOffline ResponseCode = "300"

	// DeviceNotRegistered Device not registered with route table (No mapping between IMEI and gateway ID is found in Redis).
	DeviceNotRegistered ResponseCode = "301"

	// DeviceBusy The last command has been sent and no response indicating "device busy" is received.
	DeviceBusy ResponseCode = "302"

	// SMSSentSuccessfully SMS sent successfully.
	SMSSentSuccessfully ResponseCode = "303"

	// SMSFailed SMS failed.
	SMSFailed ResponseCode = "304"

	// NetworkError Network error (Interrupted, etc.).
	NetworkError ResponseCode = "400"

	// CodeExecutionException Code execution exception.
	CodeExecutionException ResponseCode = "500"

	// RequestTimeout Request timeout.
	RequestTimeout ResponseCode = "600"
)

type Response struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data ResponseData `json:"data"`
}

type ResponseData struct {
	Code            ResponseCode `json:"_code"`
	Msg             string       `json:"_msg"`
	Type            string       `json:"_type"`
	Content         string       `json:"_content"`
	ProNo           string       `json:"_proNo"`
	Language        string       `json:"_language,omitempty"`
	IMEI            string       `json:"_imei"`
	RouteClientTime string       `json:"_route_client_time"`
	ServerFlagId    string       `json:"_serverFlagId"`
	GateId          string       `json:"_gateId"`
	RouteServerTime string       `json:"_route_server_time"`
}

func GenerateUniqueInstructionID() string {
	// Generate a random number between 0 and 999999999

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := rnd.Intn(1000000000)

	return fmt.Sprintf("%09d", randomNumber)
}
