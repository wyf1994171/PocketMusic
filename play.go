package main

import (
	"github.com/gin-gonic/gin"
	"PocketMusic/dal"
	"PocketMusic/dal/model"
	"io/ioutil"
)
type GetLikeStatus struct {
	Mid uint `json:"mid" form:"mid"`
}

type Music struct {
	Mid uint `json:"mid"`
	Mname string `json:"mname"`
	Singer string `json:"singer"`
	Lrc string `json:"lrc"`
	Source []byte `json:"source"`
	LikeStatus bool `json:"like_status"`
}

func HandleGetLikeStatus(c *gin.Context) {
	var req GetLikeStatus
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	//println("mid:",req.Mid)
	status, err := dal.GetLikeStuatus(req.Mid)
	if err != nil {
		c.Error(err)
		return
	}
	writeResponse(c,0,"",status)
}

func HandleAddLike(c *gin.Context)  {
	var req GetLikeStatus
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	//println("mid:",req.Mid)
	status, err := dal.AddLike(req.Mid)
	if err != nil {
		c.Error(err)
		return
	}
	writeResponse(c,0,"",status)
}

func HandleDeleteLike(c *gin.Context){
	var req GetLikeStatus
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	//println("mid:",req.Mid)
	status, err := dal.DeleteLike(req.Mid)
	if err != nil {
		c.Error(err)
		return
	}
	writeResponse(c,0,"",status)
}
func HandleGetMusic(c *gin.Context)  {
	var req GetLikeStatus
	var musicinfo model.MusicInfo
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	//println("mid:",req.Mid)
	err := dal.GetMusic(req.Mid,&musicinfo)
	var music Music
	music.LikeStatus,err=dal.GetLikeStuatus(req.Mid)
	music.Mid=musicinfo.Mid
	music.Mname=musicinfo.Mname
	music.Singer=musicinfo.Singer
	music.Lrc=musicinfo.Lrc
	music.Source,err=ioutil.ReadFile(musicinfo.Source)
	if err != nil {
		c.Error(err)
		return
	}
	writeResponse(c,0,"",music)
}