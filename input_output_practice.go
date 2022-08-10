package main

import("fmt")

func main() {
    var name string
    var age int
    
    fmt.Print("Enter your name: ")
    fmt.Scan(&name)
    fmt.Println("")
    
    fmt.Print("Enter your age: ")
    fmt.Scan(&age)
    fmt.Println("")
    
    fmt.Print(name, age)
}