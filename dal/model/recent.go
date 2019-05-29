package model

type Recent struct {
	RecordMeta
	Uid    string `json:"uid"`
	Mid    uint   `json:"mid"`
	Status int    `json:"status"`
}
