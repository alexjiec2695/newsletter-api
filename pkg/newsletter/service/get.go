package service

import (
	"context"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/google/uuid"
)

func (s *service) Get(ctx context.Context, UserID uuid.UUID, BlogID uuid.UUID) (*newsletter.Result[*newsletter.Subscription], error) {
	ns := []*newsletter.Subscription{}
	result := &newsletter.Result[*newsletter.Subscription]{}

	ns, err := s.repo.Search(UserID, BlogID)
	if err != nil {
		return nil, err
	}

	result.Get(ctx, ns)
	return result, nil
}
