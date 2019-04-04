package models

import (
	"time"

	"github.com/pyaesone17/blog/app/core"
)

type Post struct {
	*core.MainModel
	ID      string    `json:"_id" jsonapi:"primary,id"`
	Title   string    `json:"title" jsonapi:"attr,title" validate:"required"`
	Content string    `json:"content" jsonapi:"attr,content" validate:"required"`
	Categoy *Category `json:"category" jsonapi:"attr,category"`
}

func (p *Post) GetSaveModel() Post {
	updatedAt := time.Now()

	mainModel := &core.MainModel{
		CreatedAt: p.CreatedAt,
		UpdatedAt: updatedAt,
	}

	if time.Time.IsZero(p.CreatedAt) {
		mainModel.CreatedAt = updatedAt
	}

	post := Post{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		MainModel: mainModel,
	}

	return post
}
