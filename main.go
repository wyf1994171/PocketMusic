package main

import (
	"PocketMusic/dal"
	"fmt"
	_ "github.com/Go-SQL-Driver"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := dal.InitDB("admin:password@tcp(119.29.111.64)/testdb"); err != nil {
		fmt.Println("Db error:%v",err)
		return
	}

	r := gin.Default()

	r.GET("/like_num",HandleGetLikeNum)
	fmt.Printf("Ready!")
	r.Run()
}
