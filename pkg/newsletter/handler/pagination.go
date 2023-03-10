package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

type Pagination struct {
	Page             int    `json:"page" form:"page"`
	NumberOfPages    int    `json:"numberOfPages"`
	PaginationString string `json:"paginationString"`
	MaxPageSize      int    `json:"maxPageSize" form:"maxPageSize"`
	TotalElements    int    `json:"totalElements"`
}

func (p *Pagination) Get(ctx *gin.Context) (pagination *Pagination, err error) {
	var pag = new(Pagination)

	if err = ctx.BindQuery(pag); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return nil, err
	}
	ctx.Set("page", pag.Page)
	ctx.Set("maxPageSize", pag.MaxPageSize)

	return pag, nil
}

func (p *Pagination) New(totalElements int) *Pagination {
	var numberOfPages = CountPages(p.MaxPageSize, totalElements)
	return &Pagination{
		Page:             p.Page,
		PaginationString: fmt.Sprintf("%v/%v", p.Page, numberOfPages),
		MaxPageSize:      p.MaxPageSize,
		NumberOfPages:    CountPages(p.MaxPageSize, totalElements),
		TotalElements:    totalElements,
	}
}

func CountPages(maxPageSize int, totalElements int) int {
	if totalElements == 0 {
		return 0
	}

	if totalElements <= maxPageSize {
		return 1
	}
	return int(math.Round(float64(totalElements) / float64(maxPageSize)))
}
