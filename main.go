package main

import (
	"FireTornado/api"
	"FireTornado/model"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"net"
	"os"
)

func main() {
	fmt.Println("Starting to connect")

	// dialInfo setting
	dialInfo, _ := mgo.ParseURL(os.Getenv("uri"))

	config := &tls.Config{}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (conn net.Conn, err error) {
		connection, err := tls.Dial("tcp", addr.String(), config)
		return connection, err
	}

	dialInfo.Username = os.Getenv("user")
	dialInfo.Password = os.Getenv("pass")

	// dial
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connect success")

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// db & web setting
	db := model.Database{Session: session, DialInfo: dialInfo, DB: session.DB(os.Getenv("db"))}
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
