package main

import (
	"PocketMusic/dal"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"fmt"
	"os"
)

type AddUserReq struct {
	Uid string `json:"uid" form:"uid"`
	Nickname string `json:"nickname" form:"nickname"`
}

func HandleAddUser (ctx *gin.Context) {
	var req AddUserReq
	if err := ctx.Bind(&req); err != nil {
		ctx.Error(err)
		return
	}
	file, _, err := ctx.Request.FormFile("file")

	if err != nil {
		writeResponse(ctx,-1,"文件解析错误",nil)
		return
	}
	fileName := fmt.Sprintf("/user/%v.jpg",req.Uid)
	buf, err := ioutil.ReadAll(file)
	file2, err := os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY,0755)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	defer file.Close()
	defer file2.Close()
	_, err = file2.WriteAt(buf,0)
	if err != nil {
		writeResponse(ctx,-1,"file read fail",nil)
		return
	}
	isUser, err := dal.GetIfUser(req.Uid)
	if err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		return
	}
	if !isUser {
		err = dal.AddUser(req.Uid, req.Nickname, fileName)
		if err != nil {
			writeResponse(ctx, -1, err.Error(), nil)
			return
		}
	}
	writeResponse(ctx,0,"",nil)
}