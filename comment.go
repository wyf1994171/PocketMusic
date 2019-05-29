package main

import (
	"PocketMusic/dal"
	_ "github.com/Go-SQL-Driver/mysql"
	"github.com/gin-gonic/gin"
)

type CommentForm struct {
	Uid uint `json:"uid" form:"uid"`
	Mid	uint `json:"mid" form:"mid"`
	Content string `json:"content" form:"content"`
}

func HandleGetAllComment(ctx *gin.Context) {
	var req CommentForm
	if err := ctx.Bind(&req); err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		ctx.Error(err)
		return
	}
	res, err := dal.GetAllComment(req.Mid)
	if err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		ctx.Error(err)
		return
	}
	writeResponse(ctx,0,"",res)
}

func HandleCreateComment(ctx *gin.Context){
	var req CommentForm
	if err := ctx.Bind(&req); err != nil {
		ctx.Error(err)
		writeResponse(ctx,-1,err.Error(),nil)
		return
	}
	ID, err := dal.CreateComment(req.Uid,req.Mid,req.Content)
	if err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		ctx.Error(err)
		return
	}
	writeResponse(ctx,0,"success", ID)
}
