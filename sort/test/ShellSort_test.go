package test

import (
	"testing"
	"DataStructure/sort"
)

//todo:修改测试 减少代码重复
func TestShellSort(t *testing.T) {
	getTestResult(sort.ShellSort, "Shell Sort")
}
