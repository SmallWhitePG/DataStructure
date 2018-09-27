package sort

/*
从最后一个数开始构造大顶堆 每次都找出最大的数放在数组最后面
*/
func HeapSort(arr []int)[]int  {
	for i:=len(arr)-1;i>=0;i-- {
		buildBigTopHeap(arr,i)
		arr[0],arr[i]=arr[i],arr[0]
	}
	return arr
}
func buildBigTopHeap(arr []int,n int){
	for i:=(n-1)/2;i>=0;i-- {
		left:=2*i+1
		right:=left+1
		max:=i
		if left<=n&&arr[left]>arr[i] {
			max=left
		}
		if right<=n&&arr[right]>arr[max] {
			max=right
		}
		arr[i],arr[max]=arr[max],arr[i]
	}
}