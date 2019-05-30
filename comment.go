package main

import (
	"PocketMusic/dal"
	"PocketMusic/dal/model"
	_ "github.com/Go-SQL-Driver/mysql"
	"github.com/gin-gonic/gin"
)

func HandleGetAllComment(c *gin.Context) {
	var req model.Comment
	if err := c.Bind(&req); err != nil {
		writeResponse(c,-1,err.Error(),nil)
		return
	}
	res, err := dal.GetAllComment(req.MID)
	for key := range res {
		name, err := dal.GetUserInfoByUid(res[key]["uid"].(string))
		if err != nil {
			writeResponse(c,-1,err.Error(),nil)
			return
		}
		res[key]["name"] = name
	}
	if err != nil {
		writeResponse(c,-1,err.Error(),nil)
		return
	}
	writeResponse(c,0,"",res)
}

func HandleCreateComment(c *gin.Context){
	var req model.Comment
	if err := c.Bind(&req); err != nil {
		writeResponse(c,-1,err.Error(),nil)
		return
	}
	err := dal.CreateComment(req.UID,req.MID,req.Content)
	if err != nil {
		writeResponse(c,-1,err.Error(),nil)
		return
	}
	writeResponse(c,0,"success", err)
}
