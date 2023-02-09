package handler

import (
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
)

type Response struct {
	/*
		FilterUserID    string   `json:"userId,omitempty"`
		FilterBlogID    string   `json:"blogId,omitempty"`
		FilterInterests []string `json:"interests,omitempty"`

	*/
	Filter     Filter      `json:"filter,omitempty"`
	Pagination *Pagination `json:"pagination"`
	//Results    []*Result   `json:"results"`
	Results *newsletter.Result[*newsletter.Subscription] `json:"results"`
}

type ErrorResponse struct {
	Err     string `json:"error"`
	Message string `json:"message"`
}

type Result struct {
	UserID    string   `json:"userId"`
	BlogID    string   `json:"blogId"`
	Interests []string `json:"interests"`
}
