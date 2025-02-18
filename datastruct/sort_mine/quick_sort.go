package sort_mine

type QuickSort struct{}

func (qs *QuickSort) Sort(arr []int) {
	quickSort(0, len(arr)-1, arr)
}

func quickSort(begin int, end int, arr []int) {
	if begin < end {
		loc := partition(begin, end, arr)

		quickSort(begin, loc-1, arr)
		quickSort(loc+1, end, arr)
	}

}

func partition(begin int, end int, arr []int) int {
	i := begin + 1
	j := end

	for i < j {
		if arr[i] > arr[begin] {
			arr[i], arr[j] = arr[j], arr[i]
			j--
		} else {
			i++
		}
	}

	if arr[i] >= arr[begin] {
		i--
	}
	arr[i], arr[begin] = arr[begin], arr[i]

	return i
}
