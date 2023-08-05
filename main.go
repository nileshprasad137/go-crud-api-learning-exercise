package main

import (
	"go-crud-gin/config"
	"go-crud-gin/controller"
	"go-crud-gin/helper"
	"go-crud-gin/model"
	"go-crud-gin/repository"
	"go-crud-gin/routers"
	"go-crud-gin/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Started server!")

	// DATABASE
	db := config.DatabaseConnection()
	validate := validator.New()
	db.Table("tags").AutoMigrate(&model.Tags{})

	// RESPOSITORY
	tagsRepository := repository.NewTagsRepositoryImpl(db)

	// service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	// controller
	tagsController := controller.NewTagsController(tagsService)

	// router
	routes := routers.NewRouter(tagsController)

	server := http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
