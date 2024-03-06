package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sd-client/service/models"
)

type ServiceList struct {
	Services []*models.ServiceInfo `json:"services"`
}

var JsonFileName = "data.json"

// 读取 JSON 文件
func LoadJSONFile() ([]*models.ServiceInfo, error) {
	// 尝试打开文件
	file, err := os.Open(JsonFileName)
	// 如果文件不存在则创建新文件
	if os.IsNotExist(err) {
		fmt.Println("JSON file not found. Creating a new one.")
		newFile, err := os.Create(JsonFileName)
		if err != nil {
			return nil, err
		}
		defer newFile.Close()
		// 初始化一个空的 Data 结构体
		newData := &ServiceList{
			Services: []*models.ServiceInfo{},
		}
		// 写入空的 JSON 数据到新文件
		jsonData, err := json.MarshalIndent(newData, "", "    ")
		if err != nil {
			return nil, err
		}
		if _, err := newFile.Write(jsonData); err != nil {
			return nil, err
		}
		return newData.Services, nil
	} else if err != nil {
		// 如果发生其他错误则返回错误
		return nil, err
	}
	defer file.Close()

	// 读取文件内容
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// 将 JSON 数据解析到结构体中
	var result ServiceList
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result.Services, nil
}

// 写入 JSON 文件
func WriteJSONFile(data []*models.ServiceInfo) error {
	newData := ServiceList{
		Services: data,
	}
	jsonData, err := json.MarshalIndent(newData, "", "    ")
	if err != nil {
		return err
	}
	// 将数据写入文件
	err = ioutil.WriteFile(JsonFileName, jsonData, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
