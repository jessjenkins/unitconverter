package main

type converter struct {
	Name string             `json:"name"`
	Unit map[string]float64 `json:"units"`
}

func (cv *converter) convert(fromUnit, toUnit string, val float64) float64 {
	//TODO error handling for inputs
	fromSize := cv.Unit[fromUnit]
	toSize := cv.Unit[toUnit]
	return fromSize * val / toSize
}
