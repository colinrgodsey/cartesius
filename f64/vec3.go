package f64

var _ Vec = (*Vec3)(nil)

// Add o to v
func (v Vec3) Add(o Vec3) (res Vec3) {
	initVec3(&res, len(v))
	for i := range v {
		res[i] = v[i] + o[i]
	}
	return
}

// Mul scales v by s
func (v Vec3) Mul(s float64) (res Vec3) {
	initVec3(&res, len(v))
	for i := range v {
		res[i] = v[i] * s
	}
	return
}

// MulV multiplies v*s per dim
func (v Vec3) MulV(o Vec3) (res Vec3) {
	initVec3(&res, len(v))
	for i := range v {
		res[i] = v[i] * o[i]
	}
	return
}

// Abs returns the absolute value of v
func (v Vec3) Abs() (res Vec3) {
	initVec3(&res, len(v))
	for i := range v {
		res[i] = absfloat64(v[i])
	}
	return
}

// Sub o from v
func (v Vec3) Sub(o Vec3) Vec3 {
	return v.Add(o.Neg())
}

// Div scales v by the multiplicative inverse of s
func (v Vec3) Div(s float64) Vec3 {
	return v.Mul(1.0 / s)
}

// Neg returns the negative vector of v (-v)
func (v Vec3) Neg() Vec3 {
	return v.Mul(-1.0)
}

// Inv returns the multiplicative inverse of v
func (v Vec3) Inv() (res Vec3) {
	for i := range v {
		res[i] = 1.0 / v[i]
	}
	return
}

// Dot returns the dot product of v and o (v⋅o)
func (v Vec3) Dot(o Vec3) (d float64) {
	for i := range v {
		d += v[i] * o[i]
	}
	return
}

// Within returns true if v is within the bounds of o
// considering both values as their absolute.
func (v Vec3) Within(o Vec3) bool {
	v = v.Abs()
	o = o.Abs()
	for i := range v {
		if v[i] > o[i] {
			return false
		}
	}
	return true
}

// Eq returns true if v and o are equal
func (v Vec3) Eq(o Vec3) bool {
	for i := range v {
		if v[i] != o[i] {
			return false
		}
	}
	return true
}

// Mag returns the L2 norm of v
func (v Vec3) Mag() float64 {
	return sqrtfloat64(v.Dot(v))
}

// Unit returns the (l2) normalized vector of v
// or the zero vector if v is zero.
func (v Vec3) Unit() Vec3 {
	d := v.Mag()
	switch d {
	case 1, 0:
		return v
	}
	return v.Div(d)
}

// Get a slice of the underlying values
func (v *Vec3) Get() []float64 {
	return (*v)[:]
}

// Set vector values
func (v *Vec3) Set(vs ...float64) {
	initVec3(v, len(vs))
	for i, val := range vs {
		if i >= len(*v) {
			return
		}
		(*v)[i] = val
	}
}
