package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/diffpal/demo/internal/orders"
)

type Handler struct {
	service *orders.Service
}

func New(service *orders.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/healthz":
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	case r.Method == http.MethodPost && r.URL.Path == "/orders":
		h.createOrder(w, r)
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/orders/"):
		h.getOrder(w, r)
	default:
		writeError(w, http.StatusNotFound, "not found")
	}
}

type createOrderRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func (h *Handler) createOrder(w http.ResponseWriter, r *http.Request) {
	userID, ok := requestingUser(w, r)
	if !ok {
		return
	}
	var request createOrderRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, "invalid order request")
		return
	}
	order, err := h.service.Create(r.Context(), orders.CreateInput{
		UserID: userID, ProductID: request.ProductID, Quantity: request.Quantity,
	})
	if err != nil {
		switch {
		case errors.Is(err, orders.ErrInvalidProduct), errors.Is(err, orders.ErrInvalidQuantity):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "could not create order")
		}
		return
	}
	writeJSON(w, http.StatusCreated, order)
}

func (h *Handler) getOrder(w http.ResponseWriter, r *http.Request) {
	userID, ok := requestingUser(w, r)
	if !ok {
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/orders/")
	order, err := h.service.Get(r.Context(), id)
	if errors.Is(err, orders.ErrNotFound) {
		writeError(w, http.StatusNotFound, "order not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not load order")
		return
	}
	if order.UserID != userID {
		writeError(w, http.StatusNotFound, "order not found")
		return
	}
	writeJSON(w, http.StatusOK, order)
}

func requestingUser(w http.ResponseWriter, r *http.Request) (string, bool) {
	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		writeError(w, http.StatusUnauthorized, "X-User-ID is required")
		return "", false
	}
	return userID, true
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
