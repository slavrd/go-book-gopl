//!+

// Package lengthconv converts Feet to Meters and vice versa
package lengthconv

import "fmt"

const (
	// MeterToFeetRatio the ration of meter to feet
	MeterToFeetRatio = 0.3048
)

// Meter represents the meter length unit
type Meter float64

// Foot represents the foot length unit
type Foot float64

func (m Meter) String() string { return fmt.Sprintf("%g m.", m) }
func (f Foot) String() string  { return fmt.Sprintf("%g ft.", f) }

// MToF converts meters to feet units
func MToF(m Meter) Foot {
	return Foot(m * MeterToFeetRatio)
}

// FToM converts feet to meter units
func FToM(f Foot) Meter {
	return Meter(f / MeterToFeetRatio)
}

//!-
