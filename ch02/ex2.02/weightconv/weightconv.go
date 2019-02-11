//!+

// Package weightconv performs pounds and kilograms conversions.
package weightconv

import "fmt"

const (
	KiloToPoundRatio = 0.45359237
)

type Kilogram float64
type Pound float64

func (k Kilogram) String() string { return fmt.Sprintf("%g kg.", k) }

// KToP converts Kilograms to Pounds
func KToP(k Kilogram) Pound {
	return Pound(k / KiloToPoundRatio)
}

func (p Pound) String() string { return fmt.Sprintf("%g lb.", p) }

// PToK converts Pounds to Kilograms
func PToK(p Pound) Kilogram {
	return Kilogram(p * KiloToPoundRatio)
}

//!-
