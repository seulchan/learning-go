package termle

import "strings"

// hint represents the status of a single character in a guess.
// It's an alias for byte, making it a small and efficient way to store this information.
type hint byte

// feedback is a slice of hints, representing the feedback for an entire guessed word.
// For example, if the guess is "HELLO" and the solution is "HERTZ",
// feedback might look like: [correctPosition, absentCharacter, absentCharacter, wrongPosition, correctPosition]
type feedback []hint

// These constants define the possible states for a character in a guess.
// iota is a Go keyword that simplifies the definition of incrementing numbers.
// So, absentCharacter will be 0, wrongPosition will be 1, and correctPosition will be 2.
const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

// String returns a visual representation of a hint, using emojis.
func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "◻️"
	case wrongPosition:
		return "🟡"
	case correctPosition:
		return "💚"
	default:
		// This case should ideally not be reached if the hint values are managed correctly.
		// It's a fallback to indicate an unexpected hint value.
		return "💔"
	}
}

// String implements the Stringer interface for a slice of hints.
// This allows us to easily print the entire feedback for a guess (e.g., "💚◻️🟡◻️💚").
func (fb feedback) String() string {
	sb := strings.Builder{}
	for _, h := range fb {
		sb.WriteString(h.String())
	}
	return sb.String()
}

// Equal checks if two feedback slices are identical.
// This is useful for testing and comparing feedback results.
func (fb feedback) Equal(other feedback) bool {
	// If the lengths are different, they can't be equal.
	if len(fb) != len(other) {
		return false
	}

	for index, value := range fb { // Iterate through each hint in the feedback.
		if value != other[index] {
			return false
		}
	}

	return true
}
