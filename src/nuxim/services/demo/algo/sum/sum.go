package sum

import (
	. "nuxim/defs/services/demo/algo"
	. "nuxim/defs/util"
	"nuxim/services/demo/algo/sum/A"
	"nuxim/services/demo/algo/sum/B"
	"nuxim/services/demo/algo/sum/C"
	"nuxim/util"
	"os"
)

func sumBigDigitsA(params string, done chan int) {
	util.Log("Running Sum Big Digits with Params", params)
	data := util.ParseParams(params)
	conf := DigitSum{}
	conf.Nums = util.GetParamInt(data, PARAM_NUMS, DEFAULT_NUMS)
	conf.Tasks = util.GetParamInt(data, PARAM_TASKS, DEFAULT_TASKS)
	conf.Digit = util.GetParamInt(data, PARAM_DIGIT, DEFAULT_DIGIT)
	conf.Rand = util.GetParamInt(data, PARAM_RAND, DEFAULT_RAND)
	util.Log("Config:", conf)
	A.SumDigits(conf)
	util.Log("Done")
}
func sumBigDigitsB(params string, done chan int) {
	util.Log("Running Sum Big Digits with Params", params)
	data := util.ParseParams(params)
	conf := DigitSum{}
	conf.Nums = util.GetParamInt(data, PARAM_NUMS, DEFAULT_NUMS)
	conf.Tasks = util.GetParamInt(data, PARAM_TASKS, DEFAULT_TASKS)
	conf.Digit = util.GetParamInt(data, PARAM_DIGIT, DEFAULT_DIGIT)
	conf.Rand = util.GetParamInt(data, PARAM_RAND, DEFAULT_RAND)
	util.Log("Config:", conf)
	a := B.InitDigits(conf)
	//util.Log("A:", a)
	b := B.InitDigits(conf)
	//util.Log("B:", b)
	util.Log("a[0:10]:", a[0:10], "b[last-10:last]", b[conf.Nums-9:conf.Nums+1])
	c := B.SumDigits(conf, a, b)
	util.Log("Result[0:10]:", c[0:10], "Result[last-10:last]", c[conf.Nums-9:conf.Nums+1])
}
func sumBigDigitsC(params string, done chan int) {
	util.Log("Running Sum Big Digits with Params", params)
	data := util.ParseParams(params)
	conf := DigitSum{}
	conf.Nums = util.GetParamInt(data, PARAM_NUMS, DEFAULT_NUMS)
	conf.Tasks = util.GetParamInt(data, PARAM_TASKS, DEFAULT_TASKS)
	conf.Digit = util.GetParamInt(data, PARAM_DIGIT, DEFAULT_DIGIT)
	conf.Rand = util.GetParamInt(data, PARAM_RAND, DEFAULT_RAND)
	util.Log("Config:", conf)
	C.SumDigits(conf)
	util.Log("Done")
}
func SumBigDigits(params string, done chan int) {
	//sumBigDigitsA(params, done)
	//sumBigDigitsB(params, done)
	sumBigDigitsC(params, done)
	done <- SUCCESSFUL
	os.Exit(0)
}
