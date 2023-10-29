package db

import "time"

type BaseModel struct {
	Id         int64     `json:"id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
