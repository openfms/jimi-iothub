package client

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

type RealTimeAudioVideoDataType string

const (
	AudioVideoDataType              RealTimeAudioVideoDataType = "0"
	VideoDataType                   RealTimeAudioVideoDataType = "1"
	TwoWayIntercomDataType          RealTimeAudioVideoDataType = "2"
	MonitorDataType                 RealTimeAudioVideoDataType = "3"
	CenterBroadcastDataType         RealTimeAudioVideoDataType = "4"
	TransparentTransmissionDataType RealTimeAudioVideoDataType = "5"
)

type RealTimeCodeStreamType string

const (
	MainStream RealTimeCodeStreamType = "0"
	SubStream  RealTimeCodeStreamType = "1"
)

type BaseInstructRequest struct {
	DeviceIMEI   string             `url:"deviceImei,required"`
	ServerFlagID int64              `url:"serverFlagId,required"`
	ProNo        ProNumber          `url:"proNo,required"`
	Platform     RequestPlatform    `url:"platform,required"`
	RequestID    int64              `url:"requestId,required"`
	CmdType      RequestCommandType `url:"cmdType"`
	Language     string             `url:"language"`
	Sync         bool               `url:"sync"`
	OfflineFlag  bool               `url:"offlineFlag"`
	Timeout      int                `url:"timeOut"`
	Token        string             `url:"token"`
}
