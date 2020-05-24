package main

import (
	"crypto/tls"
	"fmt"
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

	//db := model.Database{session, dialInfo, session.DB("firstMongo")}
}
