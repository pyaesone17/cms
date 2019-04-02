package models

type Post struct {
	ID      string    `json:"_id" jsonapi:"primary,id"`
	Title   string    `json:"title" jsonapi:"attr,title"`
	Content string    `json:"content" jsonapi:"attr,content"`
	Categoy *Category `json:"category" jsonapi:"attr,category"`
}
