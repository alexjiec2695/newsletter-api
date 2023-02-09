package repository

import (
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
)

func (r *repository) Save(element *newsletter.Subscription) error {
	r.data = append(r.data, mapperSubscriptionToModel(element))
	return nil
}
