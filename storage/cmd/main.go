package main

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	handler2 "storage/internal/handler"
	"storage/internal/repository"
	"storage/internal/server"
	service2 "storage/internal/service"
	"storage/pkg/minio"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing confis: %s", err.Error())
		return
	}

	db, err := repository.NewPostgres(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DbName:   viper.GetString("db.dbname"),
		SSlmode:  viper.GetString("db.sslmode"),
	})

	defer db.Close()

	if err != nil {
		logrus.Fatalf("failed to initializ db : %s ", err.Error())
		return
	}

	client, err := minio.NewMinioClient(viper.GetString("endpoint"), viper.GetString("accessKeyId"), viper.GetString("secretaccessKeyId"))
	if err != nil {
		logrus.Fatalf("failed to initializing client. Err: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service2.NewService(repos, client)
	handlers := handler2.NewHandler(service)

	srv := new(server.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
			return
		}
	}()

	logrus.Println("server starting...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("app shutting down ...")

	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Error(err)
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
