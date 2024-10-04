package settings

type Config struct {
	Server ServerSettings
	Kafka  KafkaSettings
}

type ServerSettings struct {
	Port int
	Mode string
}

type KafkaSettings struct {
	Broker string
}
