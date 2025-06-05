package models

import "time"

type ArticleResponse struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
