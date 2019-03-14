package main

import (
	"debug/elf"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	dirpath string
)

func init() {

	flag.StringVar(&dirpath, "d", "", "the file path")

}

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("参数缺失")
		os.Exit(1)
	}

	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		if file.IsDir() == false {

			fmt.Println(file.Name())
			filepath := dirpath + file.Name()

			fs, err := elf.Open(filepath)
			if err != nil {
				log.Fatal(err)
			}

			if fs.FileHeader.Type == 2 {
				fmt.Println(file.Name(), "is ELF")
			}
		}
	}
}
