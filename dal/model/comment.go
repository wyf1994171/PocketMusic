package model

type Comment struct {
	UID			uint	`json: "uid"`
	MID			uint	`json: "mid"`
	Content 	string	`json："content"`
	Create_at	int64	`json: "create_time"`
	Update_at	int64	`json: "update_time"`
}

type CommentForm struct {
	UID		uint	`json:"uid"`
	Content	string  `json:"content"`
}