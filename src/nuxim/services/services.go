package services

import (
	. "nuxim/defs/common"
	. "nuxim/defs/services"
	. "nuxim/defs/util"
	"nuxim/services/demo"
	"nuxim/services/se"
	"nuxim/util"
)

func runTask(taskTimeout int, job CronJob, params string, quit chan int) {
	done := make(chan int)
	switch job.Platform {
	case PLATFORM_SE:
		go se.Run(job, params, done, quit)
	case PLATFORM_DEMO:
		go demo.Run(job, params, done, quit)
	}

	timeout := make(chan bool, 1)
	util.InitTimeoutTrigger(taskTimeout, timeout)
	select {
	case x := <-done:
		switch x {
		case SUCCESSFUL:
			util.Log(job.Platform, job.Service, params, "Done Successful!")
		default:
			util.Log(job.Platform, job.Service, params, "Failed!")
		}
	case <-timeout:
		util.Log("Timeout for running", job.Platform, job.Service, params)
	}
}

func runTasks(job CronJob, taskTimeout int, quit chan int) {
	for _, task := range job.Data {
		go runTask(taskTimeout, job, task.Params, quit)
	}
	<-quit
	quit <- 1
}

func Run(config Config, rootPath string, quit chan int) {
	jobs := util.InitCronJobs(config, rootPath, "services")
	c := config.Services
	for _, job := range jobs {
		//同一个调度类型（如按天，按小时或按分钟），可以有多次调度
		for _, start := range job.Start {
			go util.Schedule(job.CronType, start, job.Interval, quit,
				runTasks, job, c.TaskTimeout, quit)
		}
	}
	<-quit
	quit <- 1
}
