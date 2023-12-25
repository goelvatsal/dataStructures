package append

import (
	"dataStructures/arrays/arraysADT/display"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	arr := MyArray{display.Array{A: [20]int{2, 4, 6, 8}, Size: 20, Length: 4}}
	arr.Append(10)
	if assert.Equal(t, "Elements are: 2 4 6 8 10 \n", arr.Display()) {
		t.Fatal("outputs are not equal")
	}
}
