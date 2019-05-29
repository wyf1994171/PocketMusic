package model

type Music struct {
	RecordMeta
	Mid    uint   `json:"mid"`
	Mname  string `json:"mname"`
	Singer string `json:"singer"`
	Lrc    string `json:"lrc"`
	Url    string `json:"url"`
}

func (Music)TableName() string  {
	return "songs"
}