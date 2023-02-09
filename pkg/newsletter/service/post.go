package service

import (
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
)

func (s *service) Post(element *newsletter.Subscription) error {
	return s.repo.Save(element)
}
