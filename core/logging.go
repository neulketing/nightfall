package core

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Logger provides structured logging
type Logger struct {
	Output *os.File
}

// NewLogger creates a new logger
func NewLogger(workspace *Workspace) (*Logger, error) {
	logDir := filepath.Join(workspace.Path, "logs")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	logFile := filepath.Join(logDir, fmt.Sprintf("nightfall-%s.log", time.Now().Format("2006-01-02")))
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	return &Logger{Output: file}, nil
}

// Log writes a log entry
func (l *Logger) Log(level string, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	entry := fmt.Sprintf("[%s] [%s] %s\n", timestamp, level, message)
	l.Output.WriteString(entry)
}

// Info logs an informational message
func (l *Logger) Info(message string) {
	l.Log("INFO", message)
}

// Debug logs a debug message
func (l *Logger) Debug(message string) {
	l.Log("DEBUG", message)
}

// Warn logs a warning message
func (l *Logger) Warn(message string) {
	l.Log("WARN", message)
}

// Error logs an error message
func (l *Logger) Error(message string) {
	l.Log("ERROR", message)
}
