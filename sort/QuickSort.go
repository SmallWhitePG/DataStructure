package sort

/*
以第一个数为基准对数组进行划分 小的在右边 大的在左边
然后把该基准数两边 当成两个数组以此迭代进行排序
*/
func QuickSort(arr []int) []int {
	_quickSort(arr, 0, len(arr)-1)
	return arr
}

func _quickSort(arr []int, p, q int) {
	left := p
	right := q
	position := p
	if left>=right||left<0||right>len(arr)-1 {
		return
	}
	for left < right {
		if position==left {
			if arr[left]>arr[right] {
				arr[left],arr[right]=arr[right],arr[left]
				position=right
				left++
			}else{
				right--
			}
		}
		if position==right {
			if arr[left]>arr[right] {
				arr[left],arr[right]=arr[right],arr[left]
				position=left
				right--
			}else{
				left++
			}
		}
	}
	_quickSort(arr, p, position-1)
	_quickSort(arr, position+1, q)
}
