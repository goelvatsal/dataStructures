package negSort

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	base        [20]int
	expectedArr string
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{-10, 100, 50, 20, -45},
			"Elements are: -10 100 50 20 -45 \n",
		},
		{
			[20]int{100, 200, 300, 400, 500},
			"Elements are: 100 200 300 400 500 \n",
		},
		{
			[20]int{-300, -400, -500, -600, -700},
			"Elements are: -300 -400 -500 -600 -700 \n",
		},
		{
			[20]int{-25, 50, 100, -30, -43},
			"Elements are: -25 50 100 -30 -43 \n",
		},
		{
			[20]int{-100, 100, 200, -200, 300},
			"Elements are: -100 100 200 -200 300 \n",
		},
	}
}

var (
	impls []ArraysNegSorter
)

func TestNegSorter(t *testing.T) {
	impls = append(impls, &ImplS1A{})

	tests := testData()
	t.Logf("[#tests=%d]", len(tests))

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualArr := impl.arraysNegSorter(tt.base)
			var pOrF string
			if assert.Equal(t, tt.expectedArr, actualArr) {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test #%d, %s, expected arr = %v", i+1, pOrF, tt.expectedArr)
		}
	}
}
