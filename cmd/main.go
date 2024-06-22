package main

import (
	"context"
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/kondrushin/blog/internal/middleware"
	"github.com/kondrushin/blog/internal/repository"
	"github.com/kondrushin/blog/internal/router"
	"github.com/kondrushin/blog/internal/seeding"
	"github.com/kondrushin/blog/internal/usecase"

	"errors"
	"log/slog"
	"os"
)

func main() {
	dataFilePath := flag.String("seed", "", "Location of a data file to seed the database")
	flag.Parse()

	engine := gin.Default()
	middleware.Setup(engine)

	repository := repository.NewRepository()
	usecase := usecase.NewBlogUseCase(repository)
	router.RegisterHandlers(engine, usecase)

	slog.Info("Service started")

	if len(*dataFilePath) > 0 {
		seed(dataFilePath, repository)
	}

	engine.Run(":8080")
}

func seed(dataFilePath *string, repository *repository.Repository) {
	if _, err := os.Stat(*dataFilePath); err == nil {
		slog.Info("DB seeding started.", "source", *dataFilePath)
		err := seeding.Seed(context.Background(), *dataFilePath, repository)
		if err != nil {
			slog.Error("Error while seeding.", "error", err)
			return
		}

		slog.Info("DB seeding is completed.")

	} else if errors.Is(err, os.ErrNotExist) {
		slog.Error("Data file is not found.", "file", *dataFilePath)
	}
}
