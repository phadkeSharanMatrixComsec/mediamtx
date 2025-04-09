package eventnotifier

import "time"

// Config contains settings for event notifications
type Config struct {
	Enabled bool   `yaml:"enabled"`
	BaseURL string `yaml:"baseURL"`
}

// RecordingEventDetails contains optional details about the recording event
type RecordingEventDetails struct {
	Path  string `json:"path,omitempty"`
	Error string `json:"error,omitempty"`
}

// RecordingEventRequest represents the recording event request body
type RecordingEventRequest struct {
	Event     string                 `json:"event"`
	PathName  string                 `json:"pathName"`
	Timestamp time.Time              `json:"timestamp"`
	Details   *RecordingEventDetails `json:"details,omitempty"`
}
