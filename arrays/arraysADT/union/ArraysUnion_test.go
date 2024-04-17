package union

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	arr         [20]int
	arr2        [20]int
	expectedArr string
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{1, 2, 3, 4, 5},
			[20]int{6, 7, 8, 9, 10},
			"Elements are: 1 2 3 4 5 6 7 8 9 10 \n",
		},
		{
			[20]int{10, 20, 50, 100, 200},
			[20]int{10, 40, 90, 100, 200},
			"Elements are: 10 20 40 50 90 100 200 \n",
		},
		{
			[20]int{-50, 75, 100, 150, 500},
			[20]int{75, 99, 100, 160, 250},
			"Elements are: -50 75 99 100 150 160 250 500 \n",
		},
		{
			[20]int{10, 50, 100, 500, 1000},
			[20]int{10, 50, 100, 500, 1000},
			"Elements are: 10 50 100 500 1000 \n",
		},
		{
			[20]int{100, 250, 900, 1250, 1500},
			[20]int{100, 900, 1200, 1250, 1600},
			"Elements are: 100 250 900 1200 1250 1500 1600 \n",
		},
	}
}

var (
	impls []ArraysUnion
)

func TestArraysUnion(t *testing.T) {
	impls = append(impls, &ImplS1A{})

	tests := testData()
	t.Logf("[#tests=%d]", len(tests))

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualArr := impl.arraysUnion(tt.arr, tt.arr2)
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
