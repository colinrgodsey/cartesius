package f64

// VecS is an n-dimentional float64 vector
// with string-keyed dimensions.
type VecS map[string]float64

// At returns the value at dimension dim.
func (v *VecS) At(dim string) float64 {
	return (*v)[dim]
}

func initVecS(v *VecS, len int) {
	if *v == nil {
		*v = make(map[string]float64)
	}
}
