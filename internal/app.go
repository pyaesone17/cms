package internal

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Config *viper.Viper
	Log    *logrus.Logger
	Db     *mongo.Client
}
