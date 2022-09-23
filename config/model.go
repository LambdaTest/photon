package config

import "github.com/LambdaTest/photon/pkg/lumber"

// ConfigWrapper is a wrapper for the config
type ConfigWrapper struct {
	Config `json:"data"`
}

// Config the application's configuration
type Config struct {
	Config    string
	DB        DBConfig
	Kafka     KafkaConfig
	Port      string
	LogFile   string
	LogConfig lumber.LoggingConfig
	Env       string
	Verbose   bool
	Tracing   TracingConfig
}

// TracingConfig provides opentelemetry configurations
type TracingConfig struct {
	// OtelEndpoint for storing host name for otel collector
	OtelEndpoint string
}

// DBConfig providers the mysql db configuration.
type DBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// KafkaConfig provides the kafka configuration.
type KafkaConfig struct {
	Brokers string `json:"brokers"`
	Topic   string `json:"topic"`
}
