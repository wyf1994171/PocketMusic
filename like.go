package main

import (
	"PocketMusic/dal"
	"github.com/gin-gonic/gin"
)

type GetLikeNumReq struct {
	Uid string `json:"uid" form:"uid"`
}

type GetLikeListReq struct {
	 Uid string `json:"uid" form:"uid"`
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

func HandleGetLikeList(c *gin.Context)  {
	var req GetLikeListReq
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	mids,err := dal.GetLikeList(req.Uid)
	if err != nil {
		c.Error(err)
		return
	}
	res := make([]map[string]interface{},0)
	for key := range mids {
		song,err := dal.GetMusicInfoById(mids[key])
		if err != nil {
			c.Error(err)
			return
		}
		res = append(res,song)
	}
	writeResponse(c,0,"",res)

}