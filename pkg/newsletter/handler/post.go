package handler

import (
	"net/http"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/repository"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/service"
	"github.com/gin-gonic/gin"
)

// nolint:lll // godoc
// Post handler to create a new subscriptions
// @Summary Create subscriptions
// @Tags subscriptions
// @Description Create subscriptions
// @Accept  json
// @Produce  json
// @Param Subscription body newsletter.Subscription true "create subscriptions"
// @Success 200 {string} Success
// @Router /subscriptions [post]
// nolint:gocyclo //error checking branches
func (h *handler) Post(ctx *gin.Context) {
	r := newsletter.Subscription{}

	letterService := service.Must(repository.Must())

	if err := ctx.BindJSON(&r); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	if err := letterService.Post(&r); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, r)
}
