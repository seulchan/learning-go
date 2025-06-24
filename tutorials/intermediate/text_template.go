// --- Go Text Template Tutorial ---
//
// The `text/template` package in Go provides a powerful way to generate text output
// from templates. It's commonly used for creating dynamic content like HTML pages,
// configuration files, or formatted text messages.
//
// This tutorial demonstrates:
// 1. The basic workflow of parsing and executing a template.
// 2. Using `template.Must` for safe initialization.
// 3. An interactive example with multiple named templates and user input.
package main

import (
	"bufio"
	"fmt"
	"log" // Using the log package for handling fatal errors
	"os"
	"strings"
	"text/template"
)

// demonstrateBasicTemplate shows the fundamental steps of using a template:
// 1. Create a new template.
// 2. Parse a template string.
// 3. Execute the template with data, writing to an output stream.
func demonstrateBasicTemplate() {
	fmt.Println("--- 1. Basic Template Demonstration ---")

	// The template string. `{{.UserName}}` is an "action". The `.` refers to the
	// data object passed during execution, and `UserName` is a field of that object.
	templateString := "Welcome, {{.UserName}}! This is a basic template.\n"

	// template.New creates a new, named template.
	// .Parse associates the template string with the template.
	// This returns a template object and an error, which we must check.
	tmpl, err := template.New("basic-welcome").Parse(templateString)
	if err != nil {
		// If parsing fails (e.g., due to a syntax error like `{{.Name`),
		// the program cannot continue as intended. `log.Fatalf` prints the
		// error message and exits the program with a non-zero status code.
		log.Fatalf("Failed to parse template: %v", err)
	}

	// This is the data we want to "fill in" the template with.
	// Using a struct for data is a common and type-safe practice in Go.
	data := struct{ UserName string }{
		UserName: "Alice",
	}

	// tmpl.Execute applies the parsed template to the data (`data`) and writes
	// the output to a writer. Here, `os.Stdout` is the standard output (your console).
	fmt.Print("Executing template with a struct: ")
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		// An error here could be an I/O issue (e.g., cannot write to the console).
		log.Fatalf("Failed to execute template: %v", err)
	}
	fmt.Println()
}

// demonstrateMust shows a more concise way to parse templates when a failure
// should be considered a fatal, unrecoverable error.
func demonstrateMust() {
	fmt.Println("--- 2. Using template.Must for Guaranteed Parsing ---")

	// The template string is known at compile time. If it's invalid, it's a
	// programmer error, and the program should not start.
	templateString := "Hello, {{.UserName}}! This template was parsed with Must.\n"

	// template.Must is a helper function that wraps a call to a function like
	// template.New(...).Parse(...).
	// If the wrapped call returns an error, `Must` will panic.
	// This is useful for initializing global template variables, where you want
	// the program to fail immediately if the templates can't be parsed.
	tmpl := template.Must(template.New("must-example").Parse(templateString))

	// The data can also be a map. This is flexible but lacks the type safety
	// of a struct. The keys of the map correspond to the names in the template.
	data := map[string]string{
		"UserName": "Bob",
	}

	// Execute the template as before.
	fmt.Print("Executing template parsed with Must: ")
	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		// An execution error can still happen (e.g., I/O error), so we check it.
		log.Fatalf("Failed to execute template: %v", err)
	}
	fmt.Println()
}

// runInteractiveTemplateSelector demonstrates a more complex use case with multiple
// named templates and interactive user input.
func runInteractiveTemplateSelector() {
	fmt.Println("--- 3. Interactive Template Selector ---")

	// It's efficient to parse all related templates into a single template object.
	// We use the `{{define "name"}}...{{end}}` syntax to create multiple named
	// templates within a single string. This is a common and clean pattern.
	const allTemplateStrings = `
{{define "welcome"}}Welcome, {{.UserName}}! We're glad you joined.
{{end}}
{{define "notification"}}Hello {{.UserName}}, you have a new notification: {{.Message}}
{{end}}
{{define "error"}}Oops! An error occurred for {{.UserName}}: {{.Message}}
{{end}}
`
	// We parse all our defined templates at once. `template.Must` is used here
	// because if these core templates fail to parse, the program is in an
	// invalid state and should not continue.
	// The name passed to `New` ("interactive-set") is for identification.
	templates := template.Must(template.New("interactive-set").Parse(allTemplateStrings))

	// We'll use a bufio.Reader to read full lines of text from the user.
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("To begin, please enter your name: ")
	userName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read name: %v", err)
	}
	// The name read from the console includes the newline character, so we trim it.
	userName = strings.TrimSpace(userName)

	// An infinite loop to show the menu until the user chooses to exit.
	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Show Welcome Message")
		fmt.Println("2. Show Notification")
		fmt.Println("3. Show Error Message")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		choice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read choice, please try again.")
			continue
		}
		choice = strings.TrimSpace(choice)

		// This struct will hold the data for our notification and error templates.
		type MessageData struct {
			UserName string
			Message  string
		}

		switch choice {
		case "1":
			// For the welcome message, we only need the UserName.
			// We can pass an anonymous struct for this simple case.
			data := struct{ UserName string }{userName}
			err = templates.ExecuteTemplate(os.Stdout, "welcome", data)
		case "2":
			fmt.Print("Enter your notification message: ")
			message, readErr := reader.ReadString('\n')
			if readErr != nil {
				fmt.Println("Failed to read message, please try again.")
				continue
			}
			data := MessageData{
				UserName: userName,
				Message:  strings.TrimSpace(message),
			}
			err = templates.ExecuteTemplate(os.Stdout, "notification", data)
		case "3":
			fmt.Print("Enter your error message: ")
			message, readErr := reader.ReadString('\n')
			if readErr != nil {
				fmt.Println("Failed to read message, please try again.")
				continue
			}
			data := MessageData{
				UserName: userName,
				Message:  strings.TrimSpace(message),
			}
			err = templates.ExecuteTemplate(os.Stdout, "error", data)
		case "4":
			fmt.Println("Exiting...")
			return // Exit the function, which ends the program.
		default:
			fmt.Println("Invalid choice. Please try again.")
			continue // Skip the rest of the loop and show the menu again.
		}

		// After the switch, check if there was an error executing the template.
		if err != nil {
			fmt.Println("Error executing template:", err)
		}
	}
}

// main is the entry point of our program. It calls the demonstration functions
// in order to present the tutorial concepts sequentially.
func main() {
	fmt.Println("====== Go Text Template Tutorial ======")
	demonstrateBasicTemplate()
	demonstrateMust()
	runInteractiveTemplateSelector()
	fmt.Println("\n====== End of Tutorial ======")
}
