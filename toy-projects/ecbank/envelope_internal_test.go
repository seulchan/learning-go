// Package ecbank_test contains internal tests for XML parsing and rate calculation logic.
package ecbank

import (
	"errors"
	money "learning-go/moneyconverter"
	"strings"
	"testing"
)

// TestReadRateFromResponse tests the entire process of reading and parsing rates from an XML response.
func TestReadRateFromResponse(t *testing.T) {
	t.Run("Successful USD to RON conversion", func(t *testing.T) {
		xmlData := `<?xml version="1.0" encoding="UTF-8"?>
<gesmes:Envelope xmlns:gesmes="http://www.gesmes.org/xml/2002-08-01" xmlns="http://www.ecb.int/vocabulary/2002-08-01/eurofxref">
	<Cube>
		<Cube time='2023-10-27'>
			<Cube currency='USD' rate='1.25'/>
			<Cube currency='JPY' rate='150.0'/>
			<Cube currency='RON' rate='5.0'/>
		</Cube>
	</Cube>
</gesmes:Envelope>`
		reader := strings.NewReader(xmlData)

		// USD to RON: (EUR/RON) / (EUR/USD) = 5.0 / 1.25 = 4
		expectedRate := money.ExchangeRate(mustParseDecimal(t, "4"))
		rate, err := readRateFromResponse("USD", "RON", reader)

		if err != nil {
			t.Fatalf("readRateFromResponse failed: %v", err)
		}
		if rate != expectedRate {
			t.Errorf("expected rate %v, got %v", expectedRate, rate)
		}
	})

	t.Run("Source currency not found", func(t *testing.T) {
		xmlData := `<?xml version="1.0" encoding="UTF-8"?><gesmes:Envelope><Cube><Cube>
			<Cube currency='USD' rate='1.25'/>
		</Cube></Cube></gesmes:Envelope>`
		reader := strings.NewReader(xmlData)

		_, err := readRateFromResponse("XYZ", "USD", reader) // XYZ is not in XML
		if !errors.Is(err, ErrExchangeRateNotFound) {
			t.Errorf("expected error %v, got %v", ErrExchangeRateNotFound, err)
		}
	})

	t.Run("Target currency not found", func(t *testing.T) {
		xmlData := `<?xml version="1.0" encoding="UTF-8"?><gesmes:Envelope><Cube><Cube>
			<Cube currency='USD' rate='1.25'/>
		</Cube></Cube></gesmes:Envelope>`
		reader := strings.NewReader(xmlData)

		_, err := readRateFromResponse("USD", "XYZ", reader) // XYZ is not in XML
		if !errors.Is(err, ErrExchangeRateNotFound) {
			t.Errorf("expected error %v, got %v", ErrExchangeRateNotFound, err)
		}
	})

	t.Run("Malformed XML", func(t *testing.T) {
		xmlData := `<?xml version="1.0" encoding="UTF-8"?><MalformedXML>`
		reader := strings.NewReader(xmlData)

		_, err := readRateFromResponse("USD", "EUR", reader)
		if !errors.Is(err, ErrUnexpectedFormat) {
			t.Errorf("expected error %v, got %v", ErrUnexpectedFormat, err)
		}
	})

	t.Run("EUR to USD conversion", func(t *testing.T) {
		xmlData := `<?xml version="1.0" encoding="UTF-8"?><gesmes:Envelope><Cube><Cube>
			<Cube currency='USD' rate='1.25'/>
		</Cube></Cube></gesmes:Envelope>`
		reader := strings.NewReader(xmlData)

		// EUR to USD: USD rate is 1.25 (meaning 1 EUR = 1.25 USD)
		// So, EUR to USD rate is 1.25
		expectedRate := money.ExchangeRate(mustParseDecimal(t, "1.25"))
		rate, err := readRateFromResponse("EUR", "USD", reader)

		if err != nil {
			t.Fatalf("readRateFromResponse failed for EUR to USD: %v", err)
		}
		if rate != expectedRate {
			t.Errorf("expected rate %v for EUR to USD, got %v", expectedRate, rate)
		}
	})

	t.Run("USD to EUR conversion", func(t *testing.T) {
		xmlData := `<?xml version="1.0" encoding="UTF-8"?><gesmes:Envelope><Cube><Cube>
			<Cube currency='USD' rate='1.25'/>
		</Cube></Cube></gesmes:Envelope>`
		reader := strings.NewReader(xmlData)

		// USD to EUR: 1 / (EUR to USD rate) = 1 / 1.25 = 0.8
		expectedRate := money.ExchangeRate(mustParseDecimal(t, "0.8"))
		rate, err := readRateFromResponse("USD", "EUR", reader)

		if err != nil {
			t.Fatalf("readRateFromResponse failed for USD to EUR: %v", err)
		}
		if rate != expectedRate {
			t.Errorf("expected rate %v for USD to EUR, got %v", expectedRate, rate)
		}
	})

	t.Run("Same currency (USD to USD)", func(t *testing.T) {
		xmlData := `<?xml version="1.0" encoding="UTF-8"?><gesmes:Envelope><Cube><Cube>
			<Cube currency='USD' rate='1.25'/>
		</Cube></Cube></gesmes:Envelope>` // XML content doesn't matter here
		reader := strings.NewReader(xmlData)

		expectedRate := money.ExchangeRate(mustParseDecimal(t, "1"))
		rate, err := readRateFromResponse("USD", "USD", reader)
		if err != nil {
			t.Fatalf("readRateFromResponse failed for same currency: %v", err)
		}
		if rate != expectedRate {
			t.Errorf("expected rate %v for USD to USD, got %v", expectedRate, rate)
		}
	})
}
