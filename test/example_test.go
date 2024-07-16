// Package test -----------------------------
// @file      : example_test.go
// @author    : dingyq
// @time      : 2024/7/15 下午4:35
// -------------------------------------------
package test

import (
	"fmt"
	"github.com/dingyuqi/go-pagehelper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"testing"
)

// set your own port
var port = 8080
var DB *gorm.DB

func TestPageHelper(t *testing.T) {
	engine := gin.New()
	group := engine.Group("")
	testGroup := group.Group("test")
	// add pagehelper middleware
	testGroup.GET("/pagehelper", pagehelper.PaginationMiddleware, exampleFunc)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: engine,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Println("ListenAndServe error : ", err)
		return
	}
}

func exampleFunc(ctx *gin.Context) {
	var data []*UserData
	pageParam := pagehelper.GetPaginationParam(ctx)
	tx := DB.Model(&UserData{})
	query := tx.Select("*")
	pageParam.GetResult(query, &data)
	OkWithData(pageParam.ToResult(data), ctx)
}

// the database table that we want to fetch
type UserData struct {
	Id   int64  `json:"id,string" gorm:"column:ID;"`
	Name string `json:"name" gorm:"column:NAME;"`
}
