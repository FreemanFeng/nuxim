package stock

import (
	. "nuxim/defs/util"
	"nuxim/util"
)

func RunService(params string, done, quit chan int) {
	util.Log("Running Service For stock with params", params)
	done <- SUCCESSFUL
}
