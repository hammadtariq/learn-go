package main

// import "fmt"
// import "math/rand"
// OR

import ("fmt"
		"math/rand")

func printMyRandom(){
	fmt.Println("Random number between 0 - 100 is", rand.Intn(100))
}

func main() {
	printMyRandom()
}

