package problems

import (
    "fmt"
)

//Solve solves the problem number one
func Solve() int {
    fmt.Println("\nSolving Problem Number: 1")
    sum := 0
    for i := 1; i < 1000; i++ {
        if (i % 3 == 0 || i % 5 == 0) {
            sum += i
        }
    }
    return sum
}