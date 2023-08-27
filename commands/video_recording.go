package commands

import "fmt"

type EnableDisable uint8

const (
	Enable  EnableDisable = 1
	Disable EnableDisable = 0
)

/*
RecordAudio Whether to record audio along with the video.
Example: RECORDAUDIO,0
*/
func RecordAudio(state EnableDisable) string {
	return fmt.Sprintf("%s,%d", RECORDAUDIO, state)
}

/*
RecordAudioSub Whether to record audio along with the  video (playback video & live streaming)
Example: RECORDAUDIO_SUB,0
*/
func RecordAudioSub(state EnableDisable) string {
	return fmt.Sprintf("%s,%d", RECORDAUDIO_SUB, state)
}

type RecordSwitchCamera uint8

const (
	MainCamera RecordSwitchCamera = 1
	SubCamera  RecordSwitchCamera = 2
)

/*
RecordSwitch Set the independent switch for a camera
Example: RECORDSW,1,1
*/
func RecordSwitch(camera RecordSwitchCamera, state EnableDisable) string {
	return fmt.Sprintf("%s,%d,%d", RECORDSW, camera, state)
}

/*
Mirror Control the mirroring of the backup camera (rear-view)
Example: MIRROR,in,OFF
*/
func Mirror(state OnOffState) string {
	return fmt.Sprintf("%s,in,%s", MIRROR, state)
}

type CameraInOut string

const (
	CameraIn  CameraInOut = "IN"
	CameraOut CameraInOut = "OUT"
)

type RotationDegree int

const (
	NoRotation        RotationDegree = 0
	Rotation90Degree  RotationDegree = 90
	Rotation180Degree RotationDegree = 180
	Rotation270Degree RotationDegree = 270
)

/*
Rotation Set the rotation angle of the camera image.
Example: RATATION,IN,90
*/
func Rotation(camera CameraInOut, rotation RotationDegree) string {
	return fmt.Sprintf("%s,%s,%d", RATATION, camera, rotation)
}

type VideoQualityInward uint8

const (
	VideoQuality720P6M VideoQualityInward = 0
	VideoQuality720P3M VideoQualityInward = 1
	VideoQuality4802M  VideoQualityInward = 2
	VideoQuality320    VideoQualityInward = 3
)

func SetInwardVideoQuality(quality VideoQualityInward) string {
	return fmt.Sprintf("%s,%s,%d", CAMERA, CameraIn, quality)
}

type VideoQualityOut uint8

const (
	VideoQualityOut1080P VideoQualityOut = 0
	VideoQualityOut720P  VideoQualityOut = 1
	VideoQualityOut480P  VideoQualityOut = 2
	VideoQualityOut320P  VideoQualityOut = 3
)

func SetOutVideoQuality(quality VideoQualityOut) string {
	return fmt.Sprintf("%s,%s,%d", CAMERA, CameraOut, quality)
}

type VideoResolution uint8

const (
	VideoResolution360 VideoResolution = 0
	VideoResolution480 VideoResolution = 1
	VideoResolution720 VideoResolution = 2
)

// VideoResolutionSub Set the quality for livestreaming or playback video
func VideoResolutionSub(resolution VideoResolution) string {
	return fmt.Sprintf("%s,%d", VIDEORESOLUTION_SUB, resolution)
}

type SpeedUnit uint8

const (
	SpeedUnitKmh SpeedUnit = 0
	SpeedUnitMph SpeedUnit = 1
)

// SetSpeedUnit Set the unit of speed on the overlay of video
func SetSpeedUnit(speedUnit VideoResolution) string {
	return fmt.Sprintf("%s,%d", MILE, speedUnit)
}
