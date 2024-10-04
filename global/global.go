package global

import (
	"github.com/oeservices/open-kafka-connector/pkg/logger"
	"github.com/oeservices/open-kafka-connector/pkg/settings"
	"github.com/segmentio/kafka-go"
)

var (
	Config        *settings.Config
	Logger        *logger.Zap
	KafkaProducer *kafka.Writer
)
