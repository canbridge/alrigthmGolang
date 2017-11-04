// 数组相关的一些算法

package arr

//希尔排序
func shellSort(arr []int) []int {
	length := len(arr)
	gap := 1
	for gap < gap/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
	}
	return arr
}

// 插入排序
func insertionSort(arr []int) []int {
	for i := range arr {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = current
	}
	return arr
}

func swap(slice *[]int, a, b int) {
	(*slice)[a], (*slice)[b] = (*slice)[b], (*slice)[a]
}

func quickSort(slice *[]int, left, right int) {
	if left < right {
		p := quickSortPart(slice, left, right)
		quickSort(slice, left, p)
		quickSort(slice, p+1, right)
	}
}

func quickSortPart(slice *[]int, left, right int) int {
	// 标兵位置, 选左边的
	p := left

	//用来记录没有排好序的右边的第一个, 每次发现一个比标兵小的, 则将新发现的和该位置的元素进行交换
	smallest_big_p := left + 1

	for c := left + 1; c <= right; c++ {
		if (*slice)[c] < (*slice)[p] {
			swap(slice, c, smallest_big_p)
			smallest_big_p++

		}
	}
	swap(slice, smallest_big_p-1, p)
	// 返回左边小于标兵的
	return smallest_big_p - 1
}

//堆排
func heapSort(arr []int) []int {
	arrLen := len(arr)
	buildMaxHeap(arr, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		swap(&arr, 0, i)
		arrLen -= 1
		heapify(arr, 0, arrLen)
	}
	return arr
}

func buildMaxHeap(arr []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

func heapify(arr []int, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(&arr, i, largest)
		heapify(arr, largest, arrLen)
	}
}

// 归并排序
func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		result = append(result, left...)
	}

	if len(right) != 0 {
		result = append(result, right...)
	}

	return result
}
