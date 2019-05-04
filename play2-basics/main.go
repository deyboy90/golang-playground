// Every go file needs to belong to a package
// Packages are of two types, Executable & Reusable
// "main" is the reserved keyword which implies that
// the current file is a Executatable. Any other name when
// used makes it a Reusable file (good for dependencies)
package main

import "fmt"

func main() {
	fmt.Println("Welcome to golang")
}
