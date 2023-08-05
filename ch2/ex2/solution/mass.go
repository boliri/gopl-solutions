// Structs with the ability to represent mass values are here

package conv

const poundsUnit = "lb"
const kgUnit     = "kg"

func NewPounds(v float64) Metric { return Metric{v, poundsUnit} }
func NewKilograms(v float64) Metric { return Metric{v, kgUnit} }
