package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var newVar int16 = 20
	fmt.Println("Print Main:", newVar, &newVar)

	printReference(&newVar)
	printReference(nil)

	printAsCopy(newVar)

	fmt.Println("Print Main:", newVar, &newVar)
}

func printReference(value *int16) {
	if value != nil {
		*value = 8
		fmt.Println("Print reference:", *value, value)
		return
	}
	fmt.Println("No pointer")
}

func printAsCopy(value int16) {
	value = 4
	fmt.Println("Print as copy:", value, &value, "bytes:", unsafe.Sizeof(value))
}