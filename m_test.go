package main

import (
	"fmt"
	"testing"

	"github.com/f1748x/tw-sdk/sweetsdk/comm"
)

func TestMain(t *testing.T) {
	fmt.Println("执行----------------main")
	sdk := NewClient(&comm.Com{
		AppId: "1111111",
		Secrt: "222222",
	})
	fmt.Println(sdk.Authentication())
}
func TestMain2(t *testing.T) {
	fmt.Println("执行----------------main2")
	sdk := NewClient(&comm.Com{
		AppId: "1111111",
		Secrt: "222222",
	})
	fmt.Println(sdk.GetPclas())
}
