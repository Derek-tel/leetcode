package quickSort

import "testing"

func TestQuickSort(t *testing.T) {
	t.Log(five([]int{1, 1, 3, 4, 5, -1, -2, 6, 1}))
	t.Error(`quickSort"four" == true`)
}
