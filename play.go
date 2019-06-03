package main

import (
	"PocketMusic/dal"
	"PocketMusic/dal/model"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	//"strings"
)

type GetLikeStatus struct {
	Uid string `json:"uid" form:"uid"`
	Mid uint   `json:"mid" form:"mid"`
}
type ChangeLikeStatus struct {
	Uid    string `json:"uid" form:"uid"`
	Mid    uint   `json:"mid" form:"mid"`
	Status uint   `json:"status" form:"status"`
}
type Music struct {
	Mid        uint   `json:"mid"`
	Name       string `json:"name"`
	Singer     string `json:"singer"`
	Lyrics     string `json:"lyrics"`
	Audio      string `json:"audio"`
	LikeStatus bool   `json:"like_status"`
}

func HandleGetLikeStatus(c *gin.Context) {
	var req GetLikeStatus
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	//println("mid:",req.Mid)
	status, err := dal.GetLikeStuatus(req.Uid, req.Mid)
	if err != nil {
		c.Error(err)
		return
	}
	writeResponse(c, 0, "", status)
}

func HandleAddLike(c *gin.Context) {
	var req GetLikeStatus
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	//println("mid:",req.Mid)
	err := dal.AddLike(req.Uid, req.Mid)
	if err != nil {
		c.Error(err)
		return
	}
	writeResponse(c, 0, "", 0)
}

func HandleDeleteLike(c *gin.Context) {
	var req GetLikeStatus
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	//println("mid:",req.Mid)
	err := dal.DeleteLike(req.Uid, req.Mid)
	if err != nil {
		c.Error(err)
		return
	}
	writeResponse(c, 0, "", 0)
}
func HandleGetMusic(c *gin.Context) {
	var req GetLikeStatus
	var musicinfo model.MusicInfo
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	//println("mid:",req.Mid)
	err := dal.GetMusic(req.Mid, &musicinfo)
	var music Music
	if err != nil {
		c.Error(err)
		return
	}
	err = dal.AddRecent(req.Uid, req.Mid)
	if err != nil {
		writeResponse(c, -1, err.Error(), nil)
		return
	}
	if musicinfo.Source != "" {
		music.LikeStatus, err = dal.GetLikeStuatus(req.Uid, req.Mid)
		music.Mid = musicinfo.Mid
		music.Name = musicinfo.Mname
		music.Singer = musicinfo.Singer
		source,_:=ioutil.ReadFile(musicinfo.Lrc)
		music.Lyrics=string(source)
		//println("lrc:",music.Lyrics)
		music.Audio = musicinfo.Source
		//println("source",source)
		writeResponse(c, 0, "", music)
	} else {
		writeResponse(c, 1, "未找到歌曲资源！", music)
	}
}

type SongsReq struct {
	Url string `json:"url" form:"url"`
}

func HandleGetSong(c *gin.Context) {
	var req SongsReq
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	source, _ := ioutil.ReadFile(req.Url)
	c.Data(200, "audio/mp3", source)
}

func HandleLikestatus(c *gin.Context) {
	var req ChangeLikeStatus
	if err := c.Bind(&req); err != nil {
		c.Error(err)
		return
	}
	//println("mid:",req.Mid)
	err := dal.ChangeLike(req.Uid, req.Mid, req.Status)
	if err != nil {
		c.Error(err)
		return
	}
	writeResponse(c, 0, "", 0)
}
