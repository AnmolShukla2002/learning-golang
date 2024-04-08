package main

import "fmt"

func simpleFunction(){
	fmt.Println("Simple function")
}

func add(a,b int)(result int){       //the second int is return type of the function
	result= a+b
	return
}

func main() {
	fmt.Println("Learning functions")
	simpleFunction()
	fmt.Println("Sum of 4 and 5 is:",add(4,5))
}