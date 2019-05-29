package main

import (
	"PocketMusic/dal"
	"fmt"
	_ "github.com/Go-SQL-Driver/mysql"
	"github.com/gin-gonic/gin"
)

func main() {

	if err := dal.InitDB("admin:password@tcp(119.29.111.64)/testdb"); err != nil {
		fmt.Println("Db error:%v", err)
		return
	}

	r := gin.Default()


	r.GET("/search:Mname_Or_Singer", HandleSearch)
	r.GET("/like_num",HandleGetLikeNum)
	r.GET("/lists",HandleGetLists)
	r.POST("/lists",HandleAddList)
	r.POST("/lists/song",HandleListAddSong)
	r.DELETE("/lists/song",HandleDeleteListSong)
	r.GET("/play",HandleGetLikeStatus)
	r.POST("/play",HandleAddLike)
	r.DELETE("/play",HandleDeleteLike)
	r.GET("/play/music",HandleGetMusic)
	r.POST("/create_comment",HandleCreateComment)
	r.GET("/get_all_comment", HandleGetAllComment)
	r.GET("/songs",HandleGetSong)
	r.GET("/recent",HandleGetRecent)
	r.POST("user",HandleAddUser)
	fmt.Printf("Ready!")
	r.Run(":7007")
}
