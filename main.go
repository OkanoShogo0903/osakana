package main

import (
    "os"

    "github.com/OkanoShogo0903/osakana/controller"

    "github.com/gin-contrib/cors"

    "github.com/gin-gonic/gin"
)

func main() {
    aquarium_model := model.New()

    router := gin.Default()
    router.Use(cors.Default())

    router.LoadHTMLGlob("views/*.html")
    router.Static("/assets", "./assets")

    user := router.Group("/user")
    {
        user.POST("/signup", routers.UserSignUp)
        user.POST("/login", routers.UserLogIn)
    }

    router.GET("/", routers.Home)
    router.GET("/aquarium", controller.aquarium(aquarium_model))
    router.GET("/login", routers.LogIn)
    router.GET("/signup", routers.SignUp)
    router.NoRoute(routers.NoRoute)

    port := ""
    if len(os.Args) < 2 {
        port = os.Getenv("PORT")
    } else {
        port = os.Args[1]
    }
    router.Run(":" + port)
}
