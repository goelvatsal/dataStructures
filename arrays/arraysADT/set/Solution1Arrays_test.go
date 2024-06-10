package set

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	base        [20]int
	index       int
	newVal      int
	expectedOut string
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{10, 20, 30, 40, 60},
			5,
			50,
			"Elements are: 10 20 30 40 50 \n",
		},
		{
			[20]int{5, 10, 15, 20, 30, 30},
			5,
			25,
			"Elements are: 5 10 15 20 25 30 \n",
		},
		{
			[20]int{100, 200, 300},
			5,
			100,
			"Elements are: 100 200 300 \n",
		},
		{
			[20]int{50, 100, 250, 500, 100},
			5,
			1000,
			"Elements are: 50 100 250 500 1000 \n",
		},
	}
}

var (
	impls []Setter
)

func TestSet(t *testing.T) {
	impls = append(impls,
		&ImplS1A{})

	tests := testData()
	t.Logf("[#tests=%d] \n", len(tests))

	assert.NotEmpty(t, impls, "no implementations present!!")

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualOpt := impl.set(tt.base, tt.index, tt.newVal)
			var pOrF string
			if assert.Equal(t, tt.expectedOut, actualOpt) {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test #%d, %s, expected output=%v", i+1, pOrF, tt.expectedOut)
		}
	}
}
