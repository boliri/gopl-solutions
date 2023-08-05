// Structs with the ability to represent temperatures are here

package conv

const cUnit = "°C"
const fUnit = "°F"
const kUnit = "K"

func NewCelsius(v float64) Metric { return Metric{v, cUnit} }
func NewFahrenheit(v float64) Metric { return Metric{v, fUnit} }
func NewKelvin(v float64) Metric { return Metric{v, kUnit} }