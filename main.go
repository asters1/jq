package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/gjson"
)

func main() {

	jpath := os.Args[1]
	qstr := os.Args[2]

	content, err := ioutil.ReadFile(jpath)
	if err != nil {
		fmt.Println("读取[" + jpath + "]文件失败!请检查!!!")
		os.Exit(1)
	}
	jstr := string(content)
	res := gjson.Get(jstr, qstr)
	fmt.Println(res)
}
