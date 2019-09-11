package entity

type Account struct {
	Level       int `json:"level"`
	ListenSongs int `json:"listenSongs"`
	Profile     struct {
		Gender   int `json:"gender"`
		Province int `json:"province"`
	} `json:"profile"`
	Code int `json:"code"`
}
