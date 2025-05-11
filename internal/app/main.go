package app

import (
	"github.com/sirupsen/logrus"
	"mindlab/internal/api"
	"mindlab/internal/api/controller"
	"mindlab/internal/config"
	"mindlab/internal/db/mongo"
	"mindlab/internal/repo"
)

func Run() {
	env, err := config.Init()

	// logs
	logLevel, err := logrus.ParseLevel(env.LogLevel)
	if err != nil {
		logrus.Warnf("Invalid log level: %s", env.LogLevel)
		logLevel = logrus.WarnLevel
	}
	logrus.SetLevel(logLevel)

	// MongoDb as main storage
	storage, err := mongo.NewClient(env.Mongo)
	if err != nil {
		logrus.Fatalf("Mongo init client error: %v", err)
	}
	db := storage.Database(env.Mongo.DBName)

	// crud repos
	postRepo := repo.NewPostRepository(db)
	tokenRepo := repo.NewTokenRepository(db)

	handler := &api.Handler{
		Post:  controller.NewPostController(postRepo),
		Login: controller.NewLoginController(tokenRepo),
	}

	// routes
	r := api.SetupRouter(handler)

	err = r.Run(":" + env.APIServerPort)
	if err != nil {
		panic(err)
	}
}
