package multipleMissingElements

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	arr             [20]int
	expectedMissing []int
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{6, 7, 8, 9, 11, 12, 15, 16},
			[]int{10, 13, 14},
		},
		{
			[20]int{3, 4, 6, 7, 9, 10, 11, 12},
			[]int{5, 8},
		},
		{
			[20]int{1, 2, 3, 4, 5},
			[]int(nil),
		},
		{
			[20]int{7, 9, 10},
			[]int{8},
		},
		{
			[20]int{1, 2, 3, 4, 5, 6, 7, 8, 10},
			[]int{9},
		},
	}
}

var (
	impls []ArraysMultipleMissing
)

func TestMultipleMissingElements(t *testing.T) {
	impls = append(impls,
		&ImplS1A{},
		&ImplS2H{},
	)

	tests := testData()
	t.Logf("[#tests=%d]", len(tests))

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualOpt := impl.arraysMultipleMissing(tt.arr)
			var pOrF string
			if assert.Equal(t, tt.expectedMissing, actualOpt) {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test #%d, %s, expected arr = %v", i+1, pOrF, tt.expectedMissing)
		}
	}
}
