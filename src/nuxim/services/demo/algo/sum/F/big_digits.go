package F

import (
	"math/rand"
	. "nuxim/defs/services/demo/algo"
	//"nuxim/util"
	"time"
)

func calculateSum(a, b, c []int) int {
	step := 0
	for i := len(a) - 1; i >= 0; i-- {
		n := a[i] + b[i] + step
		step = 0
		if n > 9 {
			step = 1
			n -= 10
		}
		c[i] = n
	}
	return step
}

func addStep(c []int) {
	step := 1
	for i := len(c) - 1; i >= 0; i-- {
		n := c[i] + step
		step = 0
		if n > 9 {
			step = 1
			n -= 10
		}
		c[i] = n
	}
}

func sumDigits(a, b, c []int, i int, preCh, ch chan int) {
	step := calculateSum(a, b, c)
	x := <-ch
	if x > 0 {
		addStep(c)
	}
	//util.Log("Task", i, "Result[0:10]", c[0:10], "Step", step)
	//util.Log("Task", i, "Result", c, "Step", step)
	preCh <- step
}

func initTasks(conf DigitSum) int {
	cnt := conf.Tasks
	if cnt > MAX_TASKS {
		cnt = MAX_TASKS
	}
	return cnt
}

func initDigits(conf DigitSum, r *rand.Rand, cnt int) []int {
	a := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		a[i] = conf.Digit
		if conf.Rand > 0 {
			a[i] = r.Intn(10)
		}
	}
	return a
}

func runSumTask(conf DigitSum, i, cnt int, r *rand.Rand, preCh, ch chan int) {
	a := initDigits(conf, r, cnt)
	b := initDigits(conf, r, cnt)
	c := make([]int, cnt)
	go sumDigits(a, b, c, i, preCh, ch)
}

func Run(conf DigitSum) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	tasks := initTasks(conf)
	loops := conf.Total / (tasks * conf.Nums)
	// 初始化channel数组
	ch := make([]chan int, loops*tasks+2)
	for i := 0; i < loops*tasks+2; i++ {
		ch[i] = make(chan int)
	}

	step := 0
	for i := loops - 1; i >= 0; i-- {
		for k := 2; k < tasks+2; k++ {
			runSumTask(conf, i*tasks+k, conf.Nums, r, ch[i*tasks+k-1], ch[i*tasks+k])
		}
		// 最后一位,无进位
		ch[(i+1)*tasks+1] <- step
		// 读取进位信息, 0或1
		step = <-ch[i*tasks+1]
	}
	if conf.Total > loops*tasks*conf.Nums {
		runSumTask(conf, 1, conf.Total-loops*tasks*conf.Nums, r, ch[0], ch[1])
		ch[1] <- step
		step = <-ch[0]
	}
}
