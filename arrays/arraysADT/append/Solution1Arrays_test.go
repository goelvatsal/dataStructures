package append

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ArraysTest struct {
	init        [20]int
	new         int
	expectedOut string
}

func testData() []ArraysTest {
	return []ArraysTest{
		{
			[20]int{2, 4, 6, 8},
			10,
			"Elements are: 2 4 6 8 10 \n",
		},
		{
			[20]int{1, 2, 3, 4},
			5,
			"Elements are: 1 2 3 4 5 \n",
		},
		{
			[20]int{10, 20, 30, 40, 50},
			60,
			"Elements are: 10 20 30 40 50 60 \n",
		},
		{
			[20]int{100},
			200,
			"Elements are: 100 200 \n",
		},
	}
}

var (
	impls []Appender
)

func TestArrayAppender(t *testing.T) {
	impls = append(impls,
		&ImplS1A{},
	)

	tests := testData()
	t.Logf("[#tests=%d]\n", len(tests))

	assert.NotEmpty(t, impls, "no implementations present!!")

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualOpt := impl.append(tt.init, tt.new)
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
