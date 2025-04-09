package eventnotifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type EventNotifier struct {
	baseURL    string
	httpClient *http.Client
}

const defaultBaseURL = "http://localhost:5182"

func NewDefaultEventNotifier() *EventNotifier {
	return NewEventNotifier(defaultBaseURL)
}

func NewEventNotifier(baseURL string) *EventNotifier {
	return &EventNotifier{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

// NotifyRecordingEvent sends a recording-specific event notification
func (en *EventNotifier) NotifyRecordingEvent(event, pathName string, details *RecordingEventDetails) {
	// Run in background goroutine
	go func() {
		reqBody := RecordingEventRequest{
			Event:     event,
			PathName:  pathName,
			Timestamp: time.Now().UTC(),
			Details:   details,
		}

		jsonBody, err := json.Marshal(reqBody)
		if err != nil {
			// Log error and return since we're in a goroutine
			fmt.Printf("failed to marshal request body: %v\n", err)
			return
		}

		url := fmt.Sprintf("%s/api/Event/recording", en.baseURL)
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
		if err != nil {
			fmt.Printf("failed to create request: %v\n", err)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "*/*")

		resp, err := en.httpClient.Do(req)
		if err != nil {
			fmt.Printf("failed to send recording event notification: %v\n", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("recording event notification failed with status: %d\n", resp.StatusCode)
		}
	}()
}
