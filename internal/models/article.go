package models

import "time"

type Article struct {
	ID       int64
	Title    string
	Content  string
	CreateAt time.Time
	UpdateAt time.Time
}

func (m *Article) ToArticleResponse() ArticleResponse {
	return ArticleResponse{
		ID:       m.ID,
		Title:    m.Title,
		Content:  m.Content,
		CreateAt: m.CreateAt,
		UpdateAt: m.UpdateAt,
	}
}
