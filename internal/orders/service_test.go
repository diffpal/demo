package orders

import (
	"context"
	"testing"
	"time"
)

func TestCreateUsesCatalogPrice(t *testing.T) {
	now := func() time.Time { return time.Date(2026, 8, 12, 0, 0, 0, 0, time.UTC) }
	service := NewServiceWithClock(NewMemoryRepository(), now)

	order, err := service.Create(context.Background(), CreateInput{
		UserID: "user-123", ProductID: "keyboard", Quantity: 2,
	})
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if order.ID != "ord-000001" || order.UnitPriceCents != 7500 || order.TotalCents != 15000 {
		t.Fatalf("Create() order = %+v, want catalog-backed order", order)
	}
	if !order.CreatedAt.Equal(now()) {
		t.Fatalf("Create() CreatedAt = %v, want %v", order.CreatedAt, now())
	}
}

func TestCreateRejectsInvalidInput(t *testing.T) {
	service := NewService(NewMemoryRepository())
	cases := []CreateInput{
		{UserID: "user-123", ProductID: "missing", Quantity: 1},
		{UserID: "user-123", ProductID: "mouse", Quantity: 0},
	}
	for _, input := range cases {
		if _, err := service.Create(context.Background(), input); err == nil {
			t.Fatalf("Create(%+v) error = nil", input)
		}
	}
}
