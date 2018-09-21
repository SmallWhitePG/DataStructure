package sort

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
