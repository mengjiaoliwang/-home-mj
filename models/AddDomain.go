package main

import (
	"fmt"
	"os"
)

func AddDomain2Url(url string) bool {

	fin, err := os.Open("../conf/aap.conf")
	if err != nil {
		fmt.Println("打开文件失败")
		return false
	}
	defer fin.Close()
	return true

}
