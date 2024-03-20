package job

import (
	"sd-client/api"
	"sd-client/client"
	"sd-client/service/models"
	"sd-client/utils"
)

func SyncCloudTask() {
	cloudItemList, _ := client.GetCloudItemList()
	var itemMap = make(map[string]*models.ServiceInfo)
	itemList := api.GetItemList()
	for _, item := range itemList {
		itemMap[item.Mac] = item
	}
	//是否有内容变化
	isUpdate := false
	for _, cloudItem := range cloudItemList {
		item := itemMap[cloudItem.Mac]
		//手动修改优先级高于云端同步
		if item != nil && item.IsManuallyModify {
			continue
		} else {
			if item == nil {
				isUpdate = true
				api.AddItemList(&models.ServiceInfo{
					Ip:               cloudItem.Ip,
					Mac:              cloudItem.Mac,
					Name:             cloudItem.Name,
					Domain:           cloudItem.Domain,
					Id:               cloudItem.Id,
					Depid:            cloudItem.Depid,
					IsPing:           cloudItem.IsPing,
					IsManuallyModify: cloudItem.IsManuallyModify,
				})
			}
			if item != nil && !item.CompareContentsEqual(*cloudItem) {
				if item.Ip != cloudItem.Ip {
					utils.WriteBindZoneFile(cloudItem)
				}
				isUpdate = true
				item.Id = cloudItem.Id
				item.Ip = cloudItem.Ip
				item.Mac = cloudItem.Mac
				item.Domain = cloudItem.Domain
				item.Name = cloudItem.Name
				item.Depid = cloudItem.Depid
			}
		}
	}
	if isUpdate {
		api.PersistentItemList()
		// 调用回调接口
		go client.CllAllHost(itemList)
	}
}
