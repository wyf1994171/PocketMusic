package main

import (
	"PocketMusic/dal"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type GetListsReq struct {
	Uid uint `json:"uid" form:"uid"`
}

func HandleGetLists (ctx *gin.Context) {
	var req GetListsReq
	if err := ctx.Bind(&req); err != nil {
		ctx.Error(err)
		return
	}
	result, err := dal.GetLists(req.Uid)
	if err != nil {
		ctx.Error(err)
		return
	}
	for key := range result {
		result[key]["cover_path"], _ = ioutil.ReadFile(result[key]["cover_path"].(string))
	}
	writeResponse(ctx,0,"", result)
}