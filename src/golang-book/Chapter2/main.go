package main 
/* This is know as a “package declaration”. 
Every Go program must start with a package declaration. 
Packages are Go's way of organizing and reusing code*/

import "fmt"
/*The import keyword is how we include code from other packages to use 
with our program
The fmt package (shorthand for format) implements formatting
for input and output*/
// this is a comment

func main() {

	fmt.Println(len("Hello World"))
	fmt.Println("hello World"[1])
	fmt.Println("Hello " + "World")
}