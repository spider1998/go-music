package entity

type SingerDesc struct {
	Introduction []struct {
		Ti  string `json:"ti"`
		Txt string `json:"txt"`
	} `json:"introduction"`
	BriefDesc string `json:"briefDesc"`
	Count     int    `json:"count"`
	Code      int    `json:"code"`
}
