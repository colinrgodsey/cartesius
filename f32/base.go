/* FILE WAS AUTO-GENERATED FROM f64/base */

package f32

// Vec provides a dimensionless interface for all vecs.
type Vec interface {
	Get() []float32
	Set(vs ...float32)
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
func (a Vec3) Cross(b Vec3) Vec3 {
	return Vec3{
		a[1]*b[2] - a[2]*b[1],
		a[2]*b[0] - a[0]*b[2],
		a[0]*b[1] - a[1]*b[0],
	}
}

func initVecN(v *VecN, len int) {
	if *v == nil {
		*v = make([]float32, len)
	}
}

func initVec2(v *Vec2, len int) {}
func initVec3(v *Vec3, len int) {}
func initVec4(v *Vec4, len int) {}
