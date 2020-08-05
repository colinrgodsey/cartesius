package f64

// Mul scales v by s
func (v VecN) Mul(s float64) (res VecN) {
	initVecN(&res, len(v))
	for i := range v {
		res[i] = v[i] * s
	}
	return
}

// Abs returns the absolute value of v
func (v VecN) Abs() (res VecN) {
	initVecN(&res, len(v))
	for i := range v {
		res[i] = absfloat64(v[i])
	}
	return
}

// Sub o from v
func (v VecN) Sub(o VecN) VecN {
	return v.Add(o.Neg())
}

// Div scales v by the multiplicative inverse of s
func (v VecN) Div(s float64) VecN {
	return v.Mul(1.0 / s)
}

// Neg returns the negative vector of v (-v)
func (v VecN) Neg() VecN {
	return v.Mul(-1.0)
}

// Inv returns the multiplicative inverse of v
func (v VecN) Inv() (res VecN) {
	initVecN(&res, len(v))
	for i := range v {
		res[i] = 1.0 / v[i]
	}
	return
}

// Add o to v
func (v VecN) Add(o VecN) (res VecN) {
	initVecN(&res, len(v))
	//TODO: this doesnt technically work for different size vecNs, or vecS
	for i := range v {
		res[i] = v[i] + o.At(i)
	}
	return
}

// MulV multiplies v*s per dim
func (v VecN) MulV(o VecN) (res VecN) {
	initVecN(&res, len(v))
	for i := range v {
		res[i] = v[i] * o.At(i)
	}
	return
}

// Dot returns the dot product of v and o (v⋅o)
func (v VecN) Dot(o VecN) (d float64) {
	for i := range v {
		d += v[i] * o.At(i)
	}
	return
}

// Within returns true if v is within the bounds of o
// considering both values as their absolute.
func (v VecN) Within(o VecN) bool {
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
func (v VecN) Eq(o VecN) bool {
	for i := range v {
		if v[i] != o[i] {
			return false
		}
	}
	return true
}

// Mag returns the L2 norm of v
func (v VecN) Mag() float64 {
	return sqrtfloat64(v.Dot(v))
}

// Unit returns the (l2) normalized vector of v
// or the zero vector if v is zero.
func (v VecN) Unit() VecN {
	d := v.Mag()
	switch d {
	case 1, 0:
		return v
	}
	return v.Div(d)
}

/* VECC_START */

var _ VecC = (*VecN)(nil)

// NewVecN creates a new VecN from the provided values.
func NewVecN(vs ...float64) VecN {
	out := VecN{}
	out.Set(vs...)
	return out
}

// At returns the value at dimension dim.
func (v *VecN) At(dim int) float64 {
	if dim >= len(*v) {
		return 0
	}
	return (*v)[dim]
}

// Get a slice of the underlying values
func (v *VecN) Get() []float64 {
	return (*v)[:]
}

// Set vector values
func (v *VecN) Set(vs ...float64) {
	initVecN(v, len(vs))
	copy((*v)[:], vs)
}

// Vec2 version of this vector.
func (v *VecN) Vec2() Vec2 {
	return NewVec2((*v)[:]...)
}

// Vec3 version of this vector.
func (v *VecN) Vec3() Vec3 {
	return NewVec3((*v)[:]...)
}

// Vec4 version of this vector.
func (v *VecN) Vec4() Vec4 {
	return NewVec4((*v)[:]...)
}

/* VECC_END */
