package model

type Subscription struct {
	ID          int64  `json:"id"`
	ServiceName string `json:"service_name"`
	Price       int64  `json:"price"`
	UserId      string `json:"user_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}
