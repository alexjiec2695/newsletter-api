package newsletter

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	Get(ctx context.Context, UserID uuid.UUID, BlogID uuid.UUID) (*Result[*Subscription], error)
	Post(s *Subscription) error
}
