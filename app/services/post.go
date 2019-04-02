package services

import (
	"github.com/pyaesone17/blog/app/datastore"
	"github.com/pyaesone17/blog/app/models"
	"github.com/pyaesone17/blog/internal"
	"github.com/sirupsen/logrus"
)

type PostService struct {
	postdatastore datastore.PostDB
	app           *internal.App
}

func NewPostService(app *internal.App) *PostService {
	return &PostService{
		postdatastore: datastore.NewPostDataStore(app.Db),
		app:           app,
	}
}

func (postservice *PostService) Create(post *models.Post, category *models.Category) {
	postservice.postdatastore.CreatePost(post)
	postservice.app.Log.WithFields(logrus.Fields{
		"post": post,
	}).Info("A post has been created")

	if category != nil {
		postservice.postdatastore.AddCategory(post, category)
	}
}

func (postservice *PostService) Get() ([]*models.Post, error) {
	return postservice.postdatastore.Get()
}

func (postservice *PostService) FindPost(id string) (*models.Post, error) {
	return postservice.postdatastore.Find(id)
}
