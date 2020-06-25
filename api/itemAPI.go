package api

import (
	"FireTornado/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (web *Web) ItemRouteHandler(engine *gin.Engine) {
	routerGroup := engine.Group("/item")
	routerGroup.GET("", web.APIReadItem)
	routerGroup.GET("/group", web.APIReadItemsByCategory)
	routerGroup.GET("/all", web.APIReadItems)
	routerGroup.POST("", web.APICreateItem)
	routerGroup.PUT("", web.APIUpdateItem)
	routerGroup.DELETE("", web.APIDeleteItem)
	routerGroup.GET("/search", web.APISearchItem)
}

func (web *Web) APICreateItem(c *gin.Context) {
	item := model.Item{Id: bson.NewObjectId()}

	if err := c.ShouldBind(&item); err != nil {
		panic(err)
		return
	}

	change, err := web.DB.SaveItem(&item)
	if err != nil {
		panic(err)
		return
	}

	c.JSON(201, change)
}

func (web *Web) APIReadItems(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	items, err := web.DB.GetAllItem()
	if err != nil {
		panic(err)
		return
	}

	c.JSON(200, items)
}

func (web *Web) APIReadItemsByCategory(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	name := c.Query("name")

	items, err := web.DB.GetItemsByCategory(name)
	if err != nil {
		panic(err)
		return
	}

	c.JSON(200, items)
}

func (web *Web) APIReadItem(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	id := c.Query("id")

	if !bson.IsObjectIdHex(id) {
		c.Status(400)
		return
	}

	item, err := web.DB.GetItemById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.JSON(404, "item not found")
			return
		}
		panic(err)
		return
	}

	c.JSON(200, item)
}

func (web *Web) APIUpdateItem(c *gin.Context) {
	id := c.Query("id")

	if !bson.IsObjectIdHex(id) {
		c.Status(400)
		return
	}

	item := model.Item{Id: bson.ObjectIdHex(id)}

	if err := c.ShouldBind(&item); err != nil {
		panic(err)
		return
	}

	change, err := web.DB.SaveItem(&item)
	if err != nil {
		panic(err)
		return
	}

	c.JSON(200, change)
}

func (web *Web) APIDeleteItem(c *gin.Context) {
	id := c.Query("id")

	if !bson.IsObjectIdHex(id) {
		c.Status(400)
		return
	}

	err := web.DB.DeleteItem(id)
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

func (web *Web) APISearchItem(c *gin.Context) {
	name := c.Query("name")

	items, err := web.DB.SearchItem(name)
	if err != nil {
		panic(err)
		return
	}

	c.JSON(200, items)
}
