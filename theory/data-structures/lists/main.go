package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List

	intList.PushBack(14)
	intList.PushBack(45)
	intList.PushBack(98)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
