package f64

// Vec provides a dimensionless interface for all vecs.
type Vec interface {
	Get() []float64
	Set(vs ...float64)
}

// VecN is an n-dimentional float64 vector
type VecN []float64

// Vec2 is a 2-dimentional float64 vector
type Vec2 [2]float64

// Vec3 is a 3-dimentional float64 vector
type Vec3 [3]float64

// Vec4 is a 4-dimentional float64 vector
type Vec4 [4]float64

func initVecN(v *VecN, len int) {
	if *v == nil {
		*v = make([]float64, len)
	}
}

func initVec2(v *Vec2, len int) {}
func initVec3(v *Vec3, len int) {}
func initVec4(v *Vec4, len int) {}
