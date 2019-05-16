package model

import "time"

type RecordMeta struct {
	ID        uint      `json:"ID"`
	CreatedAt time.Time `json:"created_at"`
	Creator   string    `json:"Creator"`
	UpdatedAt time.Time `json:"updated_at"`
	Updater   string    `json:"Updater"`
}
