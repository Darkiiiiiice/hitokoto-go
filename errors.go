package hitokoto

import (
	"fmt"
	"time"
)

type HitokotoError struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Timestamp int64  `json:"ts"`
}

func (e *HitokotoError) Error() string {
	var ts = time.UnixMilli(e.Timestamp)
	return fmt.Sprintf("status: %d message: %s timestamp: %s", e.Status, e.Message, ts.Format(time.UnixDate))
}

func NewHitokotoError(status int, message string, timestamp int64) *HitokotoError {
	return &HitokotoError{
		Status:    status,
		Message:   message,
		Timestamp: timestamp,
	}
}
