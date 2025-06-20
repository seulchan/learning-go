package pikalog_test

import (
	"learning-go/pikalog"
	"testing"
)

// ExampleLogger_Debugf demonstrates the usage of Debugf.
// Go's testing package can run "Example" functions. If the function prints to standard output
// and has a "Output:" comment at the end, the test framework will compare the actual output
// to the content of the comment. This is great for documentation and basic usage tests.
func ExampleLogger_Debugf() {
	debugLogger := pikalog.New(pikalog.LevelDebug)
	debugLogger.Debugf("Hello,%s", "world")
	// The comment below specifies the expected output for this example.
	// Output:{"level":"[DEBUG]","message":"Hello,world"}
}

const (
	debugMessage = "This is a debug message for testing."
	infoMessage  = "This is an info message for testing."
	errorMessage = "This is an error message for testing."
)

// TestLogger_DebugInfoError uses a table-driven test approach to check logging behavior
// at different log levels. Table-driven tests are a common and effective way to test
// various scenarios with minimal boilerplate.
func TestLogger_DebugInfoError(t *testing.T) {
	// `tt` is our "test table". It's a map where keys are test case names (strings)
	// and values are structs defining the input and expected output for each test case.
	tt := map[string]struct {
		level    pikalog.Level // The log level to configure the logger with.
		expected string        // The expected string output from the logger.
	}{
		"debug": {
			level: pikalog.LevelDebug,
			expected: `{"level":"[DEBUG]","message":"` + debugMessage + "\"}\n" +
				`{"level":"[INFO]","message":"` + infoMessage + "\"}\n" +
				`{"level":"[ERROR]","message":"` + errorMessage + "\"}\n",
		},
		"info": {
			level: pikalog.LevelInfo,
			expected: `{"level":"[INFO]","message":"` + infoMessage + "\"}\n" +
				`{"level":"[ERROR]","message":"` + errorMessage + "\"}\n",
		},
		"error": {
			level:    pikalog.LevelError,
			expected: `{"level":"[ERROR]","message":"` + errorMessage + "\"}\n",
		},
	}

	// We iterate over each test case in our table.
	for name, tc := range tt {
		// t.Run creates a subtest. This is good practice because:
		// 1. It allows running specific subtests (e.g., `go test -run TestLogger_DebugInfoError/info`).
		// 2. If one subtest fails, others still run.
		// 3. Test output is neatly grouped by subtest name.
		t.Run(name, func(t *testing.T) {
			// For each test, we create an instance of `testWriter`.
			// This `testWriter` implements `io.Writer` and will capture
			// everything the logger writes, so we can check it.
			tw := &testWriter{}

			// Create a new logger instance for this test case.
			// We set its level according to `tc.level` and, crucially,
			// we use `pikalog.WithOutput(tw)` to make the logger write to our `testWriter`.
			testedLogger := pikalog.New(tc.level, pikalog.WithOutput(tw))

			// Perform the log actions.
			testedLogger.Debugf(debugMessage)
			testedLogger.Infof(infoMessage)
			testedLogger.Errorf(errorMessage)

			// Assert that the captured output in `tw.contents` matches `tc.expected`.
			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}

// testWriter is a helper struct that implements the io.Writer interface.
// testWriter is a struct that implements io.Writer.
// We use it to validate that we can write to a specific output.
type testWriter struct {
	contents string
}

// Write implements the io.Writer interface.
// This method will be called by the logger when it tries to write a log message.
// Instead of writing to the console or a file, it appends the written bytes (converted to a string)
// to its internal `contents` field.
func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil // Return the number of bytes written and no error.
}
