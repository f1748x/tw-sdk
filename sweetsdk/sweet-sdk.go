package sweetSdk

// package main

import (
	"fmt"

	"github.com/f1748x/tw-sdk/sweetsdk/comm"
)

// var ComManager = &Com{}
// 获取渠道平台

func NewSweetClient(cm *comm.Com) comm.ComInterface {
	// 1.检查appid
	fmt.Println("test:", cm.AppId)
	return cm
}

// github.com/f1748x/tw-sdk
