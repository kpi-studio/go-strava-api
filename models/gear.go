package models

// Gear represents equipment (bike or shoes)
type Gear struct {
	ID            string        `json:"id"`
	Primary       bool          `json:"primary"`
	ResourceState ResourceState `json:"resource_state"`
	Name          string        `json:"name"`
	Distance      float64       `json:"distance"`
	BrandName     string        `json:"brand_name"`
	ModelName     string        `json:"model_name"`
	FrameType     int           `json:"frame_type"`
	Description   string        `json:"description"`
	Weight        float64       `json:"weight"`
}