package main

import (
	"flag"
	"os"
	"os/signal"
	"sd-client/api"
	"sd-client/config"
	"sd-client/httpclient"
	"sd-client/job"
	"sd-client/logger"
	utils2 "sd-client/utils"
	"syscall"
)

var dirPath = flag.String("dir", "./", "-dir project path")
var configPath = flag.String("config", "config.yml", "-config Configuration File Address")

func main() {

	flag.Parse()
	conf, err := config.NewConfig(*dirPath + *configPath)
	conf.WorkDir = *dirPath
	if err != nil {
		panic(err)
	}
	err = logger.New(&conf.Log)
	if err != nil {
		panic(err)
	}
	utils2.SendGratutiousArp()
	if err != nil {
		panic(err)
	}
	router, err := api.NewRouter(conf)
	httpclient.NewClient(&conf.Http)
	if err != nil {
		panic(err)
	}
	//开始定时任务
	go job.StartJob(conf.Sd.Job)
	go router.Run()
	logger.Logger.Info("running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			//	router.Close()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
