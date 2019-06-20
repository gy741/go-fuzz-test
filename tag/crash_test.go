package main

import (
	"fmt"
	"os"

	"github.com/dhowden/tag"
)

func main() {
	f, err := os.Open("crashers\\f50dec762d42780c3a18cafc1a07934c67cca582")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = tag.ReadFLACTags(f)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("no err")
}