package main

import "fmt"

func main() {
	fmt.Println("Learning Arrays")

	var name[5] string
	name[0]="Anmol"
	name[2]="Yashi"

	fmt.Println("Name of person is:", name)
	fmt.Println("length of array is:", len(name))
}