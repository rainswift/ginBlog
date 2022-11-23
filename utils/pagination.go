package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ginBlog/models"
	"io/ioutil"
	"strconv"
)

func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {

	limit := 20
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}
	}
	return models.Pagination{
		Limit:  limit,
		Page:   page,
		Sort:   sort,
	}

}

func RequestId(c *gin.Context) int {

	id := 1
	//query := c.Request.URL.Query()
	bodyByts, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Printf("删除数据22222： % + v\n", string(bodyByts))
	//fmt.Printf("删除id22222： % + v\n", query)
	for key, value := range bodyByts {
		fmt.Printf("删除数据22222： % + v\n  % + v\n", key,value)
		//queryValue := value[len(value)-1]

		//switch key {
		//case "id":
		//
		//	id, _  = strconv.Atoi(queryValue)
		//
		//	break
		//}
	}
	return id

}