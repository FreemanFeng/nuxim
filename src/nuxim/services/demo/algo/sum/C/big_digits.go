package C

import (
	"math/rand"
	. "nuxim/defs/services/demo/algo"
	//"nuxim/util"
	"time"
)

func calculateSum(conf DigitSum, r *rand.Rand, c []int, cnt int) int {
	step := 0
	for i := cnt - 1; i >= 0; i-- {
		a := initDigit(conf, r)
		b := initDigit(conf, r)
		n := a + b + step
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

func runSumTask(conf DigitSum, r *rand.Rand, cnt int, preCh, ch chan int) {
	c := make([]int, cnt)
	step := calculateSum(conf, r, c, cnt)
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

func initDigit(conf DigitSum, r *rand.Rand) int {
	if conf.Rand > 0 {
		return r.Intn(10)
	}
	return conf.Digit
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
		go runSumTask(conf, r, cnt, ch[i-1], ch[i])
	}
	// 最后一位,无进位
	ch[loops+1] <- 0
	start := 1
	if conf.Total > loops*cnt {
		start = 0
		go runSumTask(conf, r, conf.Total-loops*cnt, ch[0], ch[1])
	}
	// 读取进位信息, 0或1
	<-ch[start]
}
