package main

import (
	"fmt"
	"mylearning/myutil"
)

func main() {
	fmt.Println("Learning Go Language By Anmol Shukla")
  
	// data:="Anmol"
	// fmt.Println(data)

	myutil.Myutil("Hello My Util")

	var name string= "Anmol"
	fmt.Println(name)

	var age= 21    // can be anything if data type is not specified
	fmt.Println(age)

	var salary int= 50000
	fmt.Println(salary)


	 Balance:=200000         // CAN BE EXPORTED SINCE FIRST LETTER IS CAPITAL
	 fmt.Println(Balance)    

	 amount:=200             // CANNOT BE EXPORTED SINCE FIRST LETTER IS SMALL
	 fmt.Println(amount)

	 //SAME THING GOES WITH THE FUNCTIONS AS WELL.
}
