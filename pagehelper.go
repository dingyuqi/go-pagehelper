// Package pagehelper -----------------------------
// @file      : pagehelper.go
// @author    : dingyq
// @time      : 2024/7/15 下午3:42
// -------------------------------------------
package pagehelper

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type Pagination struct {
	Page       int
	PageSize   int
	OrderParam string
	Total      int64
	NoPaginate bool
}

type PageResult struct {
	List     any   `json:"rows"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}

func (pagination *Pagination) ToResult(List any) any {
	if pagination == nil || pagination.NoPaginate == true {
		return map[string]any{"rows": List, "total": pagination.Total}
	}
	result := PageResult{}
	result.List = List
	result.Total = pagination.Total
	result.Page = pagination.Page
	result.PageSize = pagination.PageSize
	return result
}

func (pagination *Pagination) Offset() int {
	return (pagination.Page - 1) * pagination.PageSize
}

func (pagination *Pagination) Paginate() func(db *gorm.DB) *gorm.DB {
	if pagination == nil {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.Offset()).Limit(pagination.PageSize).Order(pagination.OrderParam)
	}
}

func (pagination *Pagination) Count() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(-1).Limit(-1)
	}
}

func (pagination *Pagination) GetResult(tx *gorm.DB, data any) {
	if pagination.NoPaginate == true {
		tx.Find(data)
		tx.Scopes(pagination.Count()).Count(&pagination.Total)
		return
	}
	tx.Scopes(pagination.Count()).Count(&pagination.Total)
	tx.Scopes(pagination.Paginate()).Find(data)
}

func GetPaginationParam(c *gin.Context) *Pagination {
	pagination, ok := c.Request.Context().Value("pagination").(*Pagination)
	if !ok {
		return nil
	}
	return pagination
}

func PaginationMiddleware(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("pageNum", "-1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "-1"))
	orderParam := c.DefaultQuery("orderByColumn", "") + " " + c.DefaultQuery("isAsc", "")
	if orderParam == " " {
		orderParam = ""
	}
	if page == -1 || pageSize == -1 {
		pagination := &Pagination{
			NoPaginate: true,
		}
		ctx := context.WithValue(c.Request.Context(), "pagination", pagination)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
		return
	}
	pagination := &Pagination{
		Page:       page,
		PageSize:   pageSize,
		OrderParam: orderParam,
	}

	ctx := context.WithValue(c.Request.Context(), "pagination", pagination)
	c.Request = c.Request.WithContext(ctx)

	c.Next()
}
