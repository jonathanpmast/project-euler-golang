package menu

import (
    "fmt"
    "strconv"
)



//EulerMenu Struct for creating a menu and handling input
type EulerMenu struct{
    Options []Option
}

//Option an individual menu option for the menu
type Option struct{
    ProblemNumber int
    Title string
    Description string
}

//NewMenu create an new EulerMenu with the correct options
func NewMenu() *EulerMenu {
    em := &EulerMenu{ 
        Options: []Option{
            Option{0, "Exit Program", ""},
            Option{1, "Multiples of 3 and 5", ""},
            Option{2, "Even Fibonnaci Numbers",""},
        },
    }
    return em
}

//PrintMenu Prints the menu
func (m *EulerMenu) PrintMenu() int  {
    var input int;
    
    m.printOptions()
    fmt.Print("\nPlease select a problem: ")
    if _,err := fmt.Scan(&input); err != nil {
        fmt.Println("something broke: " + err.Error())
    }
    return input
}

//ToString converts the menu option ToString 
func (mo *Option) ToString() string {
    return strconv.Itoa(mo.ProblemNumber) + ". " + mo.Title
}

func (m *EulerMenu) printOptions() {    
    for i := 0; i < len(m.Options); i++ {
        fmt.Println(m.Options[i].ToString())
    }
}

