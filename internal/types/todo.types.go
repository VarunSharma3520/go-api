package types

type Todo struct {
	UserId   string `json:"userId"`
	Title string `json:"title"`
	Description     string `json:"description"`
	Reminder string `json:"reminder"`
}