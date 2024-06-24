package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//读取文件
	open, err := os.Open("/Users/zhihu/Downloads/archive_blade_18.lab")
	if err != nil {
		log.Printf("read lab err:%+v", err)
		return
	}
	defer open.Close()
	//声明一个切片
	liu := make([]byte, 128)
	content := make([]byte, 0)
	for {
		n, err := open.Read(liu)
		if err == io.EOF {
			log.Printf("文件读完了")
			break
		}
		if err != nil {
			log.Printf("read file failed, err:%+v", err)
			return
		}
		content = append(content, liu[:n]...)
	}
	fmt.Println(content)
}
