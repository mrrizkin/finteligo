package logger

import (
	"io"
	"os"
	"path"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/mrrizkin/finteligo/app/config"
)

type Logger struct {
	*zerolog.Logger
}

func New(config *config.Config) (*Logger, error) {
	var writers []io.Writer

	if config.LOG_CONSOLE {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if config.LOG_FILE {
		rf, err := rollingFile(config)
		if err != nil {
			return nil, err
		}
		writers = append(writers, rf)
	}
	mw := io.MultiWriter(writers...)

	switch config.LOG_LEVEL {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "disable":
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}

	logger := zerolog.New(mw).With().Timestamp().Logger()

	logger.Info().
		Bool("fileLogging", config.LOG_FILE).
		Bool("jsonLogOutput", config.LOG_JSON).
		Str("logDirectory", config.LOG_DIR).
		Str("fileName", config.APP_NAME+".log").
		Int("maxSizeMB", config.LOG_MAX_SIZE).
		Int("maxBackups", config.LOG_MAX_BACKUP).
		Int("maxAgeInDays", config.LOG_MAX_AGE).
		Msg("logging configured")

	return &Logger{
		Logger: &logger,
	}, nil
}

func rollingFile(c *config.Config) (io.Writer, error) {
	err := os.MkdirAll(c.LOG_DIR, 0744)
	if err != nil {
		return nil, err
	}

	return &lumberjack.Logger{
		Filename:   path.Join(c.LOG_DIR, c.APP_NAME+".log"),
		MaxBackups: c.LOG_MAX_BACKUP, // files
		MaxSize:    c.LOG_MAX_SIZE,   // megabytes
		MaxAge:     c.LOG_MAX_AGE,    // days
	}, nil
}
