package client

import (
	"errors"
	"sd-client/config"
	"sd-client/httpclient"
	"sd-client/service/models"
)

// GetCloudItemList 获取云端服务列表
func GetCloudItemList() ([]*models.ServiceInfo, error) {
	allConfig, _ := config.GetAllConfig()
	host := allConfig.Sd.CloudHost
	var cloudItemList []*models.ServiceInfo
	var result httpclient.R
	result.Data = &cloudItemList
	err := httpclient.GET(host+"/api/list", &result)
	if err != nil {
		return nil, err
	}
	if result.Code == 0 {
		return cloudItemList, nil
	}
	return nil, errors.New(result.Msg)
}

// ModifyCloudServiceIp 更新修改云端mac对应的ip
func ModifyCloudServiceIp(mac string, ip string) error {
	allConfig, _ := config.GetAllConfig()
	host := allConfig.Sd.CloudHost
	var param = map[string]string{
		"mac": mac,
		"ip":  ip,
	}
	var result httpclient.R
	var err error
	for i := 0; i < 3; i++ {
		err = httpclient.POST(nil, host+"/api/modify/service/ip", param, &result)
		if err != nil {
			continue
		}
		if result.Code != 0 {
			err = errors.New(result.Msg)
			continue
		}
		return nil
	}
	return err
}
