/* FILE WAS AUTO-GENERATED FROM f32/vecN */

/* FILE WAS AUTO-GENERATED FROM f64/vecN */

package f32

// Mul scales v by s
func (v Vec2) Mul(s float32) (res Vec2) {
	initVec2(&res, len(v))
	for i := range v {
		res[i] = v[i] * s
	}
	return
}

// Abs returns the absolute value of v
func (v Vec2) Abs() (res Vec2) {
	initVec2(&res, len(v))
	for i := range v {
		res[i] = absfloat32(v[i])
	}
	return
}

// Sub o from v
func (v Vec2) Sub(o Vec2) Vec2 {
	return v.Add(o.Neg())
}

// Div scales v by the multiplicative inverse of s
func (v Vec2) Div(s float32) Vec2 {
	return v.Mul(1.0 / s)
}

// Neg returns the negative vector of v (-v)
func (v Vec2) Neg() Vec2 {
	return v.Mul(-1.0)
}

// Inv returns the multiplicative inverse of v
func (v Vec2) Inv() (res Vec2) {
	initVec2(&res, len(v))
	for i := range v {
		res[i] = 1.0 / v[i]
	}
	return
}

// Add o to v
func (v Vec2) Add(o Vec2) (res Vec2) {
	initVec2(&res, len(v))
	//TODO: this doesnt technically work for different size vecNs, or vecS
	for i := range v {
		res[i] = v[i] + o.At(i)
	}
	return
}

// MulV multiplies v*s per dim
func (v Vec2) MulV(o Vec2) (res Vec2) {
	initVec2(&res, len(v))
	for i := range v {
		res[i] = v[i] * o.At(i)
	}
	return
}

// Dot returns the dot product of v and o (v⋅o)
func (v Vec2) Dot(o Vec2) (d float32) {
	for i := range v {
		d += v[i] * o.At(i)
	}
	return
}

// Within returns true if v is within the bounds of o
// considering both values as their absolute.
func (v Vec2) Within(o Vec2) bool {
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
func (v Vec2) Eq(o Vec2) bool {
	for i := range v {
		if v[i] != o[i] {
			return false
		}
	}
	return true
}

// Mag returns the L2 norm of v
func (v Vec2) Mag() float32 {
	return sqrtfloat32(v.Dot(v))
}

// Unit returns the (l2) normalized vector of v
// or the zero vector if v is zero.
func (v Vec2) Unit() Vec2 {
	d := v.Mag()
	switch d {
	case 1, 0:
		return v
	}
	return v.Div(d)
}

/* VECC_START */

var _ VecC = (*Vec2)(nil)

// NewVec2 creates a new Vec2 from the provided values.
func NewVec2(vs ...float32) Vec2 {
	out := Vec2{}
	out.Set(vs...)
	return out
}

// At returns the value at dimension dim.
func (v *Vec2) At(dim int) float32 {
	if dim >= len(*v) {
		return 0
	}
	return (*v)[dim]
}

// Get a slice of the underlying values
func (v *Vec2) Get() []float32 {
	return (*v)[:]
}

// Set vector values
func (v *Vec2) Set(vs ...float32) {
	initVec2(v, len(vs))
	copy((*v)[:], vs)
}

// Vec2 version of this vector.
func (v *Vec2) Vec2() Vec2 {
	return NewVec2((*v)[:]...)
}

// Vec3 version of this vector.
func (v *Vec2) Vec3() Vec3 {
	return NewVec3((*v)[:]...)
}

// Vec4 version of this vector.
func (v *Vec2) Vec4() Vec4 {
	return NewVec4((*v)[:]...)
}

/* VECC_END */
