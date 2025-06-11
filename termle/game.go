package termle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

// Game represents the state of a Termle game.
type Game struct {
	// reader is used to get input from the player.
	reader *bufio.Reader
	// solution is the secret word the player needs to guess, stored as a slice of runes.
	// Using runes allows us to correctly handle characters from various languages.
	solution []rune
	// maxAttempts is the maximum number of guesses the player is allowed.
	maxAttempts int
}

// New creates and initializes a new Termle game.
// It takes the player's input source (e.g., os.Stdin), a list of possible words (corpus),
// and the maximum number of attempts allowed.
func New(playerInput io.Reader, corpus []string, maxAttempts int) (*Game, error) {
	// It's important to have words to choose from. If the corpus is empty,
	// we can't start a game, so we return an error.
	if len(corpus) == 0 {
		return nil, ErrCorpusIsEmpty
	}

	g := &Game{
		reader:   bufio.NewReader(playerInput),
		solution: []rune(strings.ToUpper(pickWord(corpus))),
		// The game logic assumes words are of a consistent length,
		// and comparisons are case-insensitive, so we convert the chosen word to uppercase.
		maxAttempts: maxAttempts,
	}

	return g, nil
}

func (g *Game) Play() {
	// Welcome message to the player.
	fmt.Println("Welcome to Termle!")

	// The game loop continues for each attempt, up to g.maxAttempts.
	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		// ask prompts the player for their guess and returns it.
		guess := g.ask()

		// computeFeedback compares the guess against the solution
		// and generates feedback (correct, wrong position, absent).
		fb := computeFeedback(guess, g.solution)
		// Display the feedback to the player (e.g., "💚🟡◻️◻️💚").
		fmt.Println(fb.String())

		// Check if the guess matches the solution.
		if slices.Equal(guess, g.solution) {
			fmt.Printf("🎉 You won! You found it in %d guess(es)! The word was: %s.\n", currentAttempt, string(g.solution))
			return // End the game since the player won.
		}
	}

	// If the loop finishes, it means the player used all attempts without guessing the word.
	fmt.Printf("😞 You've lost! The solution was: %s. \n", string(g.solution))
}

// ask prompts the player for a guess, reads their input, and validates it.
// It continues to prompt until a valid guess is entered.
func (g *Game) ask() []rune {
	// Inform the player about the expected length of the guess.
	fmt.Printf("Enter a %d-character guess:\n", len(g.solution))

	// Loop indefinitely until a valid guess is received.
	for {
		playerInput, _, err := g.reader.ReadLine()
		// Handle potential errors during input reading (e.g., if the input stream closes).
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Termle failed to read your guess: %s\n", err.Error())
			continue
		}
		guess := splitToUppercaseCharacters(string(playerInput))
		err = g.validateGuess(guess)
		if err != nil {
			// If validation fails, inform the player and loop again to ask for input.
			_, _ = fmt.Fprintf(os.Stderr,
				"Your attempt is invalid with Termle's solution: %s.\n",
				err.Error())
		} else {
			// If the guess is valid, return it.
			return guess
		}
	}
}

// errInvalidWordLength is returned when
// the guess has the wrong number of characters.
var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the ➥same number of characters as the solution")

// validateGuess ensures the guess is valid enough.
// For Termle, "valid enough" primarily means the guess has the same number of characters as the solution.
func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != len(g.solution) {
		// Return a formatted error that includes the expected and actual lengths,
		// and wraps the specific errInvalidWordLength for easier error checking by callers.
		return fmt.Errorf("expected %d, got %d, %w",
			len(g.solution), len(guess), errInvalidWordLength)
	}

	return nil
}

// splitToUppercaseCharacters converts the input string to uppercase
// and then splits it into a slice of runes. Using runes ensures that
// multi-byte characters (like 'é' or 'こんにちは') are handled correctly as single characters.
func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

// computeFeedback compares the player's guess against the solution and determines the status of each character.
// - correctPosition: The character is correct and in the right spot.
// - wrongPosition: The character is in the solution but in a different spot.
// - absentCharacter: The character is not in the solution.
func computeFeedback(guess, solution []rune) feedback {
	// Initialize feedback with all characters marked as absent.
	result := make(feedback, len(guess))
	// Keep track of solution characters that have already been used to provide feedback.
	// This is crucial for handling duplicate letters correctly.
	// For example, if the solution is "APPLE" and the guess is "LLLLL",
	// only two 'L's in the guess should receive positive feedback (one correct, one wrong position).
	used := make([]bool, len(solution))

	// This check is a safeguard. In a well-structured game, validateGuess should ensure this condition never occurs.
	if len(guess) != len(solution) {
		_, _ = fmt.Fprintf(os.Stderr, "Internal error! Guess and solution have different lengths: %d vs %d", len(guess), len(solution))
		return result
	}

	// First pass: Check for characters in the correct position.
	for posInGuess, character := range guess {
		if character == solution[posInGuess] {
			result[posInGuess] = correctPosition
			used[posInGuess] = true // Mark this solution character as used.
		}
	}

	// Second pass: Check for characters in the wrong position.
	for posInGuess, character := range guess {
		// Skip characters already marked as correctPosition or if they've already found a wrongPosition match.
		if result[posInGuess] != absentCharacter {
			continue
		}

		// Iterate through the solution to find a match for the current guess character.
		for posInSolution, target := range solution {
			// If this solution character was already used for a correctPosition match,
			// or for a previous wrongPosition match (for a different letter in the guess), skip it.
			if used[posInSolution] {
				continue
			}

			if character == target {
				// Found the character in the solution, but not in the current position (that was pass 1).
				result[posInGuess] = wrongPosition
				used[posInSolution] = true // Mark this solution character as used.
				break                      // Move to the next character in the guess.
			}
		}
	}

	return result
}
