package main

import (
	"errors"
	"github.com/matisiekpl/electrocardiogram-server/internal/proto"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/matisiekpl/electrocardiogram-server/internal/client"
	"github.com/matisiekpl/electrocardiogram-server/internal/controller"
	"github.com/matisiekpl/electrocardiogram-server/internal/dto"
	"github.com/matisiekpl/electrocardiogram-server/internal/repository"
	"github.com/matisiekpl/electrocardiogram-server/internal/service"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Info("Error loading .env file")
	}

	config := dto.Config{DSN: os.Getenv("DSN"), SigningSecret: os.Getenv("JWT_SECRET")}

	db, err := gorm.Open(sqlite.Open(config.DSN), &gorm.Config{})
	if err != nil {
		logrus.Panic(err)
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.CORS())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				var appError dto.AppError
				switch {
				case errors.As(err, &appError):
					return echo.NewHTTPError(400, err.Error())
				}
			}
			return err
		}
	})

	clients := client.NewClients(config)
	repositories := repository.NewRepositories(db)
	services := service.NewServices(repositories, config, clients)
	controllers := controller.NewControllers(services)
	controllers.Route(e)

	go services.Record().Connect()

	protos := proto.NewProtos(services)
	go func() {
		grpcPort := os.Getenv("GRPC_PORT")
		if grpcPort == "" {
			grpcPort = "3001"
		}
		logrus.Info("Starting GRPC (TLS) server on port " + grpcPort)
		listener, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			logrus.Panic(err)
		}
		logrus.Fatal(protos.Serve(listener))
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	logrus.Info("Starting HTTP server on port " + port)
	logrus.Fatal(e.Start(":" + port))
}
