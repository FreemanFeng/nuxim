package sdk

import (
	. "nuxim/defs/common"
	. "nuxim/defs/sdk"
	. "nuxim/defs/util"
	"nuxim/sdk/service/ucstart"
	"nuxim/util"
)

func dispatch(task Task, ch chan []byte, cache chan Cache) {
	done := make(chan int)
	defer close(done)
	switch task.Service {
	case UCSTART:
		go ucstart.RunService(task, ch, cache, done)
	default:
		ch <- []byte("Incorrect Request!")
		return
	}
	for {
		select {
		case <-done:
			return
		}
	}
}

func Run(config Config, rootPath string, quit chan int) {
	ch := make(chan Task)
	dataCh := make(chan []byte)
	m := map[string]int{}
	cacheList := []chan Cache{}
	c := config.Sdk
	go util.RunHttpServer(util.TaskHandler(ch, dataCh), c.Port, c.ReadTimeout, c.WriteTimeout, quit)
	for {
		select {
		case x := <-ch:
			_, ok := m[x.Service]
			if !ok {
				index := len(cacheList)
				m[x.Service] = index
				cacheList = append(cacheList, make(chan Cache))
				go util.RunCache(cacheList[index], quit)
			}
			index := m[x.Service]
			go dispatch(x, dataCh, cacheList[index])
		}
	}
}
