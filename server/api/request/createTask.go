package request

type CreateTask struct {
	Title      string `json:"title"`
	Date       string `json:"date"`
	CategoryID int    `json:"category_id"`
	UserID     int    `json:"user_id"`
}
