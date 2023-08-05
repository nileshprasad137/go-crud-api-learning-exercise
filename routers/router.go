package routers

import (
	"go-crud-gin/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(tagsController *controller.TagsController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	tagsRouter := router.Group("/tags")
	tagsRouter.GET("", tagsController.FindAll)
	tagsRouter.GET("/:tagId", tagsController.FindById)
	tagsRouter.POST("/", tagsController.Create)
	tagsRouter.PATCH("/:tagId", tagsController.Update)
	tagsRouter.DELETE("/:tagId", tagsController.Delete)

	return router
}
