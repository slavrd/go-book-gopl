//!+

// converter takes numerical arguments and displays them as different units for length, temp and weight
package main

import (
	"fmt"
	"os"
	"strconv"

	tempconv "github.com/go-book-gopl/ch02/ex2.01"
	"github.com/go-book-gopl/ch02/ex2.02/lengthconv"
	"github.com/go-book-gopl/ch02/ex2.02/weightconv"
)

func main() {

	// get cmd arguments
	for _, arg := range os.Args[1:] {
		a, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "converter: %v\n", err)
			os.Exit(1)
		}

		// print units as length
		fmt.Printf("Length:\n")
		feet := lengthconv.Foot(a)
		meters := lengthconv.Meter(a)
		fmt.Printf("%s == %s\n%s == %s\n\n", feet, lengthconv.FToM(feet), meters, lengthconv.MToF(meters))

		// print units as weight
		fmt.Printf("Weight:\n")
		pound := weightconv.Pound(a)
		kilo := weightconv.Kilogram(a)
		fmt.Printf("%s == %s\n%s == %s\n\n", pound, weightconv.PToK(pound), kilo, weightconv.KToP(kilo))

		// print units as temperature
		fmt.Printf("Temp:\n")
		celsius := tempconv.Celsius(a)
		fahrenheit := tempconv.Fahrenheit(a)
		fmt.Printf("%s == %s\n%s == %s\n\n", celsius, tempconv.CToF(celsius), fahrenheit, tempconv.FToC(fahrenheit))
	}
}

//!-
