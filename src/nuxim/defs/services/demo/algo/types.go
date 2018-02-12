package algo

type DigitSum struct {
	Total, Nums, Tasks, Digit, Rand int
}

type SumTask struct {
	K         int
	PreCh, Ch chan int
}

type TaskResult struct {
	K, C int
}
