package main

import (
	"PocketMusic/dal"
	"PocketMusic/dal/model"
	_ "github.com/Go-SQL-Driver/mysql"
	"github.com/gin-gonic/gin"
)

type Comment struct {
	Id	uint `json:"id" form:"id"`
	Uid uint `json:"uid" form:"uid"`
	Mid	uint `json:"mid" form:"mid"`
	Content string `json:"content" form:"content"`
}

func HandleGetAllComment(c *gin.Context) {
	var req Comment
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	num, err := dal.GetAllComment(req.Mid)
	if err != nil {
		c.Error(err)
		return
	}
	writeResponse(c,0,"",num)
}

func HandleCreateComment(c *gin.Context){
	var req model.CommentForm
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	err := dal.CreateComment(req.UID)
	if err != nil {
		c.Error(err)
		return
	}
}