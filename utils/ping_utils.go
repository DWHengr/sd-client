package utils

import (
	"fmt"
	"github.com/go-ping/ping"
	"sd-client/logger"
	"time"
)

func Ping(ip string) (float64, int64, error) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		logger.Logger.Error("创建Pinger失败:", err)
		return -1, -1, err
	}

	pinger.Count = 5 // 设置ping次数
	pinger.Timeout = time.Second * 10
	pinger.SetPrivileged(true)
	var packetLoss float64
	var avgRTT int64
	pinger.OnFinish = func(stats *ping.Statistics) {
		packetLoss = stats.PacketLoss
		avgRTT = stats.AvgRtt.Milliseconds()
	}
	err = pinger.Run()
	if err != nil {
		return 0, 0, fmt.Errorf("Ping失败: %v", err)
	}
	return packetLoss, avgRTT, nil
}
