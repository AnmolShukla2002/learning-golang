package main

import (
	"fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("denominator must be a non-zero value")
    }
    return a / b, nil
}

func main() {
    fmt.Println("Error Handling in Go")
    ans, _ := divide(3, 2)         //_ is used to ignore specific value. it is a write-only variable.
    // if err != nil {
    //     fmt.Println("Error:", err)
    // } else {
    //     fmt.Println("Division of 3 and 2 is:", ans)
    // }
		fmt.Println("Division of 3 and 2 is:", ans)
}
