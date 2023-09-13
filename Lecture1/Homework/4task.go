package Homework

type Boxes struct {
	array []int
}

func NewBoxes(a []int) *Boxes {
	return &Boxes{array: a}
}

func swap(a []int, pivotIndex, startIndex int) {
	temp := a[pivotIndex]
	a[pivotIndex] = a[startIndex]
	a[startIndex] = temp
}

// TC: O(n*logn) worst-case:O(n^2)
func partition(a []int, start, end int) int {
	var (
		pivot       = a[start]
		i, j, count = start, end, 0
	)

	for i := start + 1; i <= end; i++ {
		if a[i] <= pivot {
			count++
		}
	}
	pivotIndex := start + count
	swap(a, pivotIndex, start)
	for i < pivotIndex && j > pivotIndex {
		for a[i] <= pivot {
			i++
		}
		for a[j] > pivot {
			j--
		}
		if i < pivotIndex && j > pivotIndex {
			swap(a, i, j)
			i++
			j--
		}
	}

	return pivotIndex
}

func (b *Boxes) QuickSort() {
	start, end := 0, len(b.array)-1
	b.quickSortHelper(start, end)
}

func (b *Boxes) quickSortHelper(start, end int) {
	if start >= end {
		return
	}

	p := partition(b.array, start, end)

	b.quickSortHelper(start, p-1)
	b.quickSortHelper(p+1, end)
}
