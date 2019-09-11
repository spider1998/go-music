package entity

import (
	"cloud/util"
)

type SingerTopList struct {
	ID          string    `json:"id"`
	UserID      int       `json:"user_id"`
	Name        string    `json:"name"`
	Alias       []string  `json:"alias"`        //别名
	TopicPerson int       `json:"topic_person"` //讨论人数
	Ranking     int       `json:"ranking"`      //排名
	LastRank    int       `json:"last_rank"`    //上次排名
	Score       int       `json:"score"`        //热度
	UpdateTime  util.Time `json:"updateTime" xorm:"updated"`
	CreateTime  util.Time `json:"createTime" xorm:"created"`
}

type GetSingerTops struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Alias       []string `json:"alias"`       //别名
	TopicPerson int      `json:"topicPerson"` //讨论人数
	Ranking     int      `json:"ranking"`     //排名
	LastRank    int      `json:"lastRank"`    //上次排名
	Score       int      `json:"score"`       //热度
}

type GetSingerTopResult struct {
	List struct {
		Artists    []GetSingerTops `json:"artists"`
		UpdateTime int             `json:"updateTime"`
		Type       int             `json:"type"`
	} `json:"list"`
	Code int `json:"code"`
}
