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

func (postservice *PostService) Create(post *models.Post, category *models.Category) error {
	err := postservice.postdatastore.CreatePost(post)

	if err != nil {
		postservice.app.Log.WithFields(logrus.Fields{
			"post": post,
		}).Info("A post has been created")
	}

	if category != nil {
		postservice.postdatastore.AddCategory(post, category)
	}

	return err
}

func (postservice *PostService) Update(post *models.Post, category *models.Category) error {
	err := postservice.postdatastore.UpdatePost(post)

	if err != nil {
		postservice.app.Log.WithFields(logrus.Fields{
			"post": post,
		}).Info("A post has been updated")
	}

	return err
}

func (postservice *PostService) Get() ([]*models.Post, error) {
	return postservice.postdatastore.Get()
}

func (postservice *PostService) FindPost(id string) (*models.Post, error) {
	return postservice.postdatastore.Find(id)
}
