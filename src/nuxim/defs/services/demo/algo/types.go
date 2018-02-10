package algo

type DigitSum struct {
	Nums, Tasks, Digit, Rand int
}

type SumTask struct {
	K         int
	PreCh, Ch chan int
}

type TaskResult struct {
	K, C int
}
