package TaylorSeriesContinued

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TaylorSeriesHornersRuleTest struct {
	input       int
	precision   int
	expectedOpt float64
}

func testData() []TaylorSeriesHornersRuleTest {
	return []TaylorSeriesHornersRuleTest{
		{
			2,
			25,
			7.38905609893065,
		},
		{
			5,
			5,
			65.375,
		},
		{
			10,
			2,
			11,
		},
		{
			7,
			13,
			1067.0243116202532,
		},
		{
			9,
			7,
			1675.5625,
		},
		{
			3,
			14,
			20.085468593906093,
		},
	}
}

var (
	impls []TaylorSeriesHornerCalculator
)

func TestTaylorSeriesHornersRule(t *testing.T) {
	impls = append(impls,
		&S1R{},
	)

	tests := testData()
	t.Logf("[#tests=%d]", len(tests))

	assert.NotEmptyf(t, impls, "no implemantations present!!")

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		for i, tt := range tests {
			actualOpt := impl.recursionHornersRule(float64(tt.input), float64(tt.precision), 0)
			var pOrF string
			if assert.Equal(t, tt.expectedOpt, actualOpt) {
				//pOrF = dataStructures.Passed
			} else {
				//pOrF = dataStructures.Failed
			}
			t.Logf("test #%d, %s input=[%v, %v] and expected output = %v", i+1, pOrF, tt.input, tt.precision, tt.expectedOpt)
		}
	}
}
