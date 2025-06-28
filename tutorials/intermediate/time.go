// --- Go `time` Package Tutorial ---
//
// The `time` package in Go provides functionality for measuring and displaying time.
// It's a core package used in almost any application that deals with dates, times,
// durations, or timezones.
//
// This tutorial covers the essential features of the `time` package:
// 1. Getting the current time.
// 2. Creating specific time instances.
// 3. Parsing and formatting time strings.
// 4. Performing time arithmetic (adding/subtracting durations).
// 5. Working with timezones.
// 6. Calculating durations and comparing times.
package main

import (
	"fmt"
	"log"
	"time"
)

// main is the entry point of our program. It calls the demonstration functions.
func main() {
	fmt.Println("--- Go `time` Package Tutorial ---")

	demonstrateCurrentTime()
	demonstrateCreatingTime()
	demonstrateParsingAndFormatting()
	demonstrateTimeArithmetic()
	demonstrateTimezones()
	demonstrateDurationsAndComparisons()

	fmt.Println("\n--- End of `time` Package Tutorial ---")
}

// demonstrateCurrentTime shows how to get the current time and access its components.
func demonstrateCurrentTime() {
	fmt.Println("\n--- 1. Getting the Current Time ---")

	// `time.Now()` returns the current local time as a `time.Time` object.
	currentTime := time.Now()
	fmt.Println("Current local time:", currentTime)

	// The `time.Time` struct has methods to extract individual components.
	fmt.Printf("Year: %d, Month: %s, Day: %d\n",
		currentTime.Year(), currentTime.Month(), currentTime.Day())
	fmt.Printf("Hour: %d, Minute: %d, Second: %d\n",
		currentTime.Hour(), currentTime.Minute(), currentTime.Second())
	fmt.Println("Weekday:", currentTime.Weekday())
}

// demonstrateCreatingTime shows how to create a specific time instance.
func demonstrateCreatingTime() {
	fmt.Println("\n--- 2. Creating a Specific Time ---")

	// `time.Date` lets you create a `time.Time` for a specific date and time.
	// The arguments are: year, month, day, hour, minute, second, nanosecond, and location (timezone).
	// `time.UTC` specifies the Coordinated Universal Time timezone.
	specificTime := time.Date(2024, time.July, 30, 12, 30, 0, 0, time.UTC)
	fmt.Println("A specific time (UTC):", specificTime)
}

// demonstrateParsingAndFormatting shows how to convert between `time.Time` objects and strings.
func demonstrateParsingAndFormatting() {
	fmt.Println("\n--- 3. Parsing and Formatting Time ---")

	// --- Formatting Time to String ---
	// Go uses a unique, memorable layout string for formatting. Instead of symbols like
	// `YYYY-MM-DD`, you write out the specific reference time: Mon Jan 2 15:04:05 MST 2006.
	// `t.Format()` will format the time `t` to look like the layout string.
	now := time.Now()
	// Example 1: A custom, human-readable format.
	customFormat := now.Format("Monday, January 2, 2006 at 3:04 PM")
	fmt.Println("Custom formatted time:", customFormat)

	// Example 2: A standard format like RFC3339, which is common for APIs.
	rfcFormat := now.Format(time.RFC3339)
	fmt.Println("RFC3339 formatted time:", rfcFormat)

	// --- Parsing String to Time ---
	// `time.Parse` does the reverse. It takes a layout string and a time string,
	// and tries to parse the string according to the layout.
	dateString := "2023-10-26"
	layout := "2006-01-02" // The layout must match the format of `dateString`.

	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		// It's crucial to handle the error, as the input string might be invalid.
		log.Fatalf("Failed to parse time string: %v", err)
	}
	fmt.Printf("Parsed time from string \"%s\": %v\n", dateString, parsedTime)
}

// demonstrateTimeArithmetic shows how to add durations and round/truncate time.
func demonstrateTimeArithmetic() {
	fmt.Println("\n--- 4. Time Arithmetic ---")
	now := time.Now()
	fmt.Println("Current time:", now)

	// `time.Duration` represents a span of time. Constants like `time.Hour` make it readable.
	// `t.Add()` returns a new `time.Time` instance with the duration added.
	oneDayLater := now.Add(24 * time.Hour)
	fmt.Println("One day later:", oneDayLater)

	twoHoursAgo := now.Add(-2 * time.Hour)
	fmt.Println("Two hours ago:", twoHoursAgo)

	// `Round()` rounds a time to the nearest multiple of the given duration.
	// Here, it rounds to the nearest hour.
	roundedTime := now.Round(time.Hour)
	fmt.Println("Rounded to nearest hour:", roundedTime)

	// `Truncate()` rounds a time *down* to the multiple of the given duration.
	// Here, it truncates to the beginning of the current hour.
	truncatedTime := now.Truncate(time.Hour)
	fmt.Println("Truncated to the hour:", truncatedTime)
}

// demonstrateTimezones shows how to work with different timezones.
func demonstrateTimezones() {
	fmt.Println("\n--- 5. Working with Timezones ---")

	// `time.LoadLocation` gets a `*time.Location` for a specific timezone.
	// It's good practice to handle the error in case the timezone name is invalid.
	newYorkLocation, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalf("Could not load location: %v", err)
	}

	// Get the current time and convert it to the New York timezone.
	timeInNewYork := time.Now().In(newYorkLocation)
	fmt.Println("Current time in New York:", timeInNewYork)

	// Let's create a UTC time and convert it.
	utcTime := time.Date(2024, 7, 8, 14, 0, 0, 0, time.UTC)
	kolkataLocation, _ := time.LoadLocation("Asia/Kolkata") // Ignoring error for brevity in this example
	timeInKolkata := utcTime.In(kolkataLocation)

	fmt.Printf("The time %v in UTC is %v in Kolkata.\n", utcTime, timeInKolkata)
}

// demonstrateDurationsAndComparisons shows how to calculate the difference between
// times and how to compare them.
func demonstrateDurationsAndComparisons() {
	fmt.Println("\n--- 6. Durations and Comparisons ---")

	// Create two time points.
	startTime := time.Date(2024, time.July, 4, 12, 0, 0, 0, time.UTC)
	endTime := time.Date(2024, time.July, 4, 18, 30, 0, 0, time.UTC)

	// `t2.Sub(t1)` calculates the duration between two times.
	duration := endTime.Sub(startTime)
	fmt.Printf("The duration between %v and %v is %v.\n", startTime, endTime, duration)
	fmt.Printf("That's %.1f hours.\n", duration.Hours())

	// `After`, `Before`, and `Equal` are used for comparisons.
	fmt.Printf("Is endTime after startTime? %t\n", endTime.After(startTime))    // true
	fmt.Printf("Is startTime before endTime? %t\n", startTime.Before(endTime))  // true
	fmt.Printf("Is startTime equal to endTime? %t\n", startTime.Equal(endTime)) // false
}
