package main

import (
	"FireTornado/api"
	"FireTornado/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"os"
)

func main() {
	fmt.Println("Starting to connect")

	// dial mongodb
	databaseURL := os.Getenv("DATABASE_URL")
	dailInfo, err := mgo.ParseURL(databaseURL)
	if err != nil {
		panic(err)
	}

	session, err := mgo.Dial(databaseURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connect success")

	defer session.Close()

	// db & web setting
	db := model.Database{Session: session, DB: session.DB(dailInfo.Database)}
	web := api.Web{DB: &db}

	// set up engine
	engine := gin.New()
	engine.Use(gin.Logger())

	// use route handlers
	web.UserRouteHandler(engine)
	web.ItemRouteHandler(engine)
	web.TeapotRouteHandler(engine)

	// listen and serve
	if err = engine.Run(); err != nil {
		panic(err)
		return
	}
}
