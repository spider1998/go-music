package entity

import (
	"cloud/util"
)

type User struct {
	ID         string    `json:"id" xorm:"pk"`
	UserID     int       `json:"user_id"`
	AlbumSize  int       `json:"album_size"`
	MusicSize  int       `json:"music_size"`
	PicUrl     string    `json:"pic_url"`
	Name       string    `json:"name"`
	AccountID  int       `json:"account_id"`
	Cat        string    `json:"cat"`
	City       string    `json:"city"`
	UpdateTime util.Time `json:"updateTime" xorm:"updated"`
	CreateTime util.Time `json:"createTime" xorm:"created"`
}
