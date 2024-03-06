package api

import (
	"github.com/gin-gonic/gin"
	"sd-client/httpclient"
)

func ServiceDiscoveryApi(engine *gin.Engine) {
	group := engine.Group("/api")
	group.GET("/list", list)
}

func list(c *gin.Context) {
	httpclient.Format(itemList, nil).Context(c)
}
