package merge_sorted_array

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	input := []int{4, 5, 6, 7, 0, 1, 2}
	input2 := []int{1, 3, 9, 7, 0, 1, 2}
	k := 7
	res := merge(input, k, input2, k)
	fmt.Println(res)
}
