package main

import "fmt"

func main() {
    var str = "This is a long string that contains more than sixty-four characters in total, and it keeps going until it exceeds that length. Helllo"
 
	fmt.Printf("size of str %d\n",len(str))
	fmt.Printf("str: %v\n", str)
}
