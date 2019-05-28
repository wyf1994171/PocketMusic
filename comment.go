package main

import (
	"PocketMusic/dal"
	"github.com/gin-gonic/gin"
)

type AddCommentReq struct {
	Content string  `json:"content" form:"content"`
	Uid     uint	`json:"uid" form:"uid"`
	Mid     uint    `json:"mid" form:"mid"`
}

type CommentReq struct {
	Mid   uint `json:"mid" form:"mid"`
}

type DeleteCommentReq struct {
	Id uint	`json:"id" form:"id"`
}
func HandleAddComment(ctx *gin.Context){
	var req AddCommentReq
	if err := ctx.Bind(&req); err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		return
	}
	ID, err := dal.AddComment(req.Uid,req.Mid,req.Content)
	if err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		return
	}
	writeResponse(ctx,0,"success", ID)
}

func HandleGetComment (ctx *gin.Context) {
	var req CommentReq
	if err := ctx.Bind(&req); err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		return
	}
	res, err := dal.GetComment(req.Mid)
	if err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		return
	}
	writeResponse(ctx,0,"success", res)
}

func HandleDeleteComment (ctx *gin.Context) {
	var req DeleteCommentReq
	if err := ctx.Bind(&req); err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		return
	}
	err := dal.DeleteComment(req.Id)
	if err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		return
	}
	writeResponse(ctx,0,"success", 0)
}