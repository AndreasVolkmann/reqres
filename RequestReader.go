package main

import (
	"io/ioutil"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	path := "D:\\Dev\\go\\src\\github.com\\AndreasVolkmann\\reqres\\request.csv"
	dat, err := ioutil.ReadFile(path)
	check(err)
	fmt.Print(string(dat))
}
