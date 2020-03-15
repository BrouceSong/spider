package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	url string = "https://www.17k.com/all/book/2_14_128______1.html"
)

func main() {
	fmt.Println("开始抓取数据")
	data, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer data.Body.Close()
	content, err := ioutil.ReadAll(data.Body)
	fmt.Println(string(content))
}
