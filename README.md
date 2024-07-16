# go-pagehelper

**Read this in other languages: [简体中文](./README.zh-CN.md)**

go-pagehelper is a tool that helps paginate HTTP interface data under the GIN framework. By using this tool as the middleware of the GIN framework, you can easily return data according to the specified number of pages and paging size in the HTTP interface.

## Install
Install go-pagehelper using `go get`
```shell
go get github.com/dingyuqi/go-pagehelper@latest
```
## Usage
### Service
> [!IMPORTANT]
> The following code is just an example. For executable code, see the `example_test.go` test code in the `/test` folder.

```go
package main

import (
	"fmt"
	"github.com/dingyuqi/go-pagehelper"
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

func exampleFunc(ctx *gin.Context) {
	var data []*UserData
	pageParam := pagehelper.GetPaginationParam(ctx)
	tx := DB.Model(&UserData{})
	query := tx.Select("*")
	pageParam.GetResult(query, &data)
	OkWithData(pageParam.ToResult(data), ctx)
}

// UserData the database table that we want to fetch
type UserData struct {
	Id   int64  `json:"id,string" gorm:"column:ID;"`
	Name string `json:"name" gorm:"column:NAME;"`
}
```
### Client
For HTTP requests, four parameters can be configured according to different interface requirements:

| Serial number | Parameter name | Data type | Function |
|----|---------------|--------|------------------- --------------------------------------------------|
| 1 | pageNum | int | The number of pages you want to request, for example, if you want to request the first page, Page=1. |
| 2 | pageSize | int | Number of data items per page. |
| 3 | orderByColumn | string | The field by which you want to sort the returned data. This field needs to be the same as the field name in the database. |
| 4 | isAsc | string | Determines whether the fields specified by orderByColumn are sorted in positive order or to be sorted sequentially. The fill-in content is the string after the `ORDER BY` keyword in the query sql |

For example:
```shell
curl --request GET \
  --url http://127.0.0.1/test/pagehelper \
  --form pageNum=3 \
  --form pageSize=20 \
  --form orderByColumn=CREATE_TIME \
  --form isAsc=desc
```
The meaning of the above request is to request to obtain 20 pieces of data for one page after sorting by the field `CREATE_TIME` in reverse order, including the third page of data.

> [!TIP]
> If you want to put the data with a null value (NULL) in `CREATE_TIME` at the end, you can configure it as `isAsc=desc nulls last`
## License
[MIT](https://choosealicense.com/licenses/mit/) License