package CombinationFormula

import (
	"dataStructures"
	"github.com/stretchr/testify/assert"
	"testing"
)

type nCrTest struct {
	n           int
	r           int
	expectedOpt int
}

func testData() []nCrTest {
	return []nCrTest{
		{
			5,
			1,
			5,
		},
		{
			5,
			2,
			10,
		},
		{
			5,
			3,
			10,
		},
		{
			4,
			2,
			6,
		},
	}
}

var (
	impls []CombinationCalculator
)

func TestRecursionNCR(t *testing.T) {
	impls = append(impls,
		&ImplS1R{})

	tests := testData()
	t.Logf("[#tests=%d]", len(tests))

	assert.NotEmpty(t, impls, "no implementations present!!")

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualOpt := impl.combinationFormula(tt.n, tt.r)
			var pOrF string
			if assert.Equal(t, tt.expectedOpt, actualOpt) {
				pOrF = dataStructures.Passed
			} else {
				pOrF = dataStructures.Failed
			}
			t.Logf("test %d, %s input=[%v, %v] and expected output = %v", i+1, pOrF, tt.n, tt.r, tt.expectedOpt)
		}
	}
}
