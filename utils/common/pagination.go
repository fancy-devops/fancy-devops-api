package common

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type Pagination struct {
}

func NewPagination() *Pagination {
	return &Pagination{}
}

func (obj *Pagination) GetPager(c *gin.Context) (skip int, size int) {
	skip = 0
	pageIndex, _ := com.StrTo(c.Query("pageIndex")).Int()
	size, _ = com.StrTo(c.Query("pageSize")).Int()
	if pageIndex > 0 {
		skip = (pageIndex - 1) * size
	}

	return skip, size
}
