package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"elephant-api/controllers"
	_ "elephant-api/docs"
	"time"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			user.GET("/login", controllers.GetUserLogin)
			user.POST("/login", controllers.PostUserLogin)
			user.GET("/register", controllers.GetUserRegister)
			user.POST("/register", controllers.PostUserRegister)
			user.GET("/list", controllers.GetUserList)
			//profile := user.Group("/profile").Use(middlewares.JWT())
		}
		movie := v1.Group("/movie")
		{
			problem.GET("/list", controllers.GetMovieList)
			problem.GET("/detail", controllers.GetMovieDetail)
			problem.POST("/create", controllers.PostCreateMovie)
			problem.POST("/update", controllers.PostUpdateMovie)
			problem.POST("/delete", controllers.PostDeleteMovie)
		}
		cinema := v1.Group("/cinema")
		{
			problem.GET("/list", controllers.GetCinemaList)
			problem.GET("/detail", controllers.GetCinemaDetail)
			problem.POST("/create", controllers.PostCreateCinema)
			problem.POST("/update", controllers.PostUpdateCinema)
			problem.POST("/delete", controllers.PostDeleteCinema)
		}
	}
	return router
}
