package ucstart

import (
	. "nuxim/defs/util"
	"nuxim/util"
)

func mock_cms(task Task, ch chan []byte, cache chan Cache, done chan int) {
	for k, v := range task.Params {
		util.Log("k:", k, " v:", v)
	}
}
