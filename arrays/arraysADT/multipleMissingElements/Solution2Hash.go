package multipleMissingElements

type ImplS2H struct{}

func highestNum(arr [20]int) int {
	num := 0
	for _, v := range arr {
		if v > num {
			num = v
		}
	}
	return num
}

func lowestNum(arr [20]int) int {
	s := arr[0]
	for _, num := range arr[1:] {
		if num < s && num != 0 {
			s = num
		}
	}
	return s
}

func (s2 ImplS2H) arraysMultipleMissing(arr [20]int) []int {
	l := lenCalc(arr)
	hash := make([]int, 20)
	for i := 0; i < l; i++ {
		hash[arr[i]]++
	}

	var fin []int
	low := lowestNum(arr)
	high := highestNum(arr)
	for j := low; j <= high; j++ {
		if hash[j] == 0 {
			fin = append(fin, j)
		}
	}
	return fin
}

func (s2 ImplS2H) arraysMultipleMissingHash(arr [20]int) []int {
	var arr3 Array
	return arr3.multipleMissingElements(arr)
}
