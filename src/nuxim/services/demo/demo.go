package demo

import (
	. "nuxim/defs/common"
	. "nuxim/defs/services/demo"
	"nuxim/services/demo/algo"
	"nuxim/util"
)

func Run(job CronJob, params string, done, quit chan int) {
	switch job.Service {
	case DEMO_ALGO:
		algo.RunService(params, done, quit)
	default:
		util.Log("[FATAL ERROR]Unknown Service ID")
	}
}
