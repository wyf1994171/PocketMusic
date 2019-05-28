package model

type List struct {
	RecordMeta
	ID        uint   `json:"id"`
	UID       uint   `json:"uid"`
	CoverPath string `json:"cover_path"`
	Num       uint   `json:"num"`
	Name      string `json:"name"`
	Status    uint   `json:"status"`
}
