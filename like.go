package main

import (
	"PocketMusic/dal"
	"github.com/gin-gonic/gin"
)

type GetLikeNumReq struct {
	Uid uint `json:"uid" form:"uid"`
}

func HandleGetLikeNum(c *gin.Context) {
	var req GetLikeNumReq
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	num, err := dal.GetLikeNum(req.Uid)
	if err != nil {
		c.Error(err)
		return
	}
	writeResponse(c,0,"",num)
}