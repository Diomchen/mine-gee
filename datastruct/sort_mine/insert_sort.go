package sort_mine

type InsertStrategy struct{}

func (ins *InsertStrategy) Sort(arr []int) {
	insertSort(arr)
}

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		insValue := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > insValue {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = insValue
	}
}
