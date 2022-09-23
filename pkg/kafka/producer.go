package kafka

import (
	"context"
	"strings"
	"time"

	"github.com/LambdaTest/photon/config"
	"github.com/LambdaTest/photon/pkg/lumber"
	"github.com/LambdaTest/photon/pkg/models"
	"github.com/segmentio/kafka-go"
)

type producer struct {
	topicName     string
	writer        *kafka.Writer
	logger        lumber.Logger
	startTime     time.Time
	totalMessages int64
	totalErrors   int64
	totalBytes    int64
}

// New return a new kafka producer.
func New(cfg *config.Config, logger lumber.Logger) models.Producer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:          strings.Split(cfg.Kafka.Brokers, ","),
		Topic:            cfg.Kafka.Topic,
		ErrorLogger:      kafka.LoggerFunc(logger.Errorf),
		Logger:           kafka.LoggerFunc(logger.Debugf),
		Balancer:         &kafka.RoundRobin{},
		CompressionCodec: kafka.Snappy.Codec(),
		RequiredAcks:     int(kafka.RequireOne), // will wait for acknowledgement from only master.
	})
	logger.Infof("Kafka Producer connection created successfully")
	return &producer{
		topicName: cfg.Kafka.Topic,
		logger:    logger,
		startTime: time.Now(),
		writer:    writer,
	}
}

func (p *producer) WriteMessage(msg []byte, headers ...kafka.Header) error {
	return p.writer.WriteMessages(context.Background(), kafka.Message{
		Value:   msg,
		Headers: headers,
	})
}

func (p *producer) Close() error {
	return p.writer.Close()
}

func (p *producer) Stats() models.ProducerStats {
	if p.writer == nil {
		return models.ProducerStats{}
	}
	duration := time.Since(p.startTime)
	writerStats := p.writer.Stats()
	p.totalMessages += writerStats.Messages
	p.totalErrors += writerStats.Errors
	p.totalBytes += writerStats.Bytes
	rate := float64(p.totalMessages) / duration.Seconds()

	return models.ProducerStats{
		Duration:      duration,
		TotalMessages: p.totalMessages,
		TotalErrors:   p.totalErrors,
		TotalBytes:    p.totalBytes,
		Throughput:    rate,
	}
}
