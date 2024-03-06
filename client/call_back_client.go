package client

import (
	"sd-client/config"
	"sd-client/httpclient"
	"sd-client/logger"
	"sd-client/service/models"
)

// CllAllHost 回调全部的主机
func CllAllHost(itemList []*models.ServiceInfo) {
	allConfig, _ := config.GetAllConfig()
	hosts := allConfig.Sd.CallHosts
	var result interface{}
	for _, host := range hosts {
		for i := 1; i < 3; i++ {
			err := httpclient.POST(nil, host, itemList, &result)
			if err != nil {
				logger.Logger.Error("回调失败:", host, err)
			} else {
				break
			}
		}
	}
}
