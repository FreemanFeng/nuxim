package C

import (
	"math/rand"
	. "nuxim/defs/services/demo/algo"
	//. "nuxim/defs/util"
	"nuxim/util"
	"time"
)

func sumDigits(a, b, c []int, k int, r chan int) {
	step := k
	for i := len(a) - 1; i > 0; i-- {
		n := a[i] + b[i] + step
		step = 0
		if n > 9 {
			step = 1
			n -= 10
		}
		c[i] = n
	}
	r <- step
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

func initDigits(conf DigitSum, s *rand.Rand, cnt int) []int {
	a := make([]int, cnt+1)
	for i := 1; i < cnt+1; i++ {
		a[i] = conf.Digit
		if conf.Rand > 0 {
			a[i] = r.Intn(10)
		}
	}
	a[0] = 0
	return a
}

func runSumTask(conf DigitSum, i, cnt, step int, s *rand.Rand, r chan int) int {
	total := 0
	a := initDigits(conf, s)
	b := initDigits(conf, s)
	c := make([]int, cnt+1)
	go sumDigits(a, b, c, step, r)
	util.Log("Task", i, "Result", c)
	return <-r
}

func Run(conf DigitSum) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	r := make(chan int)
	cnt := initTasksCount(conf)
	step := 0
	loops := conf.Nums / cnt
	for i := 0; i < loops; i++ {
		step = runSumTask(conf, i, cnt, step, s, r)
	}
	if conf.Nums > loops*cnt {
		step = runSumTask(conf, loops, conf.Nums-loops*cnt, step, s, r)
	}
	util.Log("Step", step)
}
