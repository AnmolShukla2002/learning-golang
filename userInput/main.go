package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey, What's Your Name?")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hey There, My Name is", name)

	reader:=bufio.NewReader(os.Stdin)
	name,_:= reader.ReadString('\n')
	fmt.Println("Hey There, My Name is", name)

}