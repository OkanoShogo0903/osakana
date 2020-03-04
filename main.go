package main

import (
    "os"
    "net/http"

    //"github.com/OkanoShogo0903/osakana/controller"
    //"github.com/OkanoShogo0903/osakana/routers"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    //router.LoadHTMLGlob("./vue-app/dist/*.html")
    //router.Static("/vue-app/dist", "./assets")
    //router.Static("/assets", "./assets")

    router.Use(cors.Default())

    router.StaticFS("/", http.Dir("./vue-app/dist"))

    /*
    router.GET("/", func(ctx *gin.Context){
        ctx.HTML(http.StatusOK, "index.html", gin.H{})
    })
    */

    /*
    //aquarium_model := model.New()


    user := router.Group("/user")
    {
        //user.POST("/signup", routers.UserSignUp)
        user.POST("/login", routers.UserLogIn)
    }

    router.GET("/", routers.Home)
    //router.GET("/aquarium", controller.aquarium(aquarium_model))
    router.GET("/login", routers.LogIn)
    router.GET("/signup", routers.SignUp)
    router.NoRoute(routers.NoRoute)
    */

    port, is_port := os.LookupEnv("PORT")
    if is_port == false {
        port = "8080"
    }

    router.Run(":" + port)
}
