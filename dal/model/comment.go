package model
type Comment struct {
	RecordMeta
	Content string `json:"content"`
	Uid    uint    `json:"uid"`
	Mid    uint    `json:"mid"`
	Status uint    `json:"status"`
}