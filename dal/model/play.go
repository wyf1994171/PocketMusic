package model

type MusicInfo struct {
	Mid uint `json:"mid"`
	Mname string `json:"mname"`
	Singer string `json:"singer"`
	Lrc string `json:"lrc"`
	Source string `json:"source"`
	//LikeStatus bool `json:"like_status"`
} 
