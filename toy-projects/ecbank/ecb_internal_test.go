package ecbank

import (
	"errors"
	"fmt"
	money "learning-go/moneyconverter"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestEuroCentralBank_FetchExchangeRate_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `<?xml version="1.0" encoding="UTF-8"?><gesmes:Envelope><Cube><Cube>
			<Cube currency='USD' rate='2'/>
			<Cube currency='RON' rate='6'/>
		</Cube></Cube></gesmes:Envelope>`)
	}))
	defer ts.Close()

	// Create client using NewClient and override ratesURL for the test server
	ecb := NewClient(time.Second)
	ecb.ratesURL = ts.URL

	got, err := ecb.FetchExchangeRate(mustParseCurrency(t, "USD"), mustParseCurrency(t, "RON"))
	// Expected rate is 3 (RON rate 6 / USD rate 2)
	// money.ExchangeRate is an alias for money.Decimal
	want := money.ExchangeRate(mustParseDecimal(t, "3"))

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if got != want {
		t.Errorf("FetchExchangeRate() got = %v, want %v", got, want)
	}
}

func TestEuroCentralBank_FetchExchangeRate_Timeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) // Sleep longer than client timeout
	}))
	defer ts.Close()

	// Create client using NewClient and override ratesURL for the test server
	ecb := NewClient(time.Second) // Client timeout is 1 second
	ecb.ratesURL = ts.URL

	_, err := ecb.FetchExchangeRate(mustParseCurrency(t, "USD"), mustParseCurrency(t, "RON"))
	if !errors.Is(err, ErrTimeout) {
		t.Errorf("unexpected error: %v, expected %v", err, ErrTimeout)
	}
}

func mustParseCurrency(t *testing.T, code string) money.Currency {
	t.Helper()

	currency, err := money.ParseCurrency(code)
	if err != nil {
		t.Fatalf("cannot parse currency %s code", code)
	}

	return currency
}

func mustParseDecimal(t *testing.T, decimal string) money.Decimal {
	t.Helper()

	dec, err := money.ParseDecimal(decimal)
	if err != nil {
		t.Fatalf("cannot parse decimal %s", decimal)
	}

	return dec
}
