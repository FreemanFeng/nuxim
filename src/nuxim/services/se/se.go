package se

import (
	. "nuxim/defs/common"
	. "nuxim/defs/services/se"
	"nuxim/services/se/stock"
	"nuxim/util"
)

func Run(job CronJob, params string, done, quit chan int) {
	switch job.Service {
	case SE_STOCK:
		stock.RunService(params, done, quit)
	default:
		util.Log("[FATAL ERROR]Unknown Service ID")
	}
}
