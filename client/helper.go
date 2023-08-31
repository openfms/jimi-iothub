package client

import (
	"errors"
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
	// ResponseCodeSuccess The command sent successfully and received a return code indicating the device status, such as busy or error.
	ResponseCodeSuccess ResponseCode = "100"

	// ResponseCodeInvalidParameter Invalid parameter.
	ResponseCodeInvalidParameter ResponseCode = "200"

	// ResponseCodeDeviceOffline Device offline.
	ResponseCodeDeviceOffline ResponseCode = "300"

	// ResponseCodeDeviceNotRegistered Device not registered with route table (No mapping between IMEI and gateway ID is found in Redis).
	ResponseCodeDeviceNotRegistered ResponseCode = "301"

	// ResponseCodeDeviceBusy The last command has been sent and no response indicating "device busy" is received.
	ResponseCodeDeviceBusy ResponseCode = "302"

	// ResponseCodeSMSSentSuccessfully SMS sent successfully.
	ResponseCodeSMSSentSuccessfully ResponseCode = "303"

	// ResponseCodeSMSFailed SMS failed.
	ResponseCodeSMSFailed ResponseCode = "304"

	// ResponseCodeNetworkError Network error (Interrupted, etc.).
	ResponseCodeNetworkError ResponseCode = "400"

	// ResponseCodeCodeExecutionException Code execution exception.
	ResponseCodeCodeExecutionException ResponseCode = "500"

	// ResponseCodeRequestTimeout Request timeout.
	ResponseCodeRequestTimeout ResponseCode = "600"
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

type DeviceModel uint8

const (
	DeviceModelJC120  = 1
	DeviceModelJC170  = 2
	DeviceModelJC200  = 3
	DeviceModelJC400  = 4
	DeviceModelJC400P = 5
	DeviceModelJC400D = 6
	DeviceModelJC450  = 7
)

var DeviceModelNames = map[DeviceModel]string{
	DeviceModelJC120:  "JC120",
	DeviceModelJC170:  "JC170",
	DeviceModelJC200:  "JC200",
	DeviceModelJC400:  "JC400",
	DeviceModelJC400P: "JC400P",
	DeviceModelJC400D: "JC400D",
	DeviceModelJC450:  "JC450",
}

var DeviceModelValues = map[string]DeviceModel{
	"JC120":  DeviceModelJC120,
	"JC170":  DeviceModelJC170,
	"JC200":  DeviceModelJC200,
	"JC400":  DeviceModelJC400,
	"JC400P": DeviceModelJC400P,
	"JC400D": DeviceModelJC400D,
	"JC450":  DeviceModelJC450,
}

var (
	ErrUnsupportedRequest = errors.New("request not supported by device model")
	ErrEmptyCmdContent    = errors.New("command content is empty")
)

const (
	RedisServerFlagIDKey = "jimi.iothub.flagid"
	RedisRequestIDKey    = "jimi.iothub.requestid"
)
