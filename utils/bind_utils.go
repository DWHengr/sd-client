package utils

import (
	"os"
	"os/exec"
	"sd-client/config"
	"sd-client/logger"
	"sd-client/service/models"
	"text/template"
)

func WriteBindZoneFile(item *models.ServiceInfo) {
	allConfig, err := config.GetAllConfig()
	bind := allConfig.Sd.Bind
	if !bind.Enable {
		return
	}
	// 从文件中读取模板内容
	tplBytes, err := os.ReadFile("bind_zone_tpl.txt")
	if err != nil {
		logger.Logger.Error("bind_zone_tpl file err: ", err)
		return
	}
	tplStr := string(tplBytes)
	// 创建文件
	file, err := os.Create(bind.ZonesDir + bind.ZoneFilePrefix + item.Domain + bind.ZoneFileSuffix)
	if err != nil {
		logger.Logger.Error("zone file crete err: ", err)
		return
	}
	defer file.Close()
	// 创建模板并解析
	t, err := template.New("bindZoneConfig").Parse(tplStr)
	if err != nil {
		logger.Logger.Error("tlp crete err: ", err)
		return
	}
	// 将模板应用到DNS配置结构体，并写入文件
	err = t.Execute(file, item)
	if err != nil {
		logger.Logger.Error("zone file write err: ", err)
		return
	}
	if len(bind.ReloadConfigCmd) >= 0 {
		cmd := exec.Command(bind.ReloadConfigCmd)
		// 执行命令并捕获输出
		if output, err := cmd.Output(); err != nil {
			logger.Logger.Error("Error executing command:", err)
		} else {
			logger.Logger.Info(bind.ReloadConfigCmd, ": ", output)
		}
	}
}
