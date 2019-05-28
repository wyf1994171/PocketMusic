package model

type ListSong struct {
	RecordMeta
	Lid uint `json:"lid"`
	Mid uint `json:"mid"`
	Status uint `json:"status"`
}

func (ListSong) TableName() string {
	return "list_songs"
}