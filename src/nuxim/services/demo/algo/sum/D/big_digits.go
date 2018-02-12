package D

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
	preCh <- step
}

func initTaskNums(conf DigitSum) int {
	cnt := conf.Nums
	if cnt > MAX_NUMS {
		cnt = MAX_NUMS
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
	cnt := initTaskNums(conf)
	loops := conf.Total / cnt
	// 初始化channel数组
	ch := make([]chan int, loops+2)
	for i := 0; i < loops+2; i++ {
		ch[i] = make(chan int)
	}

	for i := 2; i < loops+2; i++ {
		runSumTask(conf, i, cnt, r, ch[i-1], ch[i])
	}
	// 最后一位,无进位
	ch[loops+1] <- 0
	start := 1
	if conf.Total > loops*cnt {
		start = 0
		runSumTask(conf, loops, conf.Total-loops*cnt, r, ch[0], ch[1])
	}
	// 读取进位信息, 0或1
	<-ch[start]
}
