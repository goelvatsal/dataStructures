package reverse

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
			[20]int{10, 20, 30, 40, 50},
			"Elements are: 50 40 30 20 10 \n",
		},
		{
			[20]int{4, 9, 16, 25, 36},
			"Elements are: 36 25 16 9 4 \n",
		},
		{
			[20]int{5, 10, 15, 20, 25},
			"Elements are: 25 20 15 10 5 \n",
		},
		{
			[20]int{5, 4, 3, 2, 1},
			"Elements are: 1 2 3 4 5 \n",
		},
		{
			[20]int{100, 50, 200, 100, 400},
			"Elements are: 400 100 200 50 100 \n",
		},
	}
}

var (
	impls []ArrayReverser
)

func TestReverse(t *testing.T) {
	impls = append(impls, &ImplS1A{})

	tests := testData()
	t.Logf("[#tests=%d] \n", len(tests))

	assert.NotEmpty(t, impls, "no implementations present!")

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualRev1, actualRev2 := impl.arrayReverser(tt.base)
			var pOrF string
			if assert.Equal(t, tt.expectedArr, actualRev1) && actualRev1 == actualRev2 {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test #%d, %s, expected arr = %v", i+1, pOrF, tt.expectedArr)
		}
	}
}
