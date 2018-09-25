package test

import (
	"math/rand"
	"time"
	"fmt"
)

func init()  {
	rand.Seed(time.Now().Unix())
}

func getTestResult(sortFunc func(arr []int)[]int,sortFuncName string)  {
	testArr:=make([]int,rand.Intn(20))

	for i:=0;i<len(testArr);i++{
		testArr[i]=rand.Intn(100)
	}

	fmt.Printf("testArr:%v\n",testArr)

	sortResult:=sortFunc(testArr)

	fmt.Printf("testArr after sort:%v\n",sortResult)
	for i:=1;i<len(sortResult);i++{
		if sortResult[i-1]>sortResult[i] {
			fmt.Printf("%v Error!!!\n\n",sortFuncName)
			return
		}
	}
	fmt.Printf("%v Right Sort\n\n",sortFuncName)
}