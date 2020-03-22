package sort

// 插入排序的实现
func insertSort(slice []int) {
	if slice == nil || len(slice) == 0 {
		return
	}
	size := len(slice)
	// 插入排序的主要逻辑
	for i := 1; i < size; i++ {
		// 若后面的值小于前面的值时，交换位置
		for j := i; j > 0 && slice[j] < slice[j-1]; j-- {
			slice[j], slice[j-1] = slice[j-1], slice[j]
		}
	}

	return
}

// 选择排序的实现
// 选择排序的思想就是找到当前数组中最小的元素，然后将其与第一个元素
// 交换，然后从数组中国的其他元素中找到最小的元素，
func selectSort(slice []int) {
	if slice == nil || len(slice) == 0 {
		return
	}
	size := len(slice)
	// 选择排序的主要逻辑
	for i := 0; i < size; i++ {
		min := i // 记录最小元素的索引
		for j := i + 1; j < size; j++ {
			if slice[j] < slice[min] {
				min = j
			}
		}
		slice[i], slice[min] = slice[min], slice[i]
	}

	return
}

// 希尔排序的排序的实现
// 局部有序 全局有序
func shellSort(slice []int) {
	N := len(slice)
	h := 1

	for h < N/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		// 将数组变成 h 有序
		for i := h; i < N; i++ {
			for j := i; j >= h && slice[j] < slice[j-h]; j -= h {
				slice[j], slice[j-h] = slice[j-h], slice[j]
			}
		}
		h /= 3
	}
}


// 