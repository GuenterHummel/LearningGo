// My weight loss program
package main

import "fmt"

// main is the function where all begins
func main() {
	fmt.Printf("My weight on the surface of Mars is %v kg,", 105*0.3738)
	fmt.Printf(" and I would be %v years old.\n", 52*365/687)

	fmt.Printf("My weight on the surface of %v is %v kg\n", "Earth", 100)

	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
