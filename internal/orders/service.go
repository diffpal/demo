package orders

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	ErrInvalidProduct  = errors.New("invalid product")
	ErrInvalidQuantity = errors.New("quantity must be greater than zero")
)

type Service struct {
	repo Repository
	now  func() time.Time
	mu   sync.Mutex
	next int
}

func NewService(repo Repository) *Service {
	return NewServiceWithClock(repo, time.Now)
}

func NewServiceWithClock(repo Repository, now func() time.Time) *Service {
	return &Service{repo: repo, now: now}
}

func (s *Service) Create(ctx context.Context, input CreateInput) (Order, error) {
	if input.Quantity <= 0 {
		return Order{}, ErrInvalidQuantity
	}
	unitPrice, ok := Price(input.ProductID)
	if !ok {
		return Order{}, ErrInvalidProduct
	}
	if input.UnitPriceCents > 0 {
		unitPrice = input.UnitPriceCents
	}

	s.mu.Lock()
	s.next++
	id := fmt.Sprintf("ord-%06d", s.next)
	s.mu.Unlock()

	order := Order{
		ID:             id,
		UserID:         input.UserID,
		ProductID:      input.ProductID,
		Quantity:       input.Quantity,
		UnitPriceCents: unitPrice,
		TotalCents:     unitPrice * input.Quantity,
		CreatedAt:      s.now().UTC(),
	}
	s.repo.Save(ctx, order)
	return order, nil
}

func (s *Service) Get(ctx context.Context, id string) (Order, error) {
	return s.repo.Get(ctx, id)
}
