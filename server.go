package server

import (
    "log"
    "net/http"

    "github.com/OkanoShogo0903/osakana/controller"
    db2 "github.com/OkanoShogo0903/osakana/db"

    "github.com/jmoiron/sqlx"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

type Server struct {
    db         *sqlx.DB
    router     *gin.Engine
}

func NewServer() *Server {
    return &Server{}
}

func (s *Server) Init(datasource string) {
    db := db2.NewDB(datasource)
    dbcon, err := db.Open()
    if err != nil {
        log.Fatalf("failed db init. %s", err)
    }
    s.db = dbcon
    s.router = s.Route()
}

func (s *Server) Run(port string) {
    log.Printf("Listening on port %s", port)
    s.router.Run(":" + port)
}

func (s *Server) Route() *gin.Engine {
    router := gin.Default()

    router.Use(cors.Default())

    router.StaticFS("/", http.Dir("./vue-app/dist"))

    aquarium_controller := controller.NewAquarium(s.db)

    user := router.Group("/user")
    {
        /*
        user.POST("/signup", func(c *gin.Context){
            c.HTML(http.StatusOK, "index.html", gin.H{})
        })
        */
        user.POST("/aquarium", aquarium_controller.GetUserData)
        user.PUT("/update", aquarium_controller.UpdateUserData)
        user.POST("/signup", aquarium_controller.Signup)
        //user.POST("/login", aquarium_controller.Login)

    }

    //router.GET("/", routers.Home)
    //router.GET("/aquarium", controller.aquarium(aquarium_model))
    //router.GET("/login", routers.LogIn)
    //router.GET("/signup", routers.SignUp)
    //router.NoRoute(routers.NoRoute)
    return router
}
