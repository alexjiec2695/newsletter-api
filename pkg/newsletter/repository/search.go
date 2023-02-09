package repository

import (
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/utils"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	uuid "github.com/google/uuid"
)

func (r *repository) Search(userID uuid.UUID, blogID uuid.UUID) ([]*newsletter.Subscription, error) {
	data := subscriptionDBModelToSubscription(r.data)
	data = r.search(data, userID, blogID)
	return data, nil
}

func (r *repository) search(data []*newsletter.Subscription, userID uuid.UUID, blogID uuid.UUID) []*newsletter.Subscription {
	if userID.String() == uuid.Nil.String() && blogID.String() == uuid.Nil.String() {
		return data
	}

	data = utils.Filter(data, func(el *newsletter.Subscription) bool {

		if userID.String() != uuid.Nil.String() && blogID.String() != uuid.Nil.String() {
			return userID.String() == el.UserID.String() && blogID.String() == el.BlogID.String()
		}

		if userID.String() != uuid.Nil.String() && userID.String() == el.UserID.String() {
			return true
		}
		if blogID.String() != uuid.Nil.String() && blogID.String() == el.BlogID.String() {
			return true
		}
		return false
	})

	return data
}
