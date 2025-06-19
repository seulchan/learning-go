// Package bookstore_test contains unit tests for the bookstore package.
// It's in a separate package (`_test`) to ensure it only tests the exported
// functions and types, mimicking how an external user would interact with the package.
package bookstore_test

import (
	"bookstore" // Import the package we are testing.
	"sort"      // Used for sorting slices in tests for consistent comparison.
	"testing"   // Go's built-in testing package.

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// TestBuy tests the Buy function.
// It checks if buying a book correctly decrements the copies count.
func TestBuy(t *testing.T) {
	// t.Parallel() allows this test to run in parallel with other tests marked with t.Parallel().
	// This can speed up test execution.
	t.Parallel()

	// Define the initial state of the book.
	initialBook := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}

	// Define the expected number of copies after buying one.
	wantCopies := 1

	// Call the function under test.
	// Remember that Buy operates on a *copy* of the book passed to it.
	resultBook, err := bookstore.Buy(initialBook)

	// Check if there was an unexpected error. t.Fatal will stop the test immediately.
	if err != nil {
		t.Fatalf("Buy(%#v) returned unexpected error: %v", initialBook, err)
	}

	// Get the actual number of copies from the result book.
	gotCopies := resultBook.Copies

	// Assert that the actual copies match the expected copies.
	if wantCopies != gotCopies {
		// t.Errorf logs an error message but allows the test to continue.
		t.Errorf("Buy(%#v): want copies %d, got %d", initialBook, wantCopies, gotCopies)
	}

	// Optional: Demonstrate that the original book was not changed (due to pass-by-value).
	// This is a good point to highlight in the tutorial.
	if initialBook.Copies != 2 {
		t.Errorf("Original book copies changed! Expected 2, got %d. (This demonstrates pass-by-value)", initialBook.Copies)
	}
}

// TestBuyErrorsIfNoCopiesLeft tests the error case for the Buy function.
// It checks if buying a book with zero copies returns an error.
func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()

	// Define a book with zero copies.
	bookWithNoCopies := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}

	// Attempt to buy the book. We expect an error.
	_, err := bookstore.Buy(bookWithNoCopies)

	// Assert that an error was returned.
	if err == nil {
		// If err is nil, the test fails because we expected an error.
		t.Error("want error buying from zero copies, got nil")
	}
}

// TestAddBook tests the AddBook method of the Catalog type.
// It checks if books are added correctly and if adding a duplicate ID returns an error.
func TestAddBook(t *testing.T) {
	t.Parallel()

	// Create an empty catalog.
	catalog := bookstore.Catalog{}

	// Define the books to add.
	book1 := bookstore.Book{ID: 1, Title: "Book One"}
	book2 := bookstore.Book{ID: 2, Title: "Book Two"}
	duplicateBook1 := bookstore.Book{ID: 1, Title: "Another Book One"} // Same ID as book1

	// Test Case 1: Add a new book successfully.
	err := catalog.AddBook(book1)
	if err != nil {
		t.Fatalf("AddBook(%#v) returned unexpected error: %v", book1, err)
	}
	// Check if the book was actually added.
	if _, ok := catalog[book1.ID]; !ok {
		t.Errorf("Book with ID %d was not added to the catalog", book1.ID)
	}

	// Test Case 2: Add another new book successfully.
	err = catalog.AddBook(book2)
	if err != nil {
		t.Fatalf("AddBook(%#v) returned unexpected error: %v", book2, err)
	}
	// Check if the second book was added.
	if _, ok := catalog[book2.ID]; !ok {
		t.Errorf("Book with ID %d was not added to the catalog", book2.ID)
	}

	// Test Case 3: Attempt to add a book with a duplicate ID.
	err = catalog.AddBook(duplicateBook1)
	// We expect an error here.
	if err == nil {
		t.Errorf("AddBook(%#v): want error for duplicate ID, got nil", duplicateBook1)
	}
	// Optional: Check the error message content if needed, but checking for non-nil is often sufficient.
	// if err != nil && !strings.Contains(err.Error(), "already exists") {
	// 	t.Errorf("AddBook(%#v): want error message containing 'already exists', got %v", duplicateBook1, err)
	// }

	// Check that the original book1 was not overwritten by duplicateBook1.
	if addedBook, ok := catalog[book1.ID]; ok && addedBook.Title != book1.Title {
		t.Errorf("Book with ID %d was overwritten by a duplicate add attempt", book1.ID)
	}
}

// TestGetAllBooks tests the GetAllBooks method of the Catalog type.
// It checks if the method returns all books currently in the catalog.
func TestGetAllBooks(t *testing.T) {
	t.Parallel()

	// Define the initial catalog with some books.
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}

	// Define the expected slice of books.
	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}

	// Call the method under test.
	got := catalog.GetAllBooks()

	// Maps in Go do not guarantee iteration order. To compare the resulting slice
	// reliably, we need to sort both the expected and actual slices by a consistent field (like ID).
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})

	// Use go-cmp to compare the slices.
	// cmpopts.IgnoreUnexported(bookstore.Book{}) is necessary because the Book struct
	// has an unexported field (`category`). go-cmp needs this option to compare
	// structs from different packages that have unexported fields.
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		// If they are not equal, use cmp.Diff to get a detailed difference report.
		t.Error(cmp.Diff(want, got))
	}
}

// TestGetBook tests the GetBook method for a valid book ID.
func TestGetBook(t *testing.T) {
	t.Parallel()

	// Define the catalog with books.
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}

	// Define the expected book for ID 2.
	want := bookstore.Book{
		ID:    2,
		Title: "The Power of Go: Tools",
	}

	// Call the method under test with a valid ID.
	got, _ := catalog.GetBook(2)

	// Compare the retrieved book with the expected book using go-cmp.
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

// TestGetBookBadIDReturnsError tests the GetBook method for an invalid book ID.
// It checks if the method correctly returns an error when the ID is not found.
func TestGetBookBadIDReturnsError(t *testing.T) {
	t.Parallel()

	// Create an empty catalog.
	catalog := bookstore.Catalog{}

	// Attempt to get a book with a non-existent ID.
	_, err := catalog.GetBook(999)

	// Assert that an error was returned.
	if err == nil {
		// If err is nil, the test fails. t.Fatal stops the test immediately.
		t.Fatal("want error for non-existent ID, got nil")
	}
}

// TestNetPriceCents tests the NetPriceCents method of the Book type.
// It checks if the discounted price is calculated correctly.
func TestNetPriceCents(t *testing.T) {
	t.Parallel()

	// Define a book with a price and discount.
	b := bookstore.Book{
		Title:           "For the Love of Go",
		PriceCents:      4000,
		DiscountPercent: 25,
	}

	// Calculate the expected net price: 4000 - (4000 * 25 / 100) = 4000 - 1000 = 3000.
	want := 3000

	// Call the method under test.
	got := b.NetPriceCents()

	// Assert that the calculated price matches the expected price.
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

// TestSetPriceCents tests the SetPriceCents method for valid input.
// It checks if the method correctly updates the book's price.
func TestSetPriceCents(t *testing.T) {
	t.Parallel()

	// Define an initial book.
	b := bookstore.Book{
		Title:      "For the Love of Go",
		PriceCents: 4000,
	}

	// Define the new price we want to set.
	want := 3000

	// Call the method under test. SetPriceCents uses a pointer receiver,
	// so it modifies the original `b` variable.
	err := b.SetPriceCents(want)

	// Check for unexpected errors.
	if err != nil {
		t.Fatal(err)
	}

	// Get the updated price from the book.
	got := b.PriceCents

	// Assert that the price was updated correctly.
	if want != got {
		t.Errorf("want updated price %d, got %d", want, got)
	}
}

// TestSetPriceCentsInvalid tests the SetPriceCents method for invalid input (negative price).
// It checks if the method returns an error for invalid input.
func TestSetPriceCentsInvalid(t *testing.T) {
	t.Parallel()

	// Define an initial book.
	b := bookstore.Book{
		Title:      "For the Love of Go",
		PriceCents: 4000,
	}

	// Attempt to set an invalid price (-1).
	err := b.SetPriceCents(-1)

	// Assert that an error was returned.
	if err == nil {
		t.Fatal("want error setting invalid price -1, got nil")
	}
}

// TestSetCategory tests the SetCategory method for valid category inputs.
// It checks if the method correctly updates the book's category using a pointer receiver.
func TestSetCategory(t *testing.T) {
	t.Parallel()

	// Define an initial book.
	b := bookstore.Book{
		Title: "For the Love of Go",
	}

	// Define a slice of valid categories to test.
	cats := []bookstore.Category{
		bookstore.CategoryAutobiography,
		bookstore.CategoryLargePrintRomance,
		bookstore.CategoryParticlePhysics,
	}

	// Iterate through each valid category.
	for _, cat := range cats {
		// Call the method under test. SetCategory uses a pointer receiver,
		// modifying the original `b` variable.
		err := b.SetCategory(cat)

		// Check for unexpected errors.
		if err != nil {
			t.Fatal(err)
		}

		// Get the category using the getter method.
		got := b.Category()

		// Assert that the category was set correctly.
		// Note: Comparing enums (which are int types) directly is fine.
		if cat != got {
			t.Errorf("want category %q, got %q", cat, got)
		}
	}

}

// TestSetCategoryInvalid tests the SetCategory method for invalid category input.
// It checks if the method returns an error for an unknown category.
func TestSetCategoryInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the Love of Go",
	}
	// Attempt to set an invalid category (e.g., 999, which is not defined).
	err := b.SetCategory(999)
	if err == nil {
		t.Fatal("want error for invalid category, got nil")
	}
}
