package main

import (
	"./controllers"
	"./middleware"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"log"
)

func main() {
	router := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("byandev %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	router.Use(cors.AllowAll())

	public := router.Group("/api/v1")
	{
		public.POST("/login_admin", controllers.LoginAdmin)
		public.POST("/login_petugas", controllers.LoginPetugas)
		public.POST("/login_users", controllers.LoginUsers)
	}
	v1 := router.Group("/api/v1")
	v1.Use(middleware.Auth)
	{
		v1.POST("/create_akun", controllers.CreateUsers)

		v1.GET("/akun/list", controllers.GetUsers)
		v1.GET("/akun/detail", controllers.GetUserDetail)
		v1.PUT("/akun/update", controllers.UpdateUser)

		v1.GET("/rekom/list", controllers.GetRekom)
		v1.GET("/rekom/detail", controllers.GetRekomDetail)
		v1.POST("/rekom/create", controllers.CreateRekomendasi)
		v1.PUT("/rekom/update", controllers.UpdateRekom)

		v1.GET("/privileges/role", controllers.GetRole)
		v1.GET("/privileges/type", controllers.GetType)
	}
	router.Run(":8000")
}
