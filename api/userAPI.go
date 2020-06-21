package api

import (
	"FireTornado/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (web *Web) UserRouteHandler(router *gin.Engine) {
	usersGroup := router.Group("/user")
	usersGroup.GET("", web.APIReadUser)
	usersGroup.GET("/all", web.APIReadUsers)
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

func (web *Web) APIReadUsers(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	users, err := web.DB.GetAllUsers()
	if err != nil {
		panic(err)
		return
	}

	c.JSON(200, users)
}

func (web *Web) APIReadUser(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	id := c.Query("id")

	if !bson.IsObjectIdHex(id) {
		c.Status(400)
		return
	}

	user, err := web.DB.GetUserById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.JSON(404, "user not found")
			return
		}
		panic(err)
		return
	}

	c.JSON(200, user)
}

func (web *Web) APIUpdateUser(c *gin.Context) {
	id := c.Query("id")

	if !bson.IsObjectIdHex(id) {
		c.Status(400)
		return
	}

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

	if !bson.IsObjectIdHex(id) {
		c.Status(400)
		return
	}

	err := web.DB.DeleteUser(model.User{Id: bson.ObjectIdHex(id)})
	if err != nil {
		if err == mgo.ErrNotFound {
			c.JSON(404, "item not found")
			return
		}
		panic(err)
		return
	}

	c.Status(200)
}
