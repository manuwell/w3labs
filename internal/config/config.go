package config

type Config struct {
	Logger Logger `envPrefix:"LOG_"`
}

type Logger struct {
	Level string `env:"LEVEL" envDefault:"INFO"`
}
