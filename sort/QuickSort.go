package sort

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
