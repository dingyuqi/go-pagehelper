// Package request -----------------------------
// @file      : request.go
// @author    : dingyq
// @time      : 2024/7/16 下午5:56
// -------------------------------------------
package request

import (
	"strconv"
	"strings"
)

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page          int    `json:"page" form:"page"`                   // 页码
	PageSize      int    `json:"pageSize" form:"pageSize"`           // 每页大小
	Keyword       string `json:"keyword" form:"keyword"`             //关键字
	OrderByColumn string `json:"orderByColumn" form:"orderByColumn"` // 排序列
	IsAsc         string `json:"isAsc" form:"isAsc"`
}

// GetLimitStr get postgres page limit string
func (i PageInfo) GetLimitStr() string {
	var builder strings.Builder
	limit := i.PageSize
	offset := i.PageSize * (i.Page - 1)
	if limit > 0 {
		builder.WriteString(" LIMIT ")
		builder.WriteString(strconv.Itoa(i.PageSize))
	}
	if offset > 0 {
		if i.PageSize >= 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(" OFFSET ")
		builder.WriteString(strconv.Itoa(offset))
	}
	return builder.String()
}

func (i PageInfo) GetOrderStr() string {
	if i.OrderByColumn == "create_time" {
		i.OrderByColumn = "last_time"
	} else if i.OrderByColumn == "confidence" {
		i.OrderByColumn = "confidence_result"
	} else if i.OrderByColumn == "related_group_amount" {
		i.OrderByColumn = "account_group_count"
	} else if i.OrderByColumn == "updateTime" {
		i.OrderByColumn = "update_time"
	}
	return i.OrderByColumn + " " + i.IsAsc
}

func (i PageInfo) GetOffset() int {
	return i.PageSize * (i.Page - 1)
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}
