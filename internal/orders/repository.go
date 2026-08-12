package orders

import (
	"context"
	"errors"
	"sync"
)

var ErrNotFound = errors.New("order not found")

type Repository interface {
	Save(ctx context.Context, order Order) error
	Get(ctx context.Context, id string) (Order, error)
}

type MemoryRepository struct {
	mu     sync.RWMutex
	orders map[string]Order
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{orders: make(map[string]Order)}
}

func (r *MemoryRepository) Save(_ context.Context, order Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[order.ID] = order
	return nil
}

func (r *MemoryRepository) Get(_ context.Context, id string) (Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	order, ok := r.orders[id]
	if !ok {
		return Order{}, ErrNotFound
	}
	return order, nil
}
