package commands

import "strings"

// DeviceCommand represents a command type
type DeviceCommand string

// device management commands
const (
	APN       DeviceCommand = "APN"
	ROAMING   DeviceCommand = "ROAMING"
	WIFIAP    DeviceCommand = "WIFIAP"
	SSID      DeviceCommand = "SSID"
	BTNAME    DeviceCommand = "BTNAME"
	VOLUME    DeviceCommand = "VOLUME"
	VOICESW   DeviceCommand = "VOICESW"
	LED       DeviceCommand = "LED"
	TIMESYNC  DeviceCommand = "TIMESYNC"
	TIMEZONE  DeviceCommand = "TIMEZONE"
	NETWORK   DeviceCommand = "NETWORK"
	GTRANS    DeviceCommand = "GTRANS"
	GCALIBRAT DeviceCommand = "GCALIBRAT"
	RANGE     DeviceCommand = "RANGE"
	PASSWORD  DeviceCommand = "PASSWORD"
	VERSION   DeviceCommand = "VERSION"
	RESTORE   DeviceCommand = "RESTORE"
	REBOOT    DeviceCommand = "REBOOT"
	LOGALL    DeviceCommand = "LOG,ALL"
	LOG       DeviceCommand = "LOG"
	STATUS    DeviceCommand = "STATUS"
	PARAM     DeviceCommand = "PARAM"
	PING      DeviceCommand = "PING"
	FORMAT    DeviceCommand = "FORMAT"
)

// video recording commands
const (
	RECORDAUDIO         DeviceCommand = "RECORDAUDIO"
	RECORDAUDIO_SUB     DeviceCommand = "RECORDAUDIO_SUB"
	RECORDSW            DeviceCommand = "RECORDSW"
	MIRROR              DeviceCommand = "MIRROR"
	RATATION            DeviceCommand = "RATATION"
	CAMERA              DeviceCommand = "CAMERA"
	VIDEORESOLUTION_SUB DeviceCommand = "VIDEORESOLUTION_SUB"
	MILE                DeviceCommand = "MILE"
	CAR                 DeviceCommand = "CAR"
)

// tracking commands
const (
	ACCREP           DeviceCommand = "ACCREP"
	TIMER            DeviceCommand = "TIMER"
	ANGLEREP         DeviceCommand = "ANGLEREP"
	EVENTGPS         DeviceCommand = "EVENTGPS"
	BUFFERCACHEQUERY DeviceCommand = "BUFFERCACHEQUERY"
)

// event commands
const (
	FILTER         DeviceCommand = "FILTER"
	UPLOADSW       DeviceCommand = "UPLOADSW"
	ALARMTONE      DeviceCommand = "ALARMTONE"
	VIDEOPARAM     DeviceCommand = "VIDEOPARAM"
	EXBATALM       DeviceCommand = "EXBATALM"
	FATIGUE        DeviceCommand = "FATIGUE"
	SPEED          DeviceCommand = "SPEED"
	MILEAGE        DeviceCommand = "MILEAGE"
	SOSALM         DeviceCommand = "SOSALM"
	SOS            DeviceCommand = "SOS"
	CALL           DeviceCommand = "CALL"
	SENALM         DeviceCommand = "SENALM"
	DEFENSE        DeviceCommand = "DEFENSE"
	DEFENSE_TIME   DeviceCommand = "DEFENSE_TIME"
	SHOCK          DeviceCommand = "SHOCK"
	SHAKEDELAY     DeviceCommand = "SHAKEDELAY"
	NOSDCARDALM    DeviceCommand = "NOSDCARDALM"
	CRASHALM       DeviceCommand = "CRASHALM"
	SENSOR         DeviceCommand = "SENSOR"
	RAPIDACC       DeviceCommand = "RAPIDACC"
	RAPIDDEC       DeviceCommand = "RAPIDDEC"
	RAPIDTURN      DeviceCommand = "RAPIDTURN"
	RAPIDTEST      DeviceCommand = "RAPIDTEST"
	RAPIDSW        DeviceCommand = "RAPIDSW"
	PICTIMER       DeviceCommand = "PICTIMER"
	PICTIMERSIZE   DeviceCommand = "PICTIMERSIZE"
	PICRATE        DeviceCommand = "PICRATE"
	TIMERPICRAM    DeviceCommand = "TIMERPICRAM"
	TIMERPICRAMDEL DeviceCommand = "TIMERPICRAM,DEL"
)

// Accessories commands
const (
	RELAY               DeviceCommand = "RELAY"
	SPEEDOMETER         DeviceCommand = "SPEEDOMETER"
	CARDREADER          DeviceCommand = "CARDREADER"
	DRIVERLEVEL         DeviceCommand = "DRIVERLEVEL"
	EXDEVICESW          DeviceCommand = "EXDEVICESW"
	OILPARAM            DeviceCommand = "OILPARAM"
	OILIDSET            DeviceCommand = "OILIDSET"
	TEMPCOLLECTINTERVAL DeviceCommand = "TEMPCOLLECTINTERVAL"
	UART                DeviceCommand = "UART"
	TCALIBRAT           DeviceCommand = "TCALIBRAT"
)

// DMS commands
const (
	DMSSW              DeviceCommand = "DMSSW"
	DMS_SWITCH         DeviceCommand = "DMS_SWITCH"
	DMS_VOICE_CUSTOM   DeviceCommand = "DMS_VOICE_CUSTOM"
	DMS_ALERT_CUSTOM   DeviceCommand = "DMS_ALERT_CUSTOM"
	DMSCALIBRAT        DeviceCommand = "DMSCALIBRAT"
	DMS_VIRTUAL_SPEED  DeviceCommand = "DMS_VIRTUAL_SPEED"
	DMS_CONTINUITY     DeviceCommand = "DMS_CONTINUITY"
	DMS_CALIB_ABNORMAL DeviceCommand = "DMS_CALIB_ABNORMAL"
	DMS_SECOND_EVENT   DeviceCommand = "DMS_SECOND_EVENT"
)

// communication commands
const (
	COREKITSW       DeviceCommand = "COREKITSW"
	UPLOAD          DeviceCommand = "UPLOAD"
	HTTPUPLOADLIMIT DeviceCommand = "HTTPUPLOADLIMIT"
	RSERVICE        DeviceCommand = "RSERVICE"
	FILELIST        DeviceCommand = "FILELIST"
	SERVER          DeviceCommand = "SERVER"
	Picture         DeviceCommand = "Picture"
	Video           DeviceCommand = "Video"
	RTMP            DeviceCommand = "RTMP"
	REPLAYLIST      DeviceCommand = "REPLAYLIST"
	RLOGIN          DeviceCommand = "RLOGIN"
	HVIDEO          DeviceCommand = "HVIDEO"
	EVIDEO          DeviceCommand = "EVIDEO"
	UPLOADFILE      DeviceCommand = "UPLOADFILE"
	WIFIKIT         DeviceCommand = "WIFIKIT"
)

// GenerateCommand generates a command with optional parameters
func GenerateCommand(command DeviceCommand, params ...string) string {
	// Create a slice to hold the provided parameters
	var filledParams []string

	// Filter out empty parameters
	for _, param := range params {
		if param != "" {
			filledParams = append(filledParams, param)
		}
	}

	// Combine the command and filled parameters
	fullCommand := string(command)
	if len(filledParams) > 0 {
		fullCommand += "," + strings.Join(filledParams, ",")
	}

	return fullCommand
}
