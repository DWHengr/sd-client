package api

import (
	"github.com/gin-gonic/gin"
	"sd-client/httpclient"
	"sd-client/service/models"
)

func ServiceDiscoveryApi(engine *gin.Engine) {
	group := engine.Group("/api")
	group.GET("/list", list)
	group.GET("/list/self", selfList)
	group.GET("/list/no/self", noSelfList)
}

func queryByIsSelf(itemList []*models.ServiceInfo, isSelf bool) []*models.ServiceInfo {
	var queryItemList []*models.ServiceInfo
	for _, item := range itemList {
		if item.IsSelf == isSelf {
			queryItemList = append(queryItemList, item)
		}
	}
	return queryItemList
}

func list(c *gin.Context) {
	httpclient.Format(itemList, nil).Context(c)
}

func selfList(c *gin.Context) {
	httpclient.Format(queryByIsSelf(itemList, true), nil).Context(c)
}

func noSelfList(c *gin.Context) {
	httpclient.Format(queryByIsSelf(itemList, false), nil).Context(c)
}
