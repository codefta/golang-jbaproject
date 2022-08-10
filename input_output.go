package main

import("fmt")

func main() {
    fmt.Print("Hello, World!")
    fmt.Print("Hello, World!\n")
    fmt.Print("This is my first Go program!\n")
    fmt.Println(true)
    fmt.Println(1023493)
    fmt.Println("my string")
    fmt.Println(12.4)
    fmt.Println('A')
    
    var boolean = true
    var integer = 1023493
    var str = "my string"
    var float = 12.4
    var character = 'A'
    fmt.Println(boolean, integer, str, float, character)
    
    var name string
    fmt.Scan(&name)
    fmt.Println(name)
    
    var foo int
    var str2 string
    fmt.Scan(&foo)
    
    fmt.Scan(&str2)
    
    var name2 string
    var age string
    
    fmt.Scan(&name2, &age)
    fmt.Println(name2)
    fmt.Println(age)
    
    
}