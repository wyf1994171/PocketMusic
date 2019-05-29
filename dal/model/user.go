package model

type User struct {
	RecordMeta
	Uid        string `json:"uid"`
	Nickname   string `json:"nickname"`
	AvatarPath string `json:"avatar_path"`
}
