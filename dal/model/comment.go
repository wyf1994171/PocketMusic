package model

type Comment struct {
	UID        string   `json: "uid"`
	MID        string   `json: "mid"`
	Content    string `json："content"`
	Status     uint   `json: "status"`
	Create_at  []uint8  `json: "create_time"`
	Update_at  []uint8  `json: "update_time"`
	RecordMeta
}
