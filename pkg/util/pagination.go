package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"gitlab.chad122.top/fancy-devops/fancy-devops-api/pkg/setting"
)

type Pagination struct {
}

func NewPagination() *Pagination {
	return &Pagination{}
}

func (obj *Pagination) GetPageSkip(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
