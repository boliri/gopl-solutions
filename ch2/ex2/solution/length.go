// Structs with the ability to represent length values are here

package conv

const feetUnit = "ft"
const mUnit    = "m"

func NewFeets(v float64) Metric { return Metric{v, feetUnit} }
func NewMeters(v float64) Metric { return Metric{v, mUnit} }
