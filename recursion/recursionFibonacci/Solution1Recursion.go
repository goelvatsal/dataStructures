package FibonacciElements

var slc []int

type ImplS1R struct{}

func fibInit() {
	slc = make([]int, 20)
	for i := range slc {
		slc[i] = -1
	}
}

func (s1 ImplS1R) computeFibonacci(n int) int {
	if n <= 1 {
		slc[n] = n
		return n
	} else {
		if slc[n-2] == -1 {
			slc[n-2] = s1.computeFibonacci(n - 2)
		}
		if slc[n-1] == -1 {
			slc[n-1] = s1.computeFibonacci(n - 1)
		}
		slc[n] = slc[n-2] + slc[n-1]
		return slc[n-2] + slc[n-1]
	}
}
