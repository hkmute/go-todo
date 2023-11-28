package todo

import "time"

type todoEntity struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Created_at  time.Time `json:"createdAt"`
	Updated_at  time.Time `json:"updatedAt"`
}
