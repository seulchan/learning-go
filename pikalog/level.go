package pikalog

// Level is a custom type representing the severity of a log message.
// We use `byte` as the underlying type because there are few levels,
// making it memory-efficient.
type Level byte

// const block defines the available logging levels.
// `iota` is a Go keyword that simplifies the definition of incrementing numbers.
// It starts at 0 in this const block and increments by 1 for each subsequent constant.
const (
	// LevelDebug represents the lowest level of log, mostly used for debugging purposes.
	// iota will be 0 here.
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information deemed valuable.
	// iota will be 1 here.
	LevelInfo
	// LevelError represents the highest logging level, only to be used to trace errors.
	// iota will be 2 here.
	LevelError
)

// String implements the fmt.Stringer interface
func (lvl Level) String() string {
	switch lvl {
	case LevelDebug:
		// Returns a human-readable string for the Debug level.
		return "[DEBUG]"
	case LevelInfo:
		// Returns a human-readable string for the Info level.
		return "[INFO]"
	case LevelError:
		// Returns a human-readable string for the Error level.
		return "[ERROR]"
	default:
		// This case should ideally not be reached if only predefined levels are used.
		// Returning an empty string is a safe default, but in a more robust logger,
		// you might log an internal error or panic if an unknown level is encountered,
		// depending on the desired behavior.
		return ""
	}
}
