package main

import (
	"dataStructures/arrays/arraysADT/difference"
	"fmt"
	"os"
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
	ch, _ := strconv.Atoi(os.Args[1])

	if ch > 0 && ch < 7 {
		switch ch {
		case 1:
			fmt.Printf("Enter the index and variable value: ")
			index, _ := strconv.Atoi(os.Args[2])
			val, _ := strconv.Atoi(os.Args[3])
			arr3.Array.Array.Array.Array.Array.Insert(index, val)
			return arr3.Array.Array.Array.Array.Array.Display()
		case 2:
			fmt.Printf("Enter the index: ")
			index, _ := strconv.Atoi(os.Args[2])
			arr3.Array.Array.Array.Array.Array.Delete(index)
			return arr3.Array.Array.Array.Array.Array.Display()
		case 3:
			fmt.Println("What value are you searching for? ")
			val, _ := strconv.Atoi(os.Args[2])
			return strconv.Itoa(arr3.Array.Array.Array.Array.Array.BinSearch(val))
		case 4:
			return "Sum: " + strconv.Itoa(arr3.Array.Array.Array.Array.Array.Sum())
		case 5:
			return arr3.Array.Array.Array.Array.Array.Display()
		case 6:
			return "Bye!"
		}
	} else {
		return "Incorrect argument passed!"
	}
	return ""
}

func main() {
	var arr3 Array
	arr3.Array.Array.Array.Array.Array.A = [20]int{1, 2, 3, 4, 5}
	arr3.menuProgram()
}
