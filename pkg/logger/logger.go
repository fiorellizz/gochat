package logger

import (
	"log"
	"os"
	"strings"
)

// Níveis de log aceitos
const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
)

// Logger global configurável
type Logger struct {
	level string
	logger *log.Logger
}

// New cria um novo logger com o nível definido
func New(level string) *Logger {
	if level == "" {
		level = LevelInfo
	}

	level = strings.ToLower(level)

	return &Logger{
		level:  level,
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Debug exibe mensagens apenas se o nível for debug
func (l *Logger) Debug(msg string, args ...interface{}) {
	if l.level == LevelDebug {
		l.logger.Printf("[DEBUG] "+msg, args...)
	}
}

// Info sempre é exibido
func (l *Logger) Info(msg string, args ...interface{}) {
	l.logger.Printf("[INFO] "+msg, args...)
}

// Warn exibe avisos
func (l *Logger) Warn(msg string, args ...interface{}) {
	if l.level == LevelDebug || l.level == LevelInfo || l.level == LevelWarn {
		l.logger.Printf("[WARN] "+msg, args...)
	}
}

// Error sempre é exibido
func (l *Logger) Error(msg string, args ...interface{}) {
	l.logger.Printf("[ERROR] "+msg, args...)
}
