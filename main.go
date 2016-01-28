package main

import (
    "fmt"
    "strconv"
    "os"
    menu "project-euler-golang/menu"
    problemOne "project-euler-golang/ProblemOne"
    problemTwo "project-euler-golang/ProblemTwo"
)

 
 
func main() {
    m := menu.NewMenu() 
    fmt.Print("\nWelcome to Jonathan Mast's Project Euler in Go!\n\n")
    selectedProblem := parseArgs()
    
    if(selectedProblem > 0) {
        solveProblem(selectedProblem)
    } else {
        for(selectedProblem != 0) {
            selectedProblem = m.PrintMenu()
            solveProblem(selectedProblem)
        }
    } 
    
}

func solveProblem(problemNumber int) {
    switch problemNumber {
        case 0:
            fmt.Println("See ya!")
        case 1:
            fmt.Println("Answer: " + strconv.Itoa( problemOne.Solve()) )
        default:
            fmt.Println("Invalid input.")
            
    }
}

func parseArgs() int {
   
    if(len(os.Args) > 1) {
        arg, _ := strconv.Atoi(os.Args[1])
        return arg
    } 
    return -1
    
}


