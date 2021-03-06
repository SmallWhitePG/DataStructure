package sort

/*
把数组分为两部分对两部分进行分别排序
再分别对两部分
每部分分两组排序
最后合并
*/
func MergeSort(arr []int)[]int  {
	length:=len(arr)
	if length<2 {
		return arr
	}
	middle:=length/2
	left:=arr[0:middle]
	right:=arr[middle:]
	return merge(MergeSort(left),MergeSort(right))
}

func merge(left []int,right []int)[]int  {
	var result []int
	for len(left)!=0&&len(right)!=0{
		if left[0]<right[0] {
			result=append(result,left[0])
			left=left[1:]
		}else {
			result=append(result,right[0])
			right=right[1:]
		}
	}
	for len(left)!=0 {
		result=append(result,left[0])
		left=left[1:]
	}
	for len(right)!=0 {
		result=append(result,right[0])
		right=right[1:]
	}
	return result
}
