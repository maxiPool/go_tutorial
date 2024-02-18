package main

import (
	"fmt"
	"rsc.io/quote"
)

// list all dependencies and their versions:
// go list -m -u all
// use go.mod file to set the dependencies for the project; import the required modules your code file like in Java
// go.mod is like the pom.xml

// run `go mod tidy` to download the dependencies
// then run the program with `go run .`
func main() {
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}
