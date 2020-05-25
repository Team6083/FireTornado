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
	fmt.Println("Connection start")

	dialInfo, _ := mgo.ParseURL("firstmongo-shard-00-01-3qdnz.mongodb.net:27017,firstmongo-shard-00-02-3qdnz.mongodb.net:27017")

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

	db := model.Database{Session: session, DialInfo: dialInfo, DB: session.DB("firstMongo")}
	web := api.Web{DB: &db}

	engine := gin.New()
	engine.Use(gin.Logger())

	web.RouteHandler(engine)
}
