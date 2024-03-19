package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

type Level int
type Logger struct {
	LogLevel     string `yaml:"log_level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`
	PrintConsole bool   `yaml:"print_console"`
}

func initLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	logLevel, _ := zapcore.ParseLevel(Cfg.LogLevel)
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)

	logger := zap.New(core)
	zap.ReplaceGlobals(logger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create(Cfg.Director)
	ws := io.MultiWriter(file, os.Stdout)
	return zapcore.AddSync(ws)
}
