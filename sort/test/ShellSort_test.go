package test

import (
	"testing"
	"math/rand"
	"log"
	"DataStructure/sort"
	"fmt"
)

//todo:修改测试 减少代码重复
func TestShellSort(t *testing.T)  {
	K:=rand.Intn(100)
	fmt.Println(K)
	testArr:=make([]int,K)

	for i:=0;i<len(testArr);i++{
		testArr[i]=rand.Intn(100)
	}
	fmt.Println(testArr)
	log.Printf("testArr:%v",testArr)
	log.Printf("testArr after sort:%v",sort.ShellSort(testArr))
}