package main

import (
	"FireTornado/api"
	"FireTornado/model"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"net"
)

func main() {
	fmt.Println("Starting to connect")

	dialInfo, _ := mgo.ParseURL("firstmongo-shard-00-01-3qdnz.mongodb.net:27017")

	config := &tls.Config{}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (conn net.Conn, err error) {
		connection, err := tls.Dial("tcp", addr.String(), config)
		return connection, err
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connect success")

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// db & web setting
	db := model.Database{Session: session, DialInfo: dialInfo, DB: session.DB("firstMongo")}
	web := api.Web{DB: &db}

	engine := gin.New()
	engine.Use(gin.Logger())

	// use route handlers
	web.UserRouteHandler(engine)
	web.ItemRouteHandler(engine)

	// listen and serve
	if err = engine.Run(); err != nil {
		panic(err)
		return
	}
}
