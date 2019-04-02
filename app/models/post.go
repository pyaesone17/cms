package models

import "encoding/json"

type Post struct {
	ID      string `json:"_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostWithCategory struct {
	*Post
	Categoy *Category `json:"category"`
}

func (p Post) Json() ([]byte, error) {
	return json.Marshal(p)
}

func (p PostWithCategory) Json() ([]byte, error) {
	return json.Marshal(p)
}

type PostFractal interface {
	Json() ([]byte, error)
}
