package insertSort

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	base        [20]int
	x           int
	expectedArr string
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{10, 20, 30, 50},
			40,
			"Elements are: 10 20 30 40 50 \n",
		},
		{
			[20]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			1,
			"Cannot insert element!",
		},
		{
			[20]int{4, 9, 16, 36},
			25,
			"Elements are: 4 9 16 25 36 \n",
		},
		{
			[20]int{45, 90, 135, 180},
			225,
			"Elements are: 45 90 135 180 225 \n",
		},
		{
			[20]int{-100, -20, 50, 75},
			-85,
			"Elements are: -100 -85 -20 50 75 \n",
		},
	}
}

var (
	impls []ArraysInsertSorter
)

func TestInsertSorter(t *testing.T) {
	impls = append(impls, &ImplS1A{})

	tests := testData()
	t.Logf("[#tests=%d] \n", len(tests))

	assert.NotEmpty(t, impls, "no implementations present!")

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualArr := impl.arraysInsertSorter(tt.base, tt.x)
			var pOrF string
			if assert.Equal(t, tt.expectedArr, actualArr) {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test #%d, %s, expected array = %v", i+1, pOrF, tt.expectedArr)
		}
	}
}
