package algo

import (
	. "nuxim/defs/util"
	"nuxim/services/demo/algo/sum"
	"nuxim/util"
)

func RunService(params string, done, quit chan int) {
	util.Log("Running Service For demo with params", params)
	sum.SumBigDigits(params, done)
	done <- SUCCESSFUL
}
