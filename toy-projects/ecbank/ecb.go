// Package ecbank provides a client to fetch exchange rates from the European Central Bank (ECB).
package ecbank

import (
	"errors"
	"fmt"
	money "learning-go/moneyconverter"
	"net/http"
	"net/url"
	"time"
)

// ECBError defines a custom error type for errors originating from the ecbank package.
// This allows for more specific error checking using errors.Is or errors.As.
type ECBError string

// Error implements the error interface for ECBError.
func (e ECBError) Error() string {
	return string(e)
}

// Predefined error values for common issues encountered when interacting with the ECB service.
const (
	ErrCallingServer        = ECBError("ECB client: error calling server")
	ErrTimeout              = ECBError("ECB client: timed out when waiting for response")
	ErrUnexpectedFormat     = ECBError("ECB client: unexpected response format from server")
	ErrExchangeRateNotFound = ECBError("ECB client: couldn't find the requested exchange rate")
	ErrClientSide           = ECBError("ECB client: client-side error (4xx) when contacting ECB")
	ErrServerSide           = ECBError("ECB client: server-side error (5xx) when contacting ECB")
	ErrUnknownStatusCode    = ECBError("ECB client: unknown status code received from ECB")
)

// Client is used to interact with the European Central Bank's exchange rate service.
// It holds an HTTP client configured for making requests.
type Client struct {
	httpClient *http.Client
	ratesURL   string // URL for fetching exchange rates, allowing for easier testing.
}

// NewClient creates and returns a new ECB Client.
// It takes a timeout duration, which is applied to HTTP requests made by the client.
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: &http.Client{Timeout: timeout},
		// This is the official daily Euro foreign exchange reference rates XML feed.
		ratesURL: "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml",
	}
}

// FetchExchangeRate fetches today's ExchangeRate and returns it.
// It communicates with the ECB service, parses the response, and calculates the rate.
func (c Client) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {
	// Make an HTTP GET request to the ECB's rates URL.
	resp, err := c.httpClient.Get(c.ratesURL)
	if err != nil {
		// Check if the error is a URL error (e.g., network issue, DNS problem).
		var urlErr *url.Error
		// errors.As checks if 'err' is (or wraps) a *url.Error and assigns it to urlErr.
		if errors.As(err, &urlErr) && urlErr.Timeout() {
			// If the error is specifically a timeout, wrap it with our custom ErrTimeout.
			// Wrapping (using %w) preserves the original error for further inspection if needed.
			return money.ExchangeRate{}, fmt.Errorf("%w: %v", ErrTimeout, urlErr)
		}
		// For other types of errors during the GET request, wrap them with ErrCallingServer.
		return money.ExchangeRate{}, fmt.Errorf("%w: %v", ErrCallingServer, err)
	}
	// defer ensures that resp.Body.Close() is called just before the FetchExchangeRate function returns.
	// This is crucial for releasing resources and preventing memory leaks.
	defer resp.Body.Close()

	// Check the HTTP status code of the response.
	if err = checkStatusCode(resp.StatusCode); err != nil {
		// If the status code indicates an error (e.g., 404 Not Found, 500 Server Error), return the error.
		return money.ExchangeRate{}, err
	}

	rate, err := readRateFromResponse(source.Code(), target.Code(), resp.Body)
	if err != nil {
		return money.ExchangeRate{}, err
	}
	// If everything is successful, return the fetched rate.
	return rate, nil
}

// checkStatusCode examines the HTTP status code and returns a specific error if the code indicates a problem.
// It returns nil if the status code is http.StatusOK (200).
func checkStatusCode(statusCode int) error {
	switch {
	case statusCode == http.StatusOK: // 200 OK - success
		return nil
	case statusCode >= 400 && statusCode < 500: // 4xx Client Errors (e.g., 404 Not Found)
		// %w wraps the ErrClientSide with additional context (the specific status code).
		return fmt.Errorf("%w, status code: %d", ErrClientSide, statusCode)
	case statusCode >= 500 && statusCode < 600: // 5xx Server Errors (e.g., 500 Internal Server Error)
		return fmt.Errorf("%w, status code: %d", ErrServerSide, statusCode)
	default: // Any other status codes that are not 200 or in the 4xx/5xx ranges.
		return fmt.Errorf("%w, status code: %d", ErrUnknownStatusCode, statusCode)
	}
}

// Note: The httpStatusClass function and related constants (clientErrorClass, serverErrorClass)
// were removed in favor of direct range checks in checkStatusCode for better clarity
// for beginners.
