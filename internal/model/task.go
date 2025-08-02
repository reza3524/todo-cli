package model

type Task struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Date       string `json:"date"`
	Completed  bool   `json:"completed"`
	CategoryID int    `json:"categoryId"`
	UserID     int    `json:"userId"`
}
