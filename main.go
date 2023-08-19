package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var (
	opt         string
	input       string
	key         string
	value       string
	jstr        string
	jstr_output string
)

func PrintHelp() {
	fmt.Println("-------------------帮助信息-------------------")
	fmt.Println("!!!请输入3个参数，否则会触发帮助信息!!!")
	fmt.Println("jq [option] [input] [key]")
	fmt.Println("\n[option]")
	fmt.Println("f-->input为文件")
	fmt.Println("s-->input为字符串")
	fmt.Println("G-->打印key对应值")
	fmt.Println("S-->设置key对应的值，并打印整个json")
	fmt.Println("\n[input]")
	fmt.Println("input为json的来源，可以是文件，也可以是字符串.")
	fmt.Println("\n[key]\nkey为json文件中的key值..\n用法的具体地址为-->https://github.com/tidwall/gjson")
}

func main() {
	if len(os.Args) == 4 || len(os.Args) == 5 {
		opt = os.Args[1]
		input = os.Args[2]
		key = os.Args[3]
		if len(os.Args) == 5 {
			value = os.Args[4]
		}
		if strings.Index(opt, "f") != -1 {
			b_content, err := ioutil.ReadFile(input)
			if err != nil {
				fmt.Println("地址有误！请检查--->" + input)
				os.Exit(1)
			}
			jstr = string(b_content)
		} else if strings.Index(opt, "s") != -1 {
			jstr = input
		} else {
			fmt.Println("你没有选择文件格式，文件或者字符串!!!")
			PrintHelp()
			os.Exit(1)
		}
		if strings.Index(opt, "G") != -1 {
			jstr_output = gjson.Get(jstr, key).String()
		} else if strings.Index(opt, "S") != -1 {
			if len(os.Args) != 5 {
				fmt.Println("你给的参数有误！")
				for i := 0; i < len(os.Args); i++ {
					fmt.Print(os.Args[i] + " ")
				}
				fmt.Println()
				os.Exit(1)
			}
			jstr_output, _ = sjson.Set(jstr, key, value)
		} else {
			fmt.Println("你没有选择执行方式，是Get还是Set!!!")
			PrintHelp()
			os.Exit(1)
		}
	} else {
		PrintHelp()
	}
	fmt.Println(jstr_output)
}
