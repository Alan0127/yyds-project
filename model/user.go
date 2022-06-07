package model

type UserInfo struct {
	UserName string `json:"userName"`
	Password string `json:"paasWord"`
	Country  string `json:"country"`
	Token    string `json:"token"`
}
