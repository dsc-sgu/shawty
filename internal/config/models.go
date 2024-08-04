package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type appConfig struct {
	SharedSecret string       `yaml:"shared_secret"`
	Server       serverConfig `yaml:"server"`
	Pg           pgConfig     `yaml:"postgres"`
	Log          logConfig    `yaml:"log"`
}

type serverConfig struct {
	// debug or release
	Mode string `yaml:"mode"`
	Port uint16 `yaml:"port"`
	Host string `yaml:"host"`
}

type pgConfig struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Name     string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type encoderConfig struct {
	MessageKey   string               `yaml:"message_key"`
	LevelKey     string               `yaml:"level_key"`
	LevelEncoder zapcore.LevelEncoder `yaml:"level_encoder"`
	TimeKey      string               `yaml:"time_key"`
	TimeEncoder  zapcore.TimeEncoder  `yaml:"time_encoder"`
}

type logConfig struct {
	Level            zap.AtomicLevel `yaml:"level"`
	Encoding         string          `yaml:"encoding"`
	OutputPaths      []string        `yaml:"output_paths"`
	ErrorOutputPaths []string        `yaml:"error_output_paths"`
	DevMode          bool            `yaml:"dev_mode"`
	EncoderConfig    encoderConfig   `yaml:"encoder_config"`
}
