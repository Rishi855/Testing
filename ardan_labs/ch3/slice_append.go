package main

import "fmt"

func main() {
	slice := make([]string, 0) 
	slice = append(slice, "a")
	slice = append(slice, "b")
	slice = append(slice, "c")
	// slice = append(slice, "d")

	fmt.Println("slice before function: ",slice)
	// manualAppendFuncType1(slice)
	manualAppendFuncType2(slice)
	// manualAppendFuncType3(slice)
	fmt.Println("slice after function: ",slice)
	
}

// func manualAppendFuncType1(slice []string) {
// 	slice[2] = "c1"
// }

func manualAppendFuncType2(slice []string){
	slice[1] = "b1"
	slice = append(slice, "e")
	slice[2] = "c1"
	fmt.Println("slice in function: ",slice)
}

// func manualAppendFuncType3(slice []string){
// 	slice = append(slice, "e")
// }