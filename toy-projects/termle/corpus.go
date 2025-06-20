package termle

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// ErrCorpusIsEmpty is a specific error returned when the word list (corpus) is empty.
// Defining it as a constant allows other parts of the program to check for this specific error
// using `errors.Is(err, termle.ErrCorpusIsEmpty)`.
const ErrCorpusIsEmpty = corpusError("corpus is empty")

// ReadCorpus reads a list of words from a file at the given path.
// It expects the file to contain words separated by whitespace.
func ReadCorpus(path string) ([]string, error) {
	// os.ReadFile reads the entire content of the file into a byte slice.
	data, err := os.ReadFile(path)
	if err != nil {
		// If there's an error reading the file (e.g., file not found, no permissions),
		// wrap the original error with more context using fmt.Errorf and %w.
		return nil, fmt.Errorf("unable to open %q for reading: %w", path, err)
	}

	// If the file is empty, we can't pick a word, so return the predefined ErrCorpusIsEmpty.
	if len(data) == 0 {
		return nil, ErrCorpusIsEmpty
	}

	// strings.Fields splits the file content (converted to a string) into a slice of strings.
	// It splits by any whitespace character (spaces, tabs, newlines).
	words := strings.Fields(string(data))
	return words, nil
}

// pickWord selects a random word from the provided corpus (slice of strings).
func pickWord(corpus []string) string {
	// rand.Intn returns a random integer in [0, n) where n is the length of the corpus.
	index := rand.Intn(len(corpus))
	// Note: For truly random words across different program runs, you'd typically seed
	// the random number generator, e.g., rand.Seed(time.Now().UnixNano()), usually once at program startup.
	return corpus[index]
}
