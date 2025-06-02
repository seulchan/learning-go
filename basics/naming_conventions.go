package main

import "fmt" // Import fmt for printing examples

// --- Go Naming Conventions ---
//
// Effective naming is crucial for writing clear, readable, and maintainable Go code.
// Go has established conventions that are widely followed by the community.
// This file demonstrates these conventions with examples.

// 1. MixedCaps (CamelCase) for Identifiers:
//    Go uses `MixedCaps` or `mixedCaps` (also known as CamelCase) for naming
//    variables, constants, functions, types, methods, and struct fields.
//    - `MixedCaps` (starting with an uppercase letter) is used for exported identifiers.
//    - `mixedCaps` (starting with a lowercase letter) is used for unexported identifiers.
//    Avoid using underscores (`snake_case`) for these kinds of identifiers in Go.

// Example of a struct type using MixedCaps for the type name and its fields.
// Since `Employee`, `FirstName`, `LastName`, and `Age` start with an uppercase letter,
// they are "exported" and can be accessed from other packages.
type Employee struct {
	FirstName string // Corrected typo from FisrtName
	LastName  string
	Age       int
	companyID string // Starts with lowercase, so it's unexported (internal to this package)
}

// Example of an interface. Interface names often end with "er".
type StringFormatter interface {
	FormatString(input string) string
}

// 2. Visibility: Exported vs. Unexported
//    The case of the first letter of an identifier determines its visibility:
//    - Uppercase First Letter (e.g., `MyVariable`, `CalculateTotal`): The identifier is EXPORTED.
//      This means it's visible and accessible from other packages that import this package.
//    - Lowercase First Letter (e.g., `myVariable`, `calculateTotal`): The identifier is UNEXPORTED.
//      This means it's private or internal to the current package and cannot be accessed directly
//      from other packages.

var ExportedVariable = "I can be accessed from other packages."
var unexportedVariable = "I am only accessible within this 'main' package."

const ExportedConstant = 3.14159
const internalConstant = 2.71828

// This function is exported because it starts with an uppercase letter.
func CalculateArea(width, height int) int {
	return width * height
}

// This function is unexported (internal to this package).
func calculatePerimeter(width, height int) int {
	return 2 * (width + height)
}

// 3. Package Names:
//    - Should be short, concise, all lowercase.
//    - Examples: `fmt`, `net/http`, `strings`, `mypackage`.
//    - Avoid underscores (`under_scores`) or mixedCaps (`mixedCapsNames`) in package names.
//    - The package name is the default identifier when importing the package.
//      e.g., `import "math/rand"` -> use `rand.Intn()`

// 4. Constants:
//   - Named using `MixedCaps` or `mixedCaps` just like variables.
//   - If a constant needs to be exported, it starts with an uppercase letter.
const MaxRetries = 3             // Exported constant
const defaultTimeoutSeconds = 30 // Unexported constant

// 5. Acronyms:
//    - Treat acronyms (like URL, ID, HTTP, API, JSON) as words. Keep them consistently cased.
//    - Conventionally, they are all uppercase if they are short and common.
//    - Examples: `userID`, `apiClient`, `parseURL`, `ServeHTTP`, `loadJSONConfig`.
//      If `ID` is at the start of an exported identifier: `IDProcessor`.
//      If `id` is at the start of an unexported identifier: `idProcessor`.

type UserProfile struct {
	UserID    string
	AuthToken string // Not Authtoken or AuthToken
}

func fetchAPIResponse(apiURL string) string {
	// ... logic to fetch from apiURL ...
	return "response from " + apiURL
}

// 6. Receiver Names (for methods on types):
//    - Should be short (often one or two letters) and representative of the type.
//    - Be consistent within a type's methods.
//    - Example: `e` for `*Employee`, `p` for `*UserProfile`.

func (e *Employee) GetFullName() string { // `e` is the receiver
	return e.FirstName + " " + e.LastName
}

func (e *Employee) setCompanyID(id string) { // `e` is the receiver
	e.companyID = id // Accessing unexported field
}

// 7. Getters and Setters:
//   - Go doesn't automatically generate them.
//   - If you have an unexported field (e.g., `employee.companyID`), an exported getter
//     might be `CompanyID()` and a setter `SetCompanyID()`.
//   - Getter for `companyID`:
func (e *Employee) CompanyID() string {
	return e.companyID
}

// 8. Avoid Stutter:
//    - If a package is named `user`, avoid `user.User`. Prefer `user.Profile` or `user.Account`.
//    - Example: `strings.Reader` (not `strings.StringReader`, though `strings.NewReader` is a factory function).

func main() {
	fmt.Println("--- Go Naming Conventions Demonstration ---")

	// Using an exported variable and function
	fmt.Println(ExportedVariable)
	area := CalculateArea(10, 5)
	fmt.Printf("Calculated Area (exported function): %d\n", area)

	// Using an unexported variable and function (possible because we are in the same package)
	fmt.Println(unexportedVariable)
	perimeter := calculatePerimeter(10, 5)
	fmt.Printf("Calculated Perimeter (unexported function): %d\n", perimeter)

	employee := Employee{FirstName: "John", LastName: "Doe", Age: 30}
	employee.setCompanyID("comp123") // Using an unexported method to set an unexported field

	fmt.Printf("Employee: %s, Age: %d, Company ID (via getter): %s\n",
		employee.GetFullName(), employee.Age, employee.CompanyID())

	profile := UserProfile{UserID: "usr_789", AuthToken: "secret_token_abc"}
	fmt.Printf("User Profile ID: %s, Token: %s\n", profile.UserID, profile.AuthToken)

	apiResponse := fetchAPIResponse("https://api.example.com/data")
	fmt.Println("API Response:", apiResponse)

	fmt.Println("\nReview the comments in this file for detailed explanations of Go naming conventions.")
}
