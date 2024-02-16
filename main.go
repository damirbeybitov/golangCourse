package main

import "fmt"

func main() {
	slice := []int{1,534,3,4,15,51,543,2}
	fmt.Println("task 2.1:", SortSlice(slice))
	fmt.Println("task 2.2:", IncrementOdd(slice))
	fmt.Print("task 2.3: ")
	PrintSlice(IncrementOdd(slice))
	fmt.Println("task 2.4: ", ReverseSlice(slice))
	fmt.Print("task 3.1: ")
	AppendFunc(PrintSlice, IncrementOdd1)(slice)
	fmt.Print("check: ")
	AppendFunc(ReverseSlice1, IncrementOdd1, PrintSlice)(slice)
}

func SortSlice(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}

func IncrementOdd(slice []int) []int{
    for i := 1; i < len(slice); i += 2 {
        slice[i]++
    }
	return slice
}

func PrintSlice(slice []int) {
	for _, value := range slice {
		fmt.Print(value, " ")
	}
	fmt.Println()
}

func ReverseSlice(slice []int) []int{
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func AppendFunc(dst func([]int), src ...func([]int)) func([]int) {
	return func(slice []int) {
		dst(slice)
		for _, fn := range src {
			fn(slice)
		}
	}
}

func IncrementOdd1(slice []int){
    for i := 1; i < len(slice); i += 2 {
        slice[i]++
    }
}

func ReverseSlice1(slice []int){
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}