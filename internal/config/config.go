package config

type Config struct {
	AwsEndpoint string `env:"AWS_ENDPOINT"`
	Logger      Logger `envPrefix:"LOG_"`
}

type Logger struct {
	Level string `env:"LEVEL" envDefault:"INFO"`
}
