/*
convweight provides simple utilities for conversion between kilograms and
pounds.
*/

package convweight

import "fmt"

type Kg float64

func (k Kg) String() string {
	return fmt.Sprintf("%.2f kg", k)
}

type Lb float64

func (l Lb) String() string {
	return fmt.Sprintf("%.2f lbs", l)
}

// KgInLb = number of kilograms per pound
const KgInLb = 0.45359237

// KgToLb converts kilograms to pounds
func KgToLb(k Kg) Lb {
	return Lb(k / KgInLb)
}

// LbToKg converts pounds to kilograms
func LbToKg(l Lb) Kg {
	return Kg(l * KgInLb)
}
