package models

type Project struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Twitter     string `json:"twitter"`
	Instagram   string `json:"instagram"`
	Linkedin    string `json:"linkedin"`
	Owner       string `json:"owner"`
	Image       string `json:"image"`
}
