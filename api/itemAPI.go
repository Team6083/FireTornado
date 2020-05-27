package api

import (
	"FireTornado/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (web *Web) ItemRouteHandler(engine *gin.Engine) {
	routerGroup := engine.Group("/item")
	routerGroup.POST("", web.APICreateItem)
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

func (web *Web) APIReadAllItems(c *gin.Context) {
	items, err := web.DB.GetAllItem()
	if err != nil {
		panic(err)
		return
	}

	c.JSON(200, items)
}

func (web *Web) APIReadItem(c *gin.Context) {
	id := c.Query("id")

	item, err := web.DB.GetItemById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.Status(404)
			return
		}
		panic(err)
		return
	}

	c.JSON(200, item)
}

func (web *Web) APIUpdateItem(c *gin.Context) {
	id := c.Query("id")

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

	err := web.DB.DeleteItem(model.Item{Id: bson.ObjectIdHex(id)})
	if err != nil {
		panic(err)
		return
	}

	c.Status(200)
}
