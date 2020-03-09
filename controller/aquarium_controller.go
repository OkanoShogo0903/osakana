package controller

import (
    "log"
    "net/http"

    "github.com/OkanoShogo0903/osakana/model"

    "github.com/jmoiron/sqlx"
    "github.com/gin-gonic/gin"
)

type Aquarium struct {
    db *sqlx.DB
}

func NewAquarium(db *sqlx.DB) *Aquarium {
    return &Aquarium{db: db}
}

func (a *Aquarium)GetUserData(c *gin.Context) {
    log.Printf("GetUserData")

    var user model.User
    user.IP = c.ClientIP()

    m := make([]model.AquariumDB, 0)
    err := a.db.Select(&m, `SELECT * FROM users WHERE ip = ?`, c.ClientIP())
    if err != nil {
        log.Printf(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, m)
}
func (a *Aquarium)UpdateUserData(c *gin.Context) {
    log.Printf("UpdateUserData")

    var user model.User
    user.IP = c.ClientIP()

    if err := c.ShouldBindJSON(&user.Aquarium); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    f := func(a *Aquarium, ctx *gin.Context, ip string, field string, val int64) {

        stmt, err := a.db.Prepare(`UPDATE users SET `+field+` = (?) WHERE ip = (?)`)
        if err != nil {
            log.Printf(err.Error())
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }

        defer stmt.Close()
        _, err = stmt.Exec(val, ip)
        if err != nil {
            log.Printf(err.Error())
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
    }

    f(a, c, user.IP, "background", user.Aquarium.BackgroundID)
    f(a, c, user.IP, "plant00", user.Aquarium.PlantIDs[0])
    f(a, c, user.IP, "plant01", user.Aquarium.PlantIDs[1])
    f(a, c, user.IP, "plant02", user.Aquarium.PlantIDs[2])
    f(a, c, user.IP, "plant03", user.Aquarium.PlantIDs[3])
    f(a, c, user.IP, "plant04", user.Aquarium.PlantIDs[4])
    f(a, c, user.IP, "plant05", user.Aquarium.PlantIDs[5])
    f(a, c, user.IP, "plant06", user.Aquarium.PlantIDs[6])
    f(a, c, user.IP, "plant07", user.Aquarium.PlantIDs[7])
    f(a, c, user.IP, "plant08", user.Aquarium.PlantIDs[8])
    f(a, c, user.IP, "plant09", user.Aquarium.PlantIDs[9])
    f(a, c, user.IP, "creature00", user.Aquarium.CreatureIDs[0])
    f(a, c, user.IP, "creature01", user.Aquarium.CreatureIDs[1])
    f(a, c, user.IP, "creature02", user.Aquarium.CreatureIDs[2])
    f(a, c, user.IP, "creature03", user.Aquarium.CreatureIDs[3])
    f(a, c, user.IP, "creature04", user.Aquarium.CreatureIDs[4])
    f(a, c, user.IP, "creature05", user.Aquarium.CreatureIDs[5])
    f(a, c, user.IP, "creature06", user.Aquarium.CreatureIDs[6])
    f(a, c, user.IP, "creature07", user.Aquarium.CreatureIDs[7])
    f(a, c, user.IP, "creature08", user.Aquarium.CreatureIDs[8])
    f(a, c, user.IP, "creature09", user.Aquarium.CreatureIDs[9])
    f(a, c, user.IP, "creature10", user.Aquarium.CreatureIDs[10])
    f(a, c, user.IP, "creature11", user.Aquarium.CreatureIDs[11])
    f(a, c, user.IP, "creature12", user.Aquarium.CreatureIDs[12])
    f(a, c, user.IP, "creature13", user.Aquarium.CreatureIDs[13])
    f(a, c, user.IP, "creature14", user.Aquarium.CreatureIDs[14])
    f(a, c, user.IP, "creature15", user.Aquarium.CreatureIDs[15])
    f(a, c, user.IP, "creature16", user.Aquarium.CreatureIDs[16])
    f(a, c, user.IP, "creature17", user.Aquarium.CreatureIDs[17])
    f(a, c, user.IP, "creature18", user.Aquarium.CreatureIDs[18])
    f(a, c, user.IP, "creature19", user.Aquarium.CreatureIDs[19])

    c.JSON(http.StatusOK, gin.H{"status": "request exec"})
}

func (a *Aquarium)Signup(c *gin.Context) {
    log.Printf("Signup ", c.ClientIP())

    var user model.User
    user.IP = c.ClientIP()

    if err := c.ShouldBindJSON(&user.Aquarium); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    stmt, err := a.db.Prepare(`INSERT INTO users (ip) VALUES (?)`)
    if err != nil {
        log.Printf(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    defer stmt.Close()
    _, err = stmt.Exec(user.IP)
    if err != nil {
        log.Printf(err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "request exec"})
}

func (a *Aquarium)Login(c *gin.Context) {
    log.Printf("Login")
}
