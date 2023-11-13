package main

import (
	"io"
	"log"
	"os"
)

func main() {
	//从1.txt中获取数据流
	from, err := os.Open("D:\\GOPROject\\src\\tt\\21\\1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	//创建文件4.txt
	to, err := os.OpenFile("D:\\GOPROject\\src\\tt\\21\\5.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	//从1.txt复制数据流去5.txt
	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}

}
