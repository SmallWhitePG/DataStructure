package sort

/*
1.选择一个增量序列 t1，t2，……，tk，其中 ti > tj, tk = 1；

2.按增量序列个数 k，对序列进行 k 趟排序；

3.每趟排序，根据对应的增量 ti，将待排序列分割成若干长度为 m 的子序列，分别对各子表进行直接插入排序。仅增量因子为 1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。
*/

//todo:解决希尔排序的错误
func ShellSort(arr []int) []int {
	length := len(arr)
	gap := int(length/2) //gap从数组长度的1/2开始
	for gap !=0 {
		for i := gap; i < length; i++ { //每次排序开始是从arr[gap]开始的
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp { //然后对前面相隔gap的数依次进行插入排序
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap=int(gap/2)
	}
	return arr
}
