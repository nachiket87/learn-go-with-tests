package main

import "fmt"

func Repeat(a string) string {

	var repeated string
	for range 5 {
		repeated += a
	}
	return repeated
}

func main() {
	fmt.Println(Repeat("n"))

}
