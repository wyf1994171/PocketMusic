package main

import (
	"PocketMusic/dal"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

type GetListsReq struct {
	Uid string `json:"uid" form:"uid"`
}

type AddListReq struct {
	Uid string`json:"uid" form:"uid"`
	Name string `json:"name" form:"name"`
}

type AddListSongReq struct {
	Lid uint `json:"lid" form:"lid"`
	Mid []uint `json:"mid" form:"uid"`
}

type DeleteListSongReq struct {
	Lid uint `json:"lid" form:"lid"`
	Mid []uint `json:"mid" form:"uid"`
}
type DeleteListReq struct {
	Uid string `json:"uid" form:"uid"`
	Lid []uint `json:"lid" form:"lid"`
}
type GetListSongsReq struct {
	Lid uint  `json:"lid" form:"lid"`
}
func HandleDeleteListSong (ctx *gin.Context) {
	var req DeleteListSongReq
	if err := ctx.Bind(&req); err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		return
	}
	for key := range req.Mid {
		err := dal.DeleteListSong(req.Lid,req.Mid[key])
		if err != nil {
			writeResponse(ctx,-1,err.Error(),nil)
			return
		}
	}
	err:=dal.ChangeNum(req.Lid)
	if err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		return
	}
	writeResponse(ctx,0,"success",nil)
}
func HandleListAddSong (ctx *gin.Context) {
	var req AddListSongReq
	if err := ctx.Bind(&req); err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		return
	}
	for key := range req.Mid {
		err := dal.AddListSong(req.Lid,req.Mid[key])
		if err != nil {
			writeResponse(ctx,-1,err.Error(),nil)
			return
		}
	}
	err:=dal.ChangeNum(req.Lid)
	if err != nil {
		writeResponse(ctx,-1,err.Error(),nil)
		return
	}
	writeResponse(ctx,0,"success",nil)
}
func HandleAddList (ctx *gin.Context) {
	var req AddListReq
	if err := ctx.Bind(&req); err != nil {
		ctx.Error(err)
		return
	}
	Id,err:=dal.AddList(req.Uid,req.Name)
	file, _, err := ctx.Request.FormFile("file")

	if err == nil {
		fileName := fmt.Sprintf("%v.jpg",Id)
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
		err = dal.AddListCoverPath(fileName,Id)
		if err != nil {
			writeResponse(ctx,-1,err.Error(),nil)
			return
		}
	}
	writeResponse(ctx,0,"",Id)
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

func HandleDeleteList(ctx *gin.Context)  {
	var req DeleteListReq
	if err := ctx.Bind(&req); err != nil {
		ctx.Error(err)
		return
	}
	for _,x:=range req.Lid{
		//println("uid: ",req.Uid)
		//println("lid:",x)
		err:=dal.DeleteList(req.Uid,x)
		if err!=nil{
			return
		}
	}
	writeResponse(ctx,0,"",true)
}
func HandleGetListSongs(ctx *gin.Context) {
	var req GetListSongsReq
	if err := ctx.Bind(&req); err != nil {
		ctx.Error(err)
		return
	}
	Ids,err := dal.GetListSongIds(req.Lid)
	if err != nil {
		ctx.Error(err)
		return
	}
	songs := make([]map[string]interface{},0)
	for key := range Ids {
		song,err := dal.GetMusicInfoById(Ids[key])
		if err != nil {
			ctx.Error(err)
			return
		}
		songs = append(songs,song)
	}
	writeResponse(ctx,0,"",songs)
}