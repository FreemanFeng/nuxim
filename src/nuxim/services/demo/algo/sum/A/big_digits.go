package A

import (
	"math/rand"
	. "nuxim/defs/services/demo/algo"
	//. "nuxim/defs/util"
	"nuxim/util"
	"time"
)

func runSumTask(a, b, i int, preCh, ch, r chan int) {
	c := a + b
	n := 0
	//util.Log("Task", i, "Result", c)
	x := <-ch
	//util.Log("Task", i, "Received Signal", x)
	c += x
	if c > 9 {
		n = 1
		c -= 10
	}
	r <- c
	preCh <- n
	//util.Log("Task", i, "Done")
}

func initTasksCount(conf DigitSum) int {
	cnt := conf.Nums
	if conf.Tasks < conf.Nums {
		cnt = conf.Tasks
	}
	if cnt > MAX_TASKS {
		cnt = MAX_TASKS
	}
	return cnt
}

func initDigit(conf DigitSum, s *rand.Rand) int {
	if conf.Rand > 0 {
		return s.Intn(10)
	}
	return conf.Digit
}

func runSumTasks(conf DigitSum, cnt, step int, s *rand.Rand, r chan int) int {
	ch := make([]chan int, cnt+1)
	for i := 0; i < cnt+1; i++ {
		ch[i] = make(chan int)
	}
	total := 0
	for k := 1; k < cnt+1; k++ {
		a := initDigit(conf, s)
		b := initDigit(conf, s)
		go runSumTask(a, b, k, ch[k-1], ch[k], r)
	}
	ch[cnt] <- step
	for {
		select {
		case <-r:
			total++
			if total == cnt {
				return <-ch[0]
			}
		}
	}
	return 0
}

func SumDigits(conf DigitSum) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	r := make(chan int)
	cnt := initTasksCount(conf)
	step := 0
	loops := conf.Nums / cnt
	for i := 0; i < loops; i++ {
		step = runSumTasks(conf, cnt, step, s, r)
	}
	if conf.Nums > loops*cnt {
		step = runSumTasks(conf, conf.Nums-loops*cnt, step, s, r)
	}
	util.Log("Step", step)
}
