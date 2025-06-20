package pikalog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Logger is a struct that holds the configuration for our logger.
// It's responsible for formatting and writing log messages.
type Logger struct {
	threshold        Level     // threshold is the minimum level of messages that this logger will output.
	output           io.Writer // output is where the log messages will be written (e.g., console, file).
	maxMessageLength uint      // maxMessageLength is the maximum number of characters for a single log message. 0 means no limit.
}

// New returns you a logger, ready to log at the required threshold.
// Give it a list of configuration functions to tune it at your will.
// The default output is Stdout.
// There is no default maximum length - messages aren't trimmed.
// `threshold` is the minimum log level that this logger will handle.
// `opts ...Option` is a variadic parameter, meaning you can pass zero or more Option functions.
func New(threshold Level, opts ...Option) *Logger {
	// Initialize the Logger with default values.
	lgr := &Logger{
		threshold: threshold,
		// Default output is os.Stdout (the console).
		output: os.Stdout,
		// Default maxMessageLength is 0, meaning messages are not trimmed by default.
		// The comment below is good for learners, explaining the choice for explicitness.
		maxMessageLength: 0, // we could get rid of this line and use the zero value but let's be explicit
	}

	// Apply all a.k.a "functional options" passed by the caller.
	// This loop iterates through each Option function provided in `opts`
	// and calls it, passing the logger instance `lgr` to be configured.
	for _, configFunc := range opts {
		configFunc(lgr)
	}

	return lgr
}

// Debugf formats and prints a message if the logger's threshold is LevelDebug or lower.
// It uses `fmt.Sprintf`-like formatting.
func (l *Logger) Debugf(format string, args ...any) {
	// Check if the logger's configured threshold allows Debug messages.
	// For example, if threshold is LevelInfo, LevelDebug messages will be skipped.
	if l.threshold > LevelDebug {
		return
	}
	// Delegate the actual logging to the internal logf method.
	l.logf(LevelDebug, format, args...)
}

// Infof formats and prints a message if the logger's threshold is LevelInfo or lower.
// It uses `fmt.Sprintf`-like formatting.
func (l *Logger) Infof(format string, args ...any) {
	// Check if the logger's configured threshold allows Info messages.
	if l.threshold > LevelInfo {
		return
	}
	// Delegate the actual logging to the internal logf method.
	l.logf(LevelInfo, format, args...)
}

// Errorf formats and prints a message. Error messages are always logged unless the
// threshold is set to a level higher than LevelError (which isn't defined in this example,
// so effectively, errors are always logged if this method is called and threshold is LevelError or lower).
// It uses `fmt.Sprintf`-like formatting.
func (l *Logger) Errorf(format string, args ...any) {
	// This check might seem redundant if LevelError is the highest.
	// However, it's good practice for consistency and if more levels were added above Error.
	if l.threshold > LevelError {
		return
	}
	// Delegate the actual logging to the internal logf method.
	l.logf(LevelError, format, args...)
}

// Logf formats and prints a message if the provided `lvl` is at or above the logger's threshold.
// This is a more generic logging method that can be used if the log level is determined dynamically.
func (l *Logger) Logf(lvl Level, format string, args ...any) {
	// Check if the logger's configured threshold allows messages of the given `lvl`.
	if l.threshold > lvl {
		return
	}
	// Delegate the actual logging to the internal logf method.
	l.logf(lvl, format, args...)
}

// logf is an unexported (internal) method that handles the actual formatting and writing of the log message.
// It's called by Debugf, Infof, Errorf, and Logf after they've checked the log level.
// `lvl` is the severity level of the current message.
// `format` and `args` are for `fmt.Sprintf`-style message formatting.
func (l *Logger) logf(lvl Level, format string, args ...any) {
	// Format the user-provided message string with its arguments.
	contents := fmt.Sprintf(format, args...)

	// Check if message trimming is enabled (maxMessageLength > 0)
	// and if the current message exceeds this length.
	// It's important to check length in runes, not bytes, to correctly handle
	// multi-byte characters (e.g., emojis, accented letters).
	// `[]rune(contents)` converts the string to a slice of runes (Unicode code points).
	if l.maxMessageLength != 0 && uint(len([]rune(contents))) > l.maxMessageLength {
		contents = string([]rune(contents)[:l.maxMessageLength]) + "[TRIMMED]"
	}

	msg := message{
		Level:   lvl.String(),
		Message: contents,
	}

	// Encode the structured message (level + content) into JSON format.
	// JSON is a common choice for structured logging as it's machine-readable
	// and widely supported.
	formattedMessage, err := json.Marshal(msg)
	if err != nil {
		// If JSON marshaling fails (which is rare for simple structs but possible),
		// we fall back to printing a plain error message to the logger's output.
		// This ensures that the logging attempt itself doesn't crash the application.
		// The `_, _ = ...` is used to explicitly ignore the return values (bytes written, error)
		// from Fprintf, as handling an error while handling another error can get complex.
		_, _ = fmt.Fprintf(l.output, "unable to format message for %v\n", contents)
		return
	}

	// Write the JSON-formatted log message to the configured output (e.g., console).
	// Fprintln adds a newline character at the end, which is typical for log entries.
	// Again, we ignore the return values from Fprintln for simplicity in this example.
	_, _ = fmt.Fprintln(l.output, string(formattedMessage))
}

// message represents the JSON structure of the logged messages.
// This struct is unexported (starts with a lowercase 'm') because it's only used internally by logger.go.
type message struct {
	Level   string `json:"level"`   // `json:"level"` is a struct tag defining how this field is named in the JSON output.
	Message string `json:"message"` // `json:"message"` defines the JSON key for the log content.
}
