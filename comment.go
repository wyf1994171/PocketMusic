package main

import (
	"PocketMusic/dal"
	"github.com/gin-gonic/gin"
)

type CommentForm struct {
	Uid     string `json:"uid" form:"uid"`
	Mid     string `json:"mid" form:"mid"`
	Content string `json:"content" form:"content"`
}

func HandleGetAllComment(c *gin.Context) {
	var req CommentForm
	if err := c.Bind(&req); err != nil {
		writeResponse(c, -1, err.Error(), nil)
		return
	}
	res, err := dal.GetAllComment(req.Mid)
	if err != nil {
		writeResponse(c, -1, err.Error(), nil)
		return
	}
	writeResponse(c, 0, "", res)
}

func HandleCreateComment(c *gin.Context) {
	var req CommentForm
	if err := c.Bind(&req); err != nil {
		writeResponse(c, -1, err.Error(), nil)
		return
	}
	err := dal.CreateComment(req.Uid, req.Mid, req.Content, dsn)
	if err != nil {
		writeResponse(c, -1, err.Error(), nil)
		return
	}
	writeResponse(c, 0, "success", err)
}
