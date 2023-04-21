package test

import (
	"fmt"
	_ "github.com/one-meta/meta/app/ent/runtime"
	"os"
	"testing"
)

//	func TestHello(t *testing.T) {
//		t.Log("TestHello")
//	}
//
//	func cleanuptest() {
//		fmt.Println("Cleanup.😀")
//	}
//
//	func TestClean(t *testing.T) {
//		t.Cleanup(cleanuptest)
//		fmt.Println("testing done.")
//	}
func init() {
	fmt.Println("init.")
}

func TestMain(m *testing.M) {
	// 在开始运行单元测试代码之前
	// 可以在此处添加环境初始化相关代码或者函数调用
	fmt.Println("初始化fiber")

	fiberApp, f = initFiber(false)
	defer f()

	retCode := m.Run()

	// 在全部测试代码运行结束退出之前
	// 可以在此处添加清理代码或函数调用
	fmt.Println("test done.")

	os.Exit(retCode)
}
