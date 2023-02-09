package handler

import (
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Filter struct {
	UserID    string                `json:"userId" form:"userId"`
	BlogID    string                `json:"blogId" form:"blogId"`
	Interests []newsletter.Interest `json:"interests"`
}

func (f *Filter) Get(ctx *gin.Context) (filter *Filter, err error) {
	var fill = new(Filter)

	if err = ctx.BindQuery(fill); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return nil, err
	}
	if len(fill.Interests) == 0 {
		fill.Interests = make([]newsletter.Interest, 0)
	}
	return fill, nil
}
