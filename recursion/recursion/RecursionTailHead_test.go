package RecursionTailHead

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type FibonacciTest struct {
	input       int
	expectedOpt string
}

func testData() []FibonacciTest {
	return []FibonacciTest{
		{
			3,
			"3 2 1",
		},
		{
			2,
			"3 2 1 2 1",
		},
		{
			5,
			"3 2 1 2 1 5 4 3 2 1",
		},
		{
			7,
			"3 2 1 2 1 5 4 3 2 1 7 6 5 4 3 2 1",
		},
		{
			9,
			"3 2 1 2 1 5 4 3 2 1 7 6 5 4 3 2 1 9 8 7 6 5 4 3 2 1",
		},
		{
			8,
			"3 2 1 2 1 5 4 3 2 1 7 6 5 4 3 2 1 9 8 7 6 5 4 3 2 1 8 7 6 5 4 3 2 1",
		},
	}
}

var (
	impls []RecursionComputer
)

func TestRecursionFibonacci(t *testing.T) {
	impls = append(impls,
		&ImplS1R{},
	)

	tests := testData()
	t.Logf("[#tests=%d]", len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualOpt := impl.recursionTailHead(tt.input)
			var pOrF string
			if assert.Equal(t, tt.expectedOpt, actualOpt) {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test #%d, %s input=[%v] and expected output=%v", i+1, pOrF, tt.input, tt.expectedOpt)
		}
	}
}
