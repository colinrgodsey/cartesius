package f64

// VecC provides a dimensionless interface
// for all cartesian vecs.
type VecC interface {
	At(dim int) float64
	Get() []float64
	Set(vs ...float64)
	Vec2() Vec2
	Vec3() Vec3
	Vec4() Vec4
	VecN() VecN
}

// VecN is an n-dimentional float64 vector
type VecN []float64

// Vec2 is a 2-dimentional float64 vector
type Vec2 [2]float64

// Vec3 is a 3-dimentional float64 vector
type Vec3 [3]float64

// Vec4 is a 4-dimentional float64 vector
type Vec4 [4]float64

// Cross returns the cross product of v and o (v x o)
func (v Vec3) Cross(b Vec3) Vec3 {
	return Vec3{
		v[1]*b[2] - v[2]*b[1],
		v[2]*b[0] - v[0]*b[2],
		v[0]*b[1] - v[1]*b[0],
	}
}

// VecN version of this vector.
func (v *Vec2) VecN() VecN {
	return NewVecN((*v)[:]...)
}

// VecN version of this vector.
func (v *Vec3) VecN() VecN {
	return NewVecN((*v)[:]...)
}

// VecN version of this vector.
func (v *Vec4) VecN() VecN {
	return NewVecN((*v)[:]...)
}

// VecN version of this vector.
func (v *VecN) VecN() VecN {
	return NewVecN((*v)[:]...)
}

func initVecN(v *VecN, l int) {
	if len(*v) == 0 {
		*v = make([]float64, l)
	}
}

func initVec2(v *Vec2, l int) {}
func initVec3(v *Vec3, l int) {}
func initVec4(v *Vec4, l int) {}
