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
		return
	}
	res, err := dal.GetAllComment(req.Mid)
	for key := range res {
		name, err := dal.GetUserInfoByUid(res[key]["uid"].(string))
		if err != nil {
			writeResponse(ctx,-1,err.Error(),nil)
			return
		}
		res[key]["name"] = name
	}
	if err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		return
	}
	writeResponse(ctx,0,"",res)
}

func HandleCreateComment(ctx *gin.Context){
	var req CommentForm
	if err := ctx.Bind(&req); err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		return
	}
	ID, err := dal.CreateComment(req.Uid,req.Mid,req.Content)
	if err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		return
	}
	writeResponse(ctx,0,"success", ID)
}
