package main

import "fmt"

// FizzBuzzGo:
// An implementation of the Fizz Buzz problem and described at
// The Rosetta Code:> https://rosettacode.org/wiki/FizzBuzz
// Purpose:
// This application is an exercise to get more familiar with the Go
// language and tooling such as Testing, Benchmarking and Documentation.
// ----------------------------------------------------------------------
// Version: V0.1
// 		A first draft based on very simple logic as described in
// 		the spec's
// Flowchart: https://drive.google.com/file/d/1MDRHwi5CPnfGWoJzrTFFhDJKYfb8pM1j/view?usp=sharing
//		All versions will first be designed using flow charts as a design
//		tool to verify design logic and as a roadmap to write the code.
//		While this is a simple application, the tools and methodologies
//		used to write reliable/maintainable software remain pretty much the same.
func main() {
	for x := 1; x <= 100; x++ {
		if x%3 == 0 && x%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if x%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if x%5 == 0 {
			fmt.Println("Buzz")
			continue
		}

		fmt.Println(x)
	}
}
