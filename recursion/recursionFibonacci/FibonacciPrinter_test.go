package FibonacciElements

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type FibonacciTest struct {
	input       int
	expectedOpt int
}

func testData() []FibonacciTest {
	return []FibonacciTest{
		{
			3,
			2,
		},
		{
			2,
			1,
		},
		{
			5,
			5,
		},
		{
			7,
			13,
		},
		{
			9,
			34,
		},
		{
			8,
			21,
		},
	}
}

var (
	impls []FibonacciPrinter
)

func TestRecursionFibonacci(t *testing.T) {
	impls = append(impls,
		&S1R{},
	)

	tests := testData()
	t.Logf("[#tests=%d]", len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			fibInit()
			actualOpt := impl.recursionFibonacci(tt.input)
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
