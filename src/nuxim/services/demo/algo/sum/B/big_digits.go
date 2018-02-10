package B

import (
	"math/rand"
	. "nuxim/defs/services/demo/algo"
	//. "nuxim/defs/util"
	"nuxim/util"
	"time"
)

func InitDigits(conf DigitSum) []int {
	a := make([]int, conf.Nums+1)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i < conf.Nums+1; i++ {
		a[i] = conf.Digit
		if conf.Rand > 0 {
			a[i] = r.Intn(10)
		}
	}
	a[0] = 0
	return a
}

func runSumTask(a, b, i int, preCh, ch chan int, r chan TaskResult) {
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
	r <- TaskResult{K: i, C: c}
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

func sumDigitsTest(conf DigitSum, a, b []int) []int {
	ch := make([]chan int, 4)
	for i := 0; i < 4; i++ {
		ch[i] = make(chan int)
	}
	r := make(chan TaskResult)
	go runSumTask(9, 9, 1, ch[0], ch[1], r)
	go runSumTask(9, 9, 2, ch[1], ch[2], r)
	go runSumTask(9, 9, 3, ch[2], ch[3], r)
	ch[3] <- 0
	x1 := <-r
	x2 := <-r
	x3 := <-r
	n := <-ch[0]
	util.Log("Received Result", x1, x2, x3)
	return []int{n, x1.C, x2.C, x3.C}
}

func runSumTasks(conf DigitSum, a, b, c []int, ch []chan int, start, end, step int, r chan TaskResult) int {
	//util.Log("Running Sum Tasks, Start", start, "End", end)
	total := 0
	for k := start; k < end; k++ {
		go runSumTask(a[k], b[k], k, ch[k-1], ch[k], r)
	}
	ch[end-1] <- step
	for {
		select {
		case x := <-r:
			c[x.K] = x.C
			total++
			if total == end-start {
				return <-ch[start-1]
			}
		}
	}
	return 0
}

func SumDigits(conf DigitSum, a, b []int) []int {
	c := make([]int, conf.Nums+1)
	ch := make([]chan int, conf.Nums+1)
	r := make(chan TaskResult)
	cnt := initTasksCount(conf)
	for i := 0; i < conf.Nums+1; i++ {
		ch[i] = make(chan int)
	}
	step := 0
	loops := conf.Nums / cnt
	for i := 0; i < loops; i++ {
		start := conf.Nums - (i+1)*cnt + 1
		end := conf.Nums - i*cnt + 1
		step = runSumTasks(conf, a, b, c, ch, start, end, step, r)
	}
	if conf.Nums > loops*cnt {
		start := 1
		end := conf.Nums - loops*cnt + 1
		step = runSumTasks(conf, a, b, c, ch, start, end, step, r)
	}
	c[0] = step
	return c
}
