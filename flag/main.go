package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	name string
	age  int
	sex  bool
)

func init() {
	flag.StringVar(&name, "n", "", "the name string")
	flag.IntVar(&age, "a", 17, "the defaule age is 17")
	flag.BoolVar(&sex, "s", true, "Male or Famale")
}

func main() {

	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Println("参数缺失")
		os.Exit(1)
	}

	if name == "" {
		fmt.Println("Name string missing")
		os.Exit(1)
	}
	fmt.Println(name)

	if age > 17 {
		fmt.Println("成年人")
	} else {
		fmt.Println("未成年人")
	}
	if sex == true {
		fmt.Println("男性")
	} else {
		fmt.Println("女性")
	}

}
