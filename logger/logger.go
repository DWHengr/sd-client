package logger

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	logger *zap.Logger
	// Logger logger
	Logger *zap.SugaredLogger
)

func init() {
	New(&Config{
		Level:       -1,
		Development: true,
		Sampling: Sampling{
			Initial:    10,
			Thereafter: 10,
		},
		OutputPath:      []string{"stderr"},
		ErrorOutputPath: []string{"stderr"},
	})

}

func truncateLogFile(logFilePath string) error {
	file, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

// New 创建日志
func New(conf *Config) error {
	cronStr := "0 0 0 * * ?"
	if len(conf.Cron) > 0 {
		cronStr = conf.Cron
	}
	c := cron.New(cron.WithSeconds())
	fmt.Println(cronStr)
	var err error
	_, err = c.AddFunc(cronStr, func() {
		for _, path := range conf.OutputPath {
			if path != "stderr" {
				err = truncateLogFile(path)
			}
		}
		for _, path := range conf.ErrorOutputPath {
			if path != "stderr" {
				err = truncateLogFile(path)
			}
		}
		if err == nil {
			fmt.Println("logs delete succeed")
		} else {
			fmt.Println("logs delete error:", err)
		}
	})
	if err != nil {
		fmt.Println("Error adding cron job", zap.Error(err))
	}
	c.Start()

	return newLogger(zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.Level(conf.Level)),
		Development: conf.Development,
		Sampling: &zap.SamplingConfig{
			Initial:    conf.Sampling.Initial,
			Thereafter: conf.Sampling.Thereafter,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      conf.OutputPath,
		ErrorOutputPaths: conf.ErrorOutputPath,
	})
}

func newLogger(conf zap.Config) error {
	var err error
	logger, err = conf.Build(
		zap.AddStacktrace(zap.ErrorLevel),
		zap.WithCaller(true),
	)

	if err != nil {
		return err
	}

	Logger = logger.Sugar()
	return nil
}

// Sync sync
func Sync() {
	if logger != nil {
		logger.Sync()
	}
}
