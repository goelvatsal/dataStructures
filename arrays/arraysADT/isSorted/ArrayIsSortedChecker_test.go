package isSorted

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	base         [20]int
	expectedSort string
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{10, 20, 30, 40, 50},
			"Sorted!",
		},
		{
			[20]int{100, 99, 200, 300, 400},
			"Not Sorted!",
		},
		{
			[20]int{93, 34, 10, 55, 885},
			"Not Sorted!",
		},
		{
			[20]int{100, 200, 50, 25, 12},
			"Not Sorted!",
		},
		{
			[20]int{50, 90, 100, 105, 104},
			"Not Sorted!",
		},
	}
}

var (
	impls []ArraysIsSortedChecker
)

func TestIsSortedChecker(t *testing.T) {
	impls = append(impls, &ImplS1A{})

	tests := testData()
	t.Logf("[#tests=%d] \n", len(tests))

	assert.NotEmpty(t, impls, "no implementations present!")

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualSort := impl.arraysIsSorted(tt.base)
			var pOrF string
			if assert.Equal(t, tt.expectedSort, actualSort) {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test #%d, %s, sorted or not? = %v", i+1, pOrF, tt.expectedSort)
		}
	}
}
