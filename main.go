//Declare a main package
/* A package is a way to group functions,
and it's made up of all the files in the same directory).*/
package main

/*Import the popular fmt package, which contains functions for formatting text,
including printing to the console. It will fail if the package it's not used*/

//This package is one of the standard library packages you got when you installed Go.
import (
	"fmt"
	"runtime"
)

/*
Implement a main function to print a message to the console.
A main function executes by default when you run the main package.
*/
func main() {
	fmt.Println("Hello")

	//Declaring the os variable (using a compiler automatically type determiner (:=))

	/*We store inside the os variable the scanner calling the GOOS method.
	This method Go Operating System), it is an environment variable and a standard
	library constant (runtime.GOOS) used to identify or target specific operating
	systems like linux, windows, or darwin (macOS).*/
	os := runtime.GOOS

	if os == "linux" {
		//Printing the os variable. We can concatenate using a coma (,) and the value
		fmt.Println("Your Operative System is: ", os)
	} else {
		fmt.Println("Operative SYstem not recognized")
	}
}
