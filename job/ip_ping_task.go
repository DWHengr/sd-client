package job

import (
	"sd-client/api"
	"sd-client/client"
	"sd-client/logger"
	"sd-client/utils"
)

func IpPingTask() {
	arpTables := utils.GetARPTables()
	itemList := api.GetItemList()
	//是否有内容变化
	isUpdate := false
	for i, item := range itemList {
		beforeStatus := itemList[i].IsPing
		packetLoss, _, err := utils.Ping(item.Ip)
		//丢包率小于一半
		if err == nil && packetLoss < 50 {
			itemList[i].IsPing = true
		} else {
			if item.IsManuallyModify {
				continue
			}
			logger.Logger.Error("ping不通结果:", item.Ip, err, packetLoss)
			itemList[i].IsPing = false
			//查询arp表
			ip := arpTables[item.Mac]
			if len(ip) > 0 {
				packetLoss, _, err := utils.Ping(ip)
				if err == nil && packetLoss < 50 {
					itemList[i].Ip = ip
					itemList[i].IsPing = true
					utils.WriteBindZoneFile(itemList[i])
					// 调用接口修改云端
					client.ModifyCloudServiceIp(item.Mac, ip)
				}
			}
		}
		if beforeStatus != itemList[i].IsPing {
			isUpdate = true
		}
	}
	if isUpdate {
		api.PersistentItemList()
		go client.CllAllHost(itemList)
	}
}
