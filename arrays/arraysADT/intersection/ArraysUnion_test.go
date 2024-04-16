package intersection

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
			"Elements are: 6 15 \n",
		},
	}
}

var (
	impls []ArraysIntersection
)

func TestArraysUnion(t *testing.T) {
	impls = append(impls, &ImplS1A{})

	tests := testData()
	t.Logf("[#tests=%d]", len(tests))

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualArr := impl.arraysIntersection(tt.arr, tt.arr2)
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
