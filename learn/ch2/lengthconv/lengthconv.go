package lengthconv

import "fmt"

type Meter float64
type Feet float64
type Inches float64

const (
	FeetInMeter   Meter  = 3.28
	MetersInFeet  Feet   = 0.3048
	InchesInMeter Inches = 39.37
	InchesInFoot  Feet   = 12
)

// MToF converts Meters to Feet
func MToF(m Meter) Feet { return Feet(m * FeetInMeter) }

// FToM converts Feet to Meters
func FToM(f Feet) Meter { return Meter(f * MetersInFeet) }

// IToM converts Inches to Meters
func IToM(i Inches) Meter { return Meter(i / InchesInMeter) }

// FToI converts Feet to Inches
func FToI(f Feet) Inches { return Inches(f * InchesInFoot) }

// MToI converts Meter to Inches
func MToI(m Meter) Inches { return Inches(m * 39.37) }

// IToF converts Inches to Feet
func IToF(i Inches) Feet { return Feet(i / 12) }

func (m Meter) String() string  { return fmt.Sprintf("%g m", m) }
func (f Feet) String() string   { return fmt.Sprintf("%g feet", f) }
func (i Inches) String() string { return fmt.Sprintf("%g in", i) }
