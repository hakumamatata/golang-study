// Package gosayhoho `go doc` can see desc for the package
package gosayhoho

import "fmt"

const SayWord string = "hoho"

// Say `go doc` can see desc for the function
func Say() {
	fmt.Println(SayWord)
}
