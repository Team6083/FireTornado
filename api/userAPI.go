package api

import (
	"FireTornado/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func (web *Web) UserRouteHandler(router *gin.Engine) {
	usersGroup := router.Group("/user")
	usersGroup.GET("", web.APIReadUser)
	usersGroup.POST("", web.APICreateUser)
	usersGroup.PUT("", web.APIUpdateUser)
	usersGroup.DELETE("", web.APIDeleteUser)
}

func (web *Web) APICreateUser(c *gin.Context) {
	user := model.User{Id: bson.NewObjectId()}

	err := c.ShouldBind(&user)
	if err != nil {
		panic(err)
		return
	}

	create, err := web.DB.SaveUser(user)
	if err != nil {
		panic(err)
		return
	}

	c.JSON(201, create)
}

func (web *Web) APIReadUser(c *gin.Context) {
	id := c.Query("id")

	read, err := web.DB.GetUserById(id)
	if err != nil {
		panic(err)
		return
	}

	c.JSON(200, read)
}

func (web *Web) APIUpdateUser(c *gin.Context) {
	id := c.Query("id")

	user := model.User{Id: bson.ObjectIdHex(id)}

	err := c.ShouldBind(&user)
	if err != nil {
		panic(err)
		return
	}

	update, err := web.DB.SaveUser(user)
	if err != nil {
		panic(err)
		return
	}

	c.JSON(200, update)
}

func (web *Web) APIDeleteUser(c *gin.Context) {
	id := c.Query("id")

	err := web.DB.DeleteUser(model.User{Id: bson.ObjectIdHex(id)})
	if err != nil {
		panic(err)
		return
	}

	c.Status(200)
}
