# go-pagehelper

**其他语言版本: [English](./README.md)**

go-pagehelper 是一个帮助GIN框架下HTTP接口数据分页的工具. 通过使用该工具作为GIN框架的中间件(middleware),
在HTTP接口中可以轻松按照指定的页数以及分页大小对数据进行返回.

## 安装

使用`go get`安装 go-pagehelper

```shell
go get github.com/dingyuqi/go-pagehelper@latest
```

## 用法

### 服务端

> [!IMPORTANT]  
> 下面的代码仅是个示例, 可运行的代码见 `/test` 文件夹下的 `example_test.go` 测试代码.

```go
package main

import (
	"fmt"
	"github.com/dingyuqi/go-pagehelper/respond"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// set your own port
var port = 8080
var DB *gorm.DB

func main() {
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

// UserData the database table that we want to fetch
type UserData struct {
	Id   int64  `json:"id,string" gorm:"column:ID;"`
	Name string `json:"name" gorm:"column:NAME;"`
}

func exampleFunc(ctx *gin.Context) {
	var data []*UserData
	pageParam := pagehelper.GetPaginationParam(ctx)
	tx := DB.Model(&UserData{})
	query := tx.Select("*")
	pageParam.GetResult(query, &data)
	respond.OkWithData(pageParam.ToResult(data), ctx)
}

```

### 客户端

对于HTTP请求, 可以按照不同的接口需求配置四个参数:

| 序号 | 参数名           | 数据类型   | 作用                                                              |
|----|---------------|--------|-----------------------------------------------------------------|
| 1  | pageNum       | int    | 想请求的页数, 例如想请求第一页则Page=1.                                        |
| 2  | pageSize      | int    | 每页的数据条数.                                                        |
| 3  | orderByColumn | string | 希望按照哪个字段对返回的数据进行排序. 该字段需要与数据库中的字段名相同.                           |
| 4  | isAsc         | string | 决定orderByColumn指定的字段按照正序排序还是待续排序, 填写内容为查询sql中`ORDER BY`关键词后的字符串 |

例如:
```shell
curl --request GET \
  --url http://127.0.0.1/test/pagehelper \
  --form pageNum=3 \
  --form pageSize=20 \
  --form orderByColumn=CREATE_TIME \
  --form isAsc=desc
```
上面这个请求的含义是请求获取以字段`CREATE_TIME`倒序排序后, 20条数据为一页进行分页, 其中第三页的数据.

> [!TIP]
> 如果想把`CREATE_TIME`中值为空(NULL)的数据放在最后, 可配置为`isAsc=desc nulls last`

## 认证
[MIT](https://choosealicense.com/licenses/mit/) License