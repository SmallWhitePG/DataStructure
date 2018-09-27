package sort

/*
空间换时间
给一个长度为最大值+1的数组
然后扫描要排序的数组在 在新数组中 要排序的值的位置上+1
类似位图
*/
func CountingSort(arr []int)[]int  {
	max:=arr[0]
	for _,v:=range arr{
		if v>max {
			max=v
		}
	}
	return countingSort(arr,max)
}
func countingSort(arr []int, maxValue int) []int {
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen) // 初始为0的数组

	sortedIndex := 0
	length := len(arr)

	for i := 0; i < length; i++ {
		bucket[arr[i]] += 1
	}

	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			arr[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}

	return arr
}
