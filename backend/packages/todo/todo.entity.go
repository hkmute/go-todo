package todo

import "time"

type TodoEntity struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	User_id     int       `json:"userId"`
	Created_at  time.Time `json:"createdAt"`
	Updated_at  time.Time `json:"updatedAt"`
}
