package job

import (
	"os"
	"sd-client/api"
	"sd-client/client"
	"sd-client/service/models"
	"sd-client/utils"
)

func SyncCloudTask() {
	cloudItemList, _ := client.GetCloudItemList()
	var itemMap = make(map[string]*models.ServiceInfo)
	var itemCloudIdMap = make(map[string]bool)
	//获取云端服务对应的id
	for _, item := range cloudItemList {
		itemCloudIdMap[item.Id] = true
	}
	//是否有内容变化
	isUpdate := false

	var newItemList []*models.ServiceInfo
	itemList := api.GetItemList()
	//没删除的，才做修改，移除删除的内容
	for _, item := range itemList {
		//手动修改的不删除
		if itemCloudIdMap[item.Id] || item.IsManuallyModify {
			itemMap[item.Id] = item
			newItemList = append(newItemList, item)
		} else {
			isUpdate = true
		}
	}
	if isUpdate {
		itemList = newItemList
		api.SetItemList(newItemList)
	}
	//当前主机名
	hostname, _ := os.Hostname()
	for _, cloudItem := range cloudItemList {
		item := itemMap[cloudItem.Id]
		isSelf := false
		if hostname == cloudItem.Name {
			isSelf = true
		}
		//手动修改优先级高于云端同步
		if item != nil && item.IsManuallyModify {
			item.IsSelf = isSelf
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
					IsSelf:           isSelf,
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
				item.IsSelf = isSelf
			}
		}
	}
	if isUpdate {
		api.PersistentItemList()
		// 调用回调接口
		go client.CllAllHost(itemList)
	}
}
