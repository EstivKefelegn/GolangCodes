package main

import "fmt"

func main() {
	// This is a placeholder for the main function.
	// You can add your application logic here
	var numbers [5]int
	var stringValues [5]string
	fmt.Println("Printing the 5 number spaces here:", numbers)
	fmt.Println("Printing the 5 string spaces here:", stringValues)

	numbers[0] = 200
	fmt.Println("Printing the 5 number spaces here:", numbers)
	stringValues[3] = "Estifanos"
	fmt.Println("Printing the 5 string spaces here:", stringValues)

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("The inedex of the current number is %d and the value is %d\n", i, numbers[i])
	}

	for index, value := range stringValues {
		if value != "" {
			fmt.Printf("Thhe current index is %d and the value is %s\n", index, value)
		}

		if value == "" {
			stringValues[index] = "_"
			fmt.Printf("Thhe current index is %d and the value is %s\n", index, value)
		}
	}

	var multiDeimeArray [3][3]int
	multiDeimeArray[0][0] = 1
	multiDeimeArray[0][1] = 2
	multiDeimeArray[0][2] = 3

	multiDeimeArray[1][0] = 4
	multiDeimeArray[1][1] = 5
	multiDeimeArray[1][2] = 6

	multiDeimeArray[2][0] = 7
	multiDeimeArray[2][1] = 8
	multiDeimeArray[2][2] = 9

	fmt.Println("Printign the multi-dimensional array:", multiDeimeArray[0])

	fmt.Println("Final string values:", stringValues)
	num1, _ := newNumberGen()
	fmt.Printf("The two numbers generated is %d\n", num1)

	newArrays := [3]int{1, 2, 3}
	comArrays := [3]int{0, 1, 3}

	newArrCpy := newArrays
	fmt.Println("New arrays:", newArrays)
	fmt.Println("Copied arrays:", newArrCpy)
	fmt.Println("New arrays:", newArrays[len(newArrays)-1] == comArrays[len(comArrays)-1])

	numPointers := newNumberGen2()
	fmt.Println("The number pointers are:", numPointers)


	newSlice := make([]int, 5)
	fmt.Println("The length of the slice is:", len(newSlice))
	newSlice[0] = 10
	newSlice[1] = 30
	newSlice[2] = 50
	newSlice[3] = 70
	newSlice[4] = 90
	newSlice = append(newSlice, 100)
	fmt.Println("The length of the slice is:", len(newSlice))
	fmt.Println("The new slice is:", newSlice)


	var nilSlice []int

	fmt.Println("The length of the nil slice is:", len(nilSlice))
}

func newNumberGen() (int, int) {
	return 1, 2
}


func newNumberGen2() *[5]int {
	var numberLists *[5]int

	numsLists := [5]int{1,2,3,4,5}
	numberLists = &numsLists

	fmt.Println("Printing the number lists:", *numberLists)
	return numberLists
}



/*
Stopped on the slices package and its methods.
*/