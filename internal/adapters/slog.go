package adapters

import (
	"log"
	"log/slog"
	"os"
	"w3labs/internal/config"
	"w3labs/internal/tools"
)

const (
	lwarn  = iota
	linfo  = iota
	lerror = iota
	ldebug = iota
	lfatal = iota
)

var slogLevelsMap map[string]slog.Level = map[string]slog.Level{
	"WARN":  slog.LevelWarn,
	"INFO":  slog.LevelInfo,
	"ERROR": slog.LevelError,
	"DEBUG": slog.LevelDebug,
}

type Slog struct {
	slog *slog.Logger
}

func NewSlog(cfg config.Logger) tools.Logger {
	level, ok := slogLevelsMap[cfg.Level]
	if !ok {
		level = slog.LevelInfo
	}

	slog := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
	}))
	return &Slog{
		slog: slog,
	}
}

func (l *Slog) Info(msg string, fields map[string]any) {
	l.log(linfo, msg, fields)
}

func (l *Slog) Error(msg string, fields map[string]any) {
	l.log(lerror, msg, fields)
}

func (l *Slog) Fatal(msg string, fields map[string]any) {
	l.log(lfatal, msg, fields)
}

// Warn implements [tools.Logger].
func (l *Slog) Warn(msg string, fields map[string]any) {
	l.log(lwarn, msg, fields)
}

func (l *Slog) Debug(msg string, fields map[string]any) {
	l.log(ldebug, msg, fields)
}

func (l *Slog) log(level int, msg string, fields map[string]any) {
	args := []any{}
	if fields != nil {
		for k, v := range fields {
			args = append(args, k, v)
		}
	}

	switch level {
	case lwarn:
		l.slog.Warn(msg, args...)
	case linfo:
		l.slog.Info(msg, args...)
	case lerror:
		l.slog.Error(msg, args...)
	case lfatal:
		l.slog.Error(msg, args...)
		log.Fatal(msg)
	case ldebug:
		l.slog.Debug(msg, args...)
	}
}
