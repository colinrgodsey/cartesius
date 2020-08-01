/* FILE WAS AUTO-GENERATED FROM f64/base */

package f32

// VecC provides a dimensionless interface
// for all cartesian vecs.
type VecC interface {
	At(dim int) float32
	Get() []float32
	Set(vs ...float32)
	Vec2() Vec2
	Vec3() Vec3
	Vec4() Vec4
	VecN() VecN
}

// VecN is an n-dimentional float32 vector
type VecN []float32

// Vec2 is a 2-dimentional float32 vector
type Vec2 [2]float32

// Vec3 is a 3-dimentional float32 vector
type Vec3 [3]float32

// Vec4 is a 4-dimentional float32 vector
type Vec4 [4]float32

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
		*v = make([]float32, l)
	}
}

func initVec2(v *Vec2, l int) {}
func initVec3(v *Vec3, l int) {}
func initVec4(v *Vec4, l int) {}
