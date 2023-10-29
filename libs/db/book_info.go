package db

import (
	"time"
)

type BookInfo struct {
	ID         int64     `json:"id,omitempty"          xorm:"id"`
	Author     string    `json:"author,omitempty"      xorm:"author"`
	BookName   string    `json:"book_name,omitempty"   xorm:"book_name"`
	Title      string    `json:"title,omitempty"       xorm:"title"`
	URL        string    `json:"url,omitempty"         xorm:"url"`
	Brief      string    `json:"brief,omitempty"       xorm:"brief"`
	KindID     int64     `json:"kind_id,omitempty"     xorm:"kind_id"`
	Bad        int64     `json:"bad,omitempty"         xorm:"bad"`
	NotBad     int64     `json:"not_bad,omitempty"     xorm:"not_bad"`
	Normal     int64     `json:"normal,omitempty"      xorm:"normal"`
	Good       int64     `json:"good,omitempty"        xorm:"good"`
	VeryGood   int64     `json:"very_good,omitempty"   xorm:"very_good"`
	CreateTime time.Time `json:"create_time,omitempty" xorm:"create_time"`
	UpdateTime time.Time `json:"update_time,omitempty" xorm:"update_time"`
}

func GetBookTitleByKind(kindID int64) (titles []string, err error) {
	var books []BookInfo
	err = Engine.Table("book_info").Where("kind_id = ?", kindID).Find(&books)
	for _, book := range books {
		titles = append(titles, book.Title)
	}
	return titles, err
}
