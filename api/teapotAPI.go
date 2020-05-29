package api

import "github.com/gin-gonic/gin"

func (web *Web) TeapotRouteHandler(router *gin.Engine) {
	router.GET("/tea", web.APITeapot)
}

func (web *Web) APITeapot(c *gin.Context) {
	c.Status(418)
}
