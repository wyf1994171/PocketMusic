package model

type Comment struct {
	UID        uint   `json: "uid"`
	MID        uint   `json: "mid"`
	Content    string `json："content"`
	Status     uint   `json: "status"`
	Create_at  int64  `json: "create_time"`
	Update_at  int64  `json: "update_time"`
	RecordMeta
}
