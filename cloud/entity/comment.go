package entity

const TableComment = "comment"

//人员数据库实体
type Comments struct {
	ID        string `json:"id"`
	ComTxt    string `json:"com_txt"`    //词频前五
	ComInt    int    `json:"com_int"`    //评论数
	ComMale   int    `json:"com_male"`   //女性评论
	ComFemale int    `json:"com_female"` //男性评论
	ComArea   string `json:"com_area"`   //城市前五
}

type Comment struct {
	CreatedAt string  `json:"created_at"`
	Text      string  `json:"text"`
	User      ComUser `json:"user"`
	Gender    string  `json:"gender"`
}

type ComUser struct {
	Gender   string `json:"gender"`
	Location string `json:"location"`
}

func (Comment) TableName() string {
	return TableComment
}
