package main

import (
    /*
    "os"

    //"github.com/OkanoShogo0903/osakana/controller"
    "github.com/OkanoShogo0903/osakana/routers"

    */
    "github.com/gin-contrib/cors"
    "net/http"
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

    router.Run(":8080")

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

    port := ""
    if len(os.Args) < 2 {
        port = os.Getenv("PORT")
    } else {
        port = os.Args[1]
    }
    router.Run(":" + port)
    */
}
