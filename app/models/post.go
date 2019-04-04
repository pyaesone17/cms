package models

import (
	"time"
)

type Post struct {
	ID        string    `json:"_id" jsonapi:"primary,id"`
	Title     string    `json:"title" jsonapi:"attr,title" validate:"required"`
	Content   string    `json:"content" jsonapi:"attr,content" validate:"required"`
	Categoy   *Category `json:"category" jsonapi:"attr,category"`
	CreatedAt time.Time `json:"created_at" jsonapi:"attr,created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" jsonapi:"attr,updated_at" bson:"updated_at"`
}

func (p *Post) GetSaveModel() Post {
	updatedAt := time.Now()

	post := Post{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: updatedAt,
		UpdatedAt: updatedAt,
	}

	return post
}

func (p *Post) GetUpdateModel() Post {
	updatedAt := time.Now()

	post := Post{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: updatedAt,
	}

	return post
}
