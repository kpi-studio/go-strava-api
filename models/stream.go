package models

// StreamSet represents a collection of stream data
type StreamSet struct {
	Time           *TimeStream        `json:"time"`
	Distance       *DistanceStream    `json:"distance"`
	LatLng         *LatLngStream      `json:"latlng"`
	Altitude       *AltitudeStream    `json:"altitude"`
	VelocitySmooth *VelocityStream    `json:"velocity_smooth"`
	Heartrate      *HeartrateStream   `json:"heartrate"`
	Cadence        *CadenceStream     `json:"cadence"`
	Watts          *PowerStream       `json:"watts"`
	Temperature    *TemperatureStream `json:"temp"`
	Moving         *MovingStream      `json:"moving"`
	GradeSmooth    *GradeStream       `json:"grade_smooth"`
}

// BaseStream represents the common fields for all stream types
type BaseStream struct {
	Type         string        `json:"type"`
	Data         []interface{} `json:"data"`
	SeriesType   string        `json:"series_type"`
	OriginalSize int           `json:"original_size"`
	Resolution   string        `json:"resolution"`
}

// TimeStream represents time data
type TimeStream struct {
	BaseStream
	Data []int `json:"data"`
}

// DistanceStream represents distance data
type DistanceStream struct {
	BaseStream
	Data []float64 `json:"data"`
}

// LatLngStream represents latitude/longitude data
type LatLngStream struct {
	BaseStream
	Data [][]float64 `json:"data"`
}

// AltitudeStream represents altitude data
type AltitudeStream struct {
	BaseStream
	Data []float64 `json:"data"`
}

// VelocityStream represents velocity data
type VelocityStream struct {
	BaseStream
	Data []float64 `json:"data"`
}

// HeartrateStream represents heart rate data
type HeartrateStream struct {
	BaseStream
	Data []int `json:"data"`
}

// CadenceStream represents cadence data
type CadenceStream struct {
	BaseStream
	Data []int `json:"data"`
}

// PowerStream represents power data
type PowerStream struct {
	BaseStream
	Data []int `json:"data"`
}

// TemperatureStream represents temperature data
type TemperatureStream struct {
	BaseStream
	Data []int `json:"data"`
}

// MovingStream represents moving/stationary data
type MovingStream struct {
	BaseStream
	Data []bool `json:"data"`
}

// GradeStream represents grade data
type GradeStream struct {
	BaseStream
	Data []float64 `json:"data"`
}

// StreamType represents the type of stream data
type StreamType string

const (
	StreamTypeTime        StreamType = "time"
	StreamTypeDistance    StreamType = "distance"
	StreamTypeLatLng      StreamType = "latlng"
	StreamTypeAltitude    StreamType = "altitude"
	StreamTypeVelocity    StreamType = "velocity_smooth"
	StreamTypeHeartrate   StreamType = "heartrate"
	StreamTypeCadence     StreamType = "cadence"
	StreamTypePower       StreamType = "watts"
	StreamTypeTemperature StreamType = "temp"
	StreamTypeMoving      StreamType = "moving"
	StreamTypeGrade       StreamType = "grade_smooth"
)