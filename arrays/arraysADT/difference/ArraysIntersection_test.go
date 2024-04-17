package difference

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	arr, arr2   [20]int
	expectedArr string
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{2, 6, 10, 15, 25},
			[20]int{3, 6, 7, 15, 20},
			"Elements are: 2 10 \n",
		},
		{
			[20]int{4, 12, 45, 60, 75},
			[20]int{3, 19, 40, 60, 75},
			"Elements are: 4 12 45 \n",
		},
		{
			[20]int{5, 20, 23, 44, 50},
			[20]int{2, 20, 30, 43, 54},
			"Elements are: 5 23 44 50 \n",
		},
		{
			[20]int{1, 3, 5, 7, 9},
			[20]int{2, 4, 6, 8, 10},
			"Elements are: 1 3 5 7 9 \n",
		},
		{
			[20]int{5, 15, 40, 90, 100},
			[20]int{5, 15, 40, 90, 100},
			"Elements are: \n",
		},
	}
}

var (
	impls []ArraysDifference
)

func TestArraysDifference(t *testing.T) {
	impls = append(impls, &ImplS1A{})
	
	tests := testData()
	t.Logf("[#tests=%d]", len(tests))
	
	for _, impl := range impls {
		t.Logf("Using: %T", impl)
		
		for i, tt := range tests {
			actualArr := impl.arraysDifference(tt.arr, tt.arr2)
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
