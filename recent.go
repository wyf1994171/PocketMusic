package main

import (
	"PocketMusic/dal"
	"github.com/gin-gonic/gin"
)

type GetRecentReq struct {
	Uid string `json:"uid" form:"uid"`
}

func HandleGetRecent(ctx *gin.Context) {
	var req GetRecentReq
	if err := ctx.Bind(&req); err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		return
	}
	mids, err := dal.GetRecent(req.Uid)
	if err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		return
	}
	res := make([]map[string]interface{},0)
	for key := range mids {
		m,err := dal.GetMusicInfoById(mids[key])
		if err != nil {
			writeResponse(ctx,-1,err.Error(),nil)
			return
		}
		res = append(res,m)
	}
	writeResponse(ctx,0,"",res)
}