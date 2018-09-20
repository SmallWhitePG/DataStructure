package test

import (
	"math/rand"
	"log"
	"time"
)

func init()  {
	rand.Seed(time.Now().Unix())
}

func getTestResult(sortFunc func(arr []int)[]int,sortFuncName string)  {
	testArr:=make([]int,rand.Intn(20))

	for i:=0;i<len(testArr);i++{
		testArr[i]=rand.Intn(100)
	}
	log.Printf("testArr:%v",testArr)
	log.Printf("testArr after sort:%v",sortFunc(testArr))
	for i:=1;i<len(testArr);i++{
		if testArr[i-1]>testArr[i] {
			log.Printf("%v Error!!!",sortFuncName)
			return
		}
	}
	log.Printf("%v Right Sort",sortFuncName)
}