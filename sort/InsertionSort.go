package sort

/*
从第二个数开始 一个一个向前找 直到找到比小的数 就在这个位置插入
*/
func InsertionSort(arr []int)[]int  {
	for i:=range arr{
		preIndex:=i-1
		current:=arr[i]
		for preIndex>=00&&arr[preIndex]>current {
			arr[preIndex+1]=arr[preIndex]
			preIndex-=1
		}
		arr[preIndex+1]=current
	}
	return arr
}
