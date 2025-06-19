// Package bookstore provides types and functions for managing a collection of books.
// It's designed as a simple example to demonstrate basic Go concepts like structs,
// maps, methods, error handling, and constants for a beginner tutorial.
package bookstore

import (
	"errors"
	"fmt"
	// We'll use sort later if we need to order books from the map.
)

// Category represents the genre or subject of a book.
// Using a custom integer type makes the code more readable and type-safe
// compared to using raw integers directly.
type Category int

// These constants define the specific categories available in our bookstore.
// `iota` is a special Go constant that simplifies the definition of incrementing constants.
const (
	CategoryAutobiography Category = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

type Book struct {
	// Title is the title of the book. It's exported (starts with uppercase)
	// so it can be accessed from other packages.
	Title string
	// Author is the name of the book's author. Exported.
	Author string
	// Copies is the number of copies of this book currently in stock. Exported.
	Copies int
	// ID is a unique identifier for the book. Exported.
	ID int
	// PriceCents is the price of the book in cents (to avoid floating-point issues with money). Exported.
	PriceCents int
	// DiscountPercent is the discount applied to the book's price, as a percentage. Exported.
	DiscountPercent int
	// category is the book's genre. It's unexported (starts with lowercase)
	// meaning it can only be accessed or modified within the `bookstore` package.
	// We provide exported methods (SetCategory, Category) to interact with it.
	category Category
	// isFiction is an example of another unexported field.
	isFiction bool
}

type Catalog map[int]Book

var validCategory = map[Category]bool{
	CategoryAutobiography:     true,
	CategoryLargePrintRomance: true,
	CategoryParticlePhysics:   true,
}

// Buy simulates the purchase of a single copy of a book.
// It takes a Book value as input. IMPORTANT: Go passes structs by value,
// meaning this function operates on a *copy* of the original book.
// The original book outside this function will NOT be modified.
// This is a key concept in Go to understand!
// It returns the updated Book value (the copy) and an error if no copies are left.
func Buy(b Book) (Book, error) {
	// Check if there are any copies available.
	if b.Copies == 0 {
		// If not, return an empty Book struct and a specific error.
		// errors.New creates a simple error message.
		return Book{}, errors.New("no copies left")
	}
	// Decrement the number of copies in the *copy* of the book.
	b.Copies--
	// Return the updated copy of the book and nil (indicating no error).
	return b, nil
}

// AddBook adds a book to the catalog.
// It takes a pointer receiver `*Catalog` because it needs to modify the original map
// (adding a new entry). Maps in Go are reference types, but modifying the map itself
// (like adding or deleting keys) requires a pointer if the map is passed to a function/method.
// It returns an error if a book with the same ID already exists.
func (c Catalog) AddBook(book Book) error {
	// Check if a book with this ID already exists in the catalog.
	if _, exists := c[book.ID]; exists {
		// If it exists, return an error. fmt.Errorf is used to create formatted errors.
		return fmt.Errorf("book with ID %d already exists", book.ID)
	}
	// Add the book to the catalog map.
	c[book.ID] = book
	// Return nil to indicate success.
	return nil
}

// GetAllBooks retrieves all books from the catalog as a slice.
// It takes a value receiver `Catalog` because it only needs to read from the map, not modify it.
// Note: Iterating over a map in Go does not guarantee any specific order.
func (c Catalog) GetAllBooks() []Book {
	// Create an empty slice to store the books.
	result := []Book{}
	// Iterate over the values (books) in the catalog map.
	for _, b := range c {
		// Append each book to the result slice.
		result = append(result, b)
	}
	// Return the slice containing all books.
	// For consistent test results, you might want to sort this slice,
	// but the method itself doesn't guarantee order. Sorting is often done by the caller or in tests.
	return result
}

// GetBook retrieves a single book from the catalog by its ID.
// It takes a value receiver `Catalog` as it only reads from the map.
// It returns the found Book and nil, or an empty Book and an error if the ID is not found.
func (c Catalog) GetBook(ID int) (Book, error) {
	// Look up the book in the map using the ID as the key.
	// The map lookup returns the value (the Book) and a boolean indicating if the key was found.
	b, ok := c[ID]
	// Check if the key was NOT found (`!ok`).
	if !ok {
		// If not found, return an empty Book struct and a formatted error.
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	// If found, return the Book and nil (indicating success).
	return b, nil
}

// NetPriceCents calculates the final price of the book after applying the discount.
// It takes a value receiver `Book` as it only reads the book's fields.
func (b Book) NetPriceCents() int {
	// Calculate the amount of saving based on the discount percentage.
	saving := b.PriceCents * b.DiscountPercent / 100
	// Subtract the saving from the original price to get the net price.
	return b.PriceCents - saving
}

func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		return fmt.Errorf("negative price %d", price)
	}
	b.PriceCents = price
	return nil
}

// SetCategory sets the category for the book.
// It takes a pointer receiver `*Book` because it needs to modify the original book's `category` field.
// It validates the provided category against the list of valid categories.
func (b *Book) SetCategory(category Category) error {
	// Check if the provided category exists in our `validCategory` map.
	if !validCategory[category] {
		// If the category is not valid, return a formatted error.
		return fmt.Errorf("unknown cateogry %v", category)
	}
	// If the category is valid, update the book's unexported `category` field.
	b.category = category
	// Return nil to indicate success.
	return nil
}

// Category is a getter method that returns the book's category.
// It takes a value receiver `Book` as it only reads the `category` field.
func (b Book) Category() Category {
	return b.category
}
