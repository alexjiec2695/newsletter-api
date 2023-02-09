package newsletter

import (
	"github.com/google/uuid"
)

type Repository interface {
	Search(userID uuid.UUID, blogID uuid.UUID) ([]*Subscription, error)
	Save(s *Subscription) error
}
