package job

import (
	"sd-client/config"
	"time"
	_ "time"
)

func StartJob(job config.Job) {
	var ipMonitorTicker = time.NewTicker(time.Duration(job.IpMonitor) * time.Second)
	var ipPingTicker = time.NewTicker(time.Duration(job.IpPing) * time.Second)
	var syncCloudTicker = time.NewTicker(time.Duration(job.SyncCloud) * time.Second)
	for {
		select {
		case <-ipMonitorTicker.C:
			go IpMonitorTask()
		case <-ipPingTicker.C:
			go IpPingTask()
		case <-syncCloudTicker.C:
			go SyncCloudTask()
		}

	}
	defer ipMonitorTicker.Stop()
	defer ipPingTicker.Stop()
	defer syncCloudTicker.Stop()
}
