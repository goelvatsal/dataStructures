package main

import (
	"dataStructures/arrays/arraysADT/difference"
	"fmt"
	"strconv"
)

type Array struct {
	difference.Array
}

func (arr3 Array) menuProgram() string {
	fmt.Println("What operation do you want to do?")
	fmt.Println("1.  Insert")
	fmt.Println("2.  Delete")
	fmt.Println("3.  Search")
	fmt.Println("4.  Sum")
	fmt.Println("5.  Display")
	fmt.Println("6.  Exit")

	fmt.Printf("Enter choice: ")
	var ch int
	fmt.Scan(&ch)

	if ch > 0 && ch < 7 {
		var index, val int
		switch ch {
		case 1:
			fmt.Printf("Enter the index value: ")
			fmt.Scan(&index)
			fmt.Printf("Enter the variable value: ")
			fmt.Scan(&val)
			arr3.Array.Array.Array.Array.Array.Insert(index, val)
			return arr3.Array.Array.Array.Array.Array.Display()
		case 2:
			fmt.Printf("Enter the index: ")
			fmt.Scan(&index)
			arr3.Array.Array.Array.Array.Array.Delete(index)
			return arr3.Array.Array.Array.Array.Array.Display()
		case 3:
			fmt.Printf("What value are you searching for? ")
			fmt.Scan(&val)
			fmt.Println("Value found! Location is:", strconv.Itoa(arr3.Array.Array.Array.Array.Array.BinSearch(val)))
			return ""
		case 4:
			fmt.Println("Sum: " + strconv.Itoa(arr3.Array.Array.Array.Array.Array.Sum()))
			return ""
		case 5:
			return arr3.Array.Array.Array.Array.Array.Display()
		case 6:
			fmt.Println("Bye!")
		}
	} else {
		fmt.Println("Incorrect argument passed!")
		return ""
	}
	return ""
}

func main() {
	var arr3 Array
	arr3.Array.Array.Array.Array.Array.A = [20]int{1, 2, 3, 4, 5}
	arr3.Array.Array.Array.Array.Array.Length = 5
	arr3.menuProgram()
}
