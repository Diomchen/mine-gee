package sort_mine

func shellSort(arr []int) {
	n := len(arr)
	// 外层：初始化 gap，比如 n/2
	for gap := n / 2; gap > 0; gap /= 2 {
		// 内层：从 i = gap 开始，到 n-1
		for i := gap; i < n; i++ {
			// 保存当前元素，防止被覆盖
			temp := arr[i]
			j := i

			// 关键部分：往左比较，步长为 gap
			// 条件是：j >= gap 且 arr[j - gap] > temp
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}

			// 把 temp 放到正确位置
			arr[j] = temp
		}
	}
}

type ShellStrategy struct{}

func (s *ShellStrategy) Sort(arr []int) {
	shellSort(arr)
}
