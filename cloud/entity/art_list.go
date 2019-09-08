package entity

const TableArtList = "art_list"

//人员数据库实体
type ArtList struct {
	ID string `json:"id" xorm:"pk"`
	UserID int `json:"user_id"`
	AlbumSize int `json:"album_size"`
	MusicSize int `json:"music_size"`
	PicUrl string `json:"pic_url"`
	Name string `json:"name"`
	AccountID int `json:"account_id"`
	Cat string `json:"cat"`
	UpdateTime string `json:"updateTime" xorm:"updated"`
	CreateTime string `json:"createTime" xorm:"created"`
}


type Artlists struct {
	Artists []ArtlistResult `json:"artists"`
	More bool `json:"more"`
	Code int `json:"code"`
}

type ArtlistResult struct {
	ID int `json:"id"`
	AlbumSize int `json:"albumSize"`
	MusicSize int `json:"musicSize"`
	PicUrl string `json:"picUrl"`
	Name string `json:"name"`
	AccountID int `json:"accountId"`
	Cat string `json:"cat"`

}

func (ArtList) TableName() string {
	return TableArtList
}

var ArtAZ  =map[int]string{
	1:"a",
	2:"b",
	3:"c",
	4:"d",
	5:"e",
	6:"f",
	7:"g",
	8:"h",
	9:"i",
	10:"j",
	11:"k",
	12:"l",
	13:"m",
	14:"n",
	15:"o",
	16:"p",
	17:"q",
	18:"r",
	19:"s",
	20:"t",
	21:"u",
	22:"v",
	23:"w",
	24:"x",
	25:"y",
	26:"z",
}
