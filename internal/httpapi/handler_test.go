package httpapi

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/diffpal/demo/internal/orders"
)

func newHandler() *Handler {
	clock := func() time.Time { return time.Date(2026, 8, 12, 0, 0, 0, 0, time.UTC) }
	return New(orders.NewServiceWithClock(orders.NewMemoryRepository(), clock))
}

func TestHealthz(t *testing.T) {
	recorder := httptest.NewRecorder()
	newHandler().ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/healthz", nil))
	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusOK)
	}
}

func TestCreateOrder(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(`{"product_id":"mouse","quantity":2}`))
	request.Header.Set("X-User-ID", "user-123")
	newHandler().ServeHTTP(recorder, request)
	if recorder.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d: %s", recorder.Code, http.StatusCreated, recorder.Body.String())
	}
	if got := recorder.Body.String(); !bytes.Contains([]byte(got), []byte(`"unit_price_cents":3500`)) {
		t.Fatalf("response = %s, want server-side price", got)
	}
}

func TestCreateOrderRejectsInvalidProductAndQuantity(t *testing.T) {
	for _, body := range []string{
		`{"product_id":"missing","quantity":1}`,
		`{"product_id":"keyboard","quantity":0}`,
	} {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(body))
		request.Header.Set("X-User-ID", "user-123")
		newHandler().ServeHTTP(recorder, request)
		if recorder.Code != http.StatusBadRequest {
			t.Fatalf("body %s: status = %d, want %d", body, recorder.Code, http.StatusBadRequest)
		}
	}
}

func TestGetOrderByOwner(t *testing.T) {
	handler := newHandler()
	create := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBufferString(`{"product_id":"keyboard","quantity":1}`))
	create.Header.Set("X-User-ID", "user-123")
	handler.ServeHTTP(httptest.NewRecorder(), create)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/orders/ord-000001", nil)
	request.Header.Set("X-User-ID", "user-123")
	handler.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d: %s", recorder.Code, http.StatusOK, recorder.Body.String())
	}
}
