package model

import "time"

type RecordMeta struct {
	ID        uint      `json:"ID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
