package model

import "time"

type Subscription struct {
	ID          int64
	ServiceName string
	Price       int64
	UserId      string
	StartDate   time.Time
	EndDate     *time.Time
}

type ExternalData struct {
	ID          int64
	ServiceName string `json:"service_name"`
	Price       int64  `json:"price"`
	UserId      string `json:"user_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}
