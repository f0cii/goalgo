package log

import (
	"time"
)

// Entry represents a single log entry.
type Entry struct {
	ID      uint64    `json:"id"`
	Time    time.Time `json:"time"`
	Level   Level     `json:"level"`
	Message string    `json:"message"`
}

func NewEntry(id uint64, level Level, message string) *Entry {
	return &Entry{
		ID:      id,
		Time:    time.Now(),
		Level:   level,
		Message: message,
	}
}
