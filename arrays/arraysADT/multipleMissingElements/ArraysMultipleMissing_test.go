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
			[20]int{6, 7, 8, 9, 11, 12, 15, 16, 17, 18, 19},
			[]int{10, 13, 14},
		},
		{
			[20]int{100, 105, 110},
			[]int{101, 102, 103, 104, 106, 107, 108, 109},
		},
		{
			[20]int{1, 2, 3, 4, 5},
			[]int(nil),
		},
		{
			[20]int{20, 21, 22, 23, 25},
			[]int{24},
		},
		{
			[20]int{50, 59},
			[]int{51, 52, 53, 54, 55, 56, 57, 58},
		},
	}
}

var (
	impls []ArraysMultipleMissing
)

func TestMultipleMissingElements(t *testing.T) {
	impls = append(impls, &ImplS1A{})

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
