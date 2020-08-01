/* FILE WAS AUTO-GENERATED FROM f64/vecN */

package f64

// Mul scales v by s
func (v VecS) Mul(s float64) (res VecS) {
	initVecS(&res, len(v))
	for i := range v {
		res[i] = v[i] * s
	}
	return
}

// Abs returns the absolute value of v
func (v VecS) Abs() (res VecS) {
	initVecS(&res, len(v))
	for i := range v {
		res[i] = absfloat64(v[i])
	}
	return
}

// Sub o from v
func (v VecS) Sub(o VecS) VecS {
	return v.Add(o.Neg())
}

// Div scales v by the multiplicative inverse of s
func (v VecS) Div(s float64) VecS {
	return v.Mul(1.0 / s)
}

// Neg returns the negative vector of v (-v)
func (v VecS) Neg() VecS {
	return v.Mul(-1.0)
}

// Inv returns the multiplicative inverse of v
func (v VecS) Inv() (res VecS) {
	initVecS(&res, len(v))
	for i := range v {
		res[i] = 1.0 / v[i]
	}
	return
}

// Add o to v
func (v VecS) Add(o VecS) (res VecS) {
	initVecS(&res, len(v))
	//TODO: this doesnt technically work for different size vecNs, or vecS
	for i := range v {
		res[i] = v[i] + o.At(i)
	}
	return
}

// MulV multiplies v*s per dim
func (v VecS) MulV(o VecS) (res VecS) {
	initVecS(&res, len(v))
	for i := range v {
		res[i] = v[i] * o.At(i)
	}
	return
}

// Dot returns the dot product of v and o (v⋅o)
func (v VecS) Dot(o VecS) (d float64) {
	for i := range v {
		d += v[i] * o.At(i)
	}
	return
}

// Within returns true if v is within the bounds of o
// considering both values as their absolute.
func (v VecS) Within(o VecS) bool {
	v = v.Abs()
	o = o.Abs()
	for i := range v {
		if v[i] > o.At(i) {
			return false
		}
	}
	return true
}

// Eq returns true if v and o are equal
func (v VecS) Eq(o VecS) bool {
	for i := range v {
		if v[i] != o[i] {
			return false
		}
	}
	return true
}

// Mag returns the L2 norm of v
func (v VecS) Mag() float64 {
	return sqrtfloat64(v.Dot(v))
}

// Unit returns the (l2) normalized vector of v
// or the zero vector if v is zero.
func (v VecS) Unit() VecS {
	d := v.Mag()
	switch d {
	case 1, 0:
		return v
	}
	return v.Div(d)
}

/*

var _ VecC = (*VecS)(nil)

// NewVecS creates a new VecS from the provided values.
func NewVecS(vs ...float64) VecS {
	out := VecS{}
	out.Set(vs...)
	return out
}

// At returns the value at dimension dim.
func (v *VecS) At(dim int) float64 {
	if dim >= len(*v) {
		return 0
	}
	return (*v)[dim]
}

// Get a slice of the underlying values
func (v *VecS) Get() []float64 {
	return (*v)[:]
}

// Set vector values
func (v *VecS) Set(vs ...float64) {
	initVecS(v, len(vs))
	for i, val := range vs {
		if i >= len(*v) {
			return
		}
		(*v)[i] = val
	}
}

// Vec2 version of this vector.
func (v *VecS) Vec2() Vec2 {
	return NewVec2((*v)[:]...)
}

// Vec3 version of this vector.
func (v *VecS) Vec3() Vec3 {
	return NewVec3((*v)[:]...)
}

// Vec4 version of this vector.
func (v *VecS) Vec4() Vec4 {
	return NewVec4((*v)[:]...)
}

*/
