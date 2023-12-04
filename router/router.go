package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuliussetyawan/task-5-pbi-btpns-yulius/controllers"
	"github.com/yuliussetyawan/task-5-pbi-btpns-yulius/middlewares"
)

func StartRoute() *gin.Engine {
	route := gin.Default()

	userRouter := route.Group("/users")
	{
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.PUT("/:userId", controllers.UserUpdate)
		userRouter.DELETE("/:userId", controllers.UserDelete)
	}

	photoRouter := route.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())

		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.ListPhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	return route
}
