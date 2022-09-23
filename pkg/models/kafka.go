package models

import (
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	// EventHeader is the kafka header for webhook event type.
	EventHeader = "event_type"
	// RepoIDHeader is the kafka header for webhook repo_id.
	RepoIDHeader = "repo_id"
	// GitSCMHeader is the kafka header for webhook git_scm provider.
	GitSCMHeader = "git_driver"
)

// Producer represents the Kafka producer.
type Producer interface {
	// WriteMessage write the message to the kafka producer.
	WriteMessage(msg []byte, headers ...kafka.Header) error
	// Close closes the kafka producer connection.
	Close() error
	// Stats returns the kafka producer stats.
	Stats() ProducerStats
}

// ProducerStats are the metrics of kafka producer.
type ProducerStats struct {
	Duration      time.Duration
	TotalMessages int64
	TotalErrors   int64
	TotalBytes    int64
	Throughput    float64
}
