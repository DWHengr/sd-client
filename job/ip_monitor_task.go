package job

import (
	"net"
	"sd-client/api"
	"sd-client/client"
	pkglogger "sd-client/logger"
	utils2 "sd-client/utils"
)

type MacInfo struct {
	Mac string
	Ip  string
}

var macInfosMap = make(map[string]string)

func IpMonitorTask() {
	var itemCloudIdMap = make(map[string]bool)
	itemList := api.GetItemList()
	for _, item := range itemList {
		itemCloudIdMap[item.Mac] = true
	}
	// 获取本地网络接口信息
	iFaces, err := net.Interfaces()
	if err != nil {
		pkglogger.Logger.Error("获取网络接口信息出错: ", err)
		return
	}
	isUpdate := false
	// 遍历网络接口信息
	for _, iFace := range iFaces {
		// 过滤非物理网卡
		if iFace.Flags&net.FlagLoopback != 0 || iFace.Flags&net.FlagUp == 0 {
			continue
		}
		// 获取物理网卡的MAC地址
		mac := iFace.HardwareAddr
		if mac == nil {
			continue
		}
		// 获取物理网卡的IP地址信息
		addrs, err := iFace.Addrs()
		if err != nil {
			pkglogger.Logger.Error("获取IP地址信息出错: ", err)
			continue
		}
		ip := ""
		//mac地址和ip信息
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				break
			}
		}
		macStr := mac.String()
		//mac地址对应的ip信息变化更新本地和远程
		if itemCloudIdMap[mac.String()] {
			err := client.ModifyCloudServiceIp(mac.String(), ip)
			if err == nil {
				macInfosMap[macStr] = ip
				isUpdate = true
			}
		}
	}
	if isUpdate {
		utils2.SendGratutiousArp()
		go SyncCloudTask()
	}

}
