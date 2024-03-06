package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sd-client/client"
	pkglogger "sd-client/logger"
	"sd-client/service/models"
	"sd-client/utils"
	"strings"
)

var itemList []*models.ServiceInfo

func GetItemList() []*models.ServiceInfo {
	return itemList
}

func SetItemList(newItemList []*models.ServiceInfo) {
	itemList = newItemList
}
func AddItemList(newItem *models.ServiceInfo) {
	itemList = append(itemList, newItem)
}

func PersistentItemList() {
	utils.WriteJSONFile(itemList)
}

type Query struct {
	Name string `json:"name"`
	Mac  string `json:"mac"`
}

func ServiceDiscovery(engine *gin.Engine) {
	//加载json内容
	itemList, _ = utils.LoadJSONFile()
	group := engine.Group("/")
	group.GET("/", FindAanQuery)
	group.GET("/delete/manually/:id", DelManuallyService)
	group.POST("/update", UpdateService)
}

func FindAanQuery(c *gin.Context) {
	var query = Query{
		Name: c.Query("name"),
		Mac:  c.Query("mac"),
	}
	if err := c.Bind(&query); err != nil {
		pkglogger.Logger.Error(err)
	}
	var queryItemList []*models.ServiceInfo
	if len(query.Name) > 0 || len(query.Mac) > 0 {
		for _, item := range itemList {
			if (len(query.Name) > 0 && strings.Contains(item.Name, query.Name)) ||
				(len(query.Mac) > 0 && strings.Contains(item.Mac, query.Mac)) {
				queryItemList = append(queryItemList, item)
			}
		}
	} else {
		queryItemList = itemList
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"items": queryItemList,
	})
}

func DelManuallyService(c *gin.Context) {
	id := c.Param("id")
	if len(id) <= 0 {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}
	for _, item := range itemList {
		if item.Id == id {
			item.IsManuallyModify = false
			break
		}
	}
	c.Redirect(http.StatusFound, "/")
}

func UpdateService(c *gin.Context) {
	// 更新
	var newItem models.ServiceInfo
	if err := c.Bind(&newItem); err != nil {
		pkglogger.Logger.Error(err, "无效的请求")
		return
	}
	for i, item := range itemList {
		if item.Ip != newItem.Ip {
			utils.WriteBindZoneFile(&newItem)
		}
		if item.Id == newItem.Id {
			itemList[i] = &newItem
			itemList[i].IsManuallyModify = true
			break
		}
	}
	utils.WriteJSONFile(itemList)
	// 调用回调接口
	client.CllAllHost(itemList)
	c.Redirect(http.StatusFound, "/")
}
