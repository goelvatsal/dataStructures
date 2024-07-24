package merger

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	arr          [20]int
	arr2         [20]int
	expectedarr3 string
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{2, 6, 10, 15, 25},
			[20]int{3, 4, 7, 18, 20},
			"Elements are: 2 3 4 6 7 10 15 18 20 25 \n",
		},
		{
			[20]int{1, 3, 5, 7, 9},
			[20]int{2, 4, 6, 8, 10},
			"Elements are: 1 2 3 4 5 6 7 8 9 10 \n",
		},
		{
			[20]int{20, 40, 90, 100, 200},
			[20]int{10, 50, 70, 91, 140},
			"Elements are: 10 20 40 50 70 90 91 100 140 200 \n",
		},
		{
			[20]int{-10, 20, 40, 200, 500},
			[20]int{-50, 10, 30, 70, 400},
			"Elements are: -50 -10 10 20 30 40 70 200 400 500 \n",
		},
		{
			[20]int{-500, -400, -300, -200, -100},
			[20]int{40, 50, 60, 70, 80},
			"Elements are: -500 -400 -300 -200 -100 40 50 60 70 80 \n",
		},
	}
}

var (
	impls []ArraysMerger
)

func TestMerger(t *testing.T) {
	impls = append(impls, &ImplS1A{})

	tests := testData()
	t.Logf("[#tests=%d]", len(tests))

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualArr := impl.arraysMerger(tt.arr, tt.arr2)
			var pOrF string
			if assert.Equal(t, tt.expectedarr3, actualArr) {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test #%d, %s, expected arr = %v", i+1, pOrF, tt.expectedarr3)
		}
	}
}
