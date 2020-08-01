/* FILE WAS AUTO-GENERATED FROM f64/vecN */

package f64

var _ Vec = (*Vec4)(nil)

// Add o to v
func (v Vec4) Add(o Vec4) (res Vec4) {
	initVec4(&res, len(v))
	for i := range v {
		res[i] = v[i] + o[i]
	}
	return
}

// Mul scales v by s
func (v Vec4) Mul(s float64) (res Vec4) {
	initVec4(&res, len(v))
	for i := range v {
		res[i] = v[i] * s
	}
	return
}

// MulV multiplies v*s per dim
func (v Vec4) MulV(o Vec4) (res Vec4) {
	initVec4(&res, len(v))
	for i := range v {
		res[i] = v[i] * o[i]
	}
	return
}

// Abs returns the absolute value of v
func (v Vec4) Abs() (res Vec4) {
	initVec4(&res, len(v))
	for i := range v {
		res[i] = absfloat64(v[i])
	}
	return
}

// Sub o from v
func (v Vec4) Sub(o Vec4) Vec4 {
	return v.Add(o.Neg())
}

// Div scales v by the multiplicative inverse of s
func (v Vec4) Div(s float64) Vec4 {
	return v.Mul(1.0 / s)
}

// Neg returns the negative vector of v (-v)
func (v Vec4) Neg() Vec4 {
	return v.Mul(-1.0)
}

// Inv returns the multiplicative inverse of v
func (v Vec4) Inv() (res Vec4) {
	for i := range v {
		res[i] = 1.0 / v[i]
	}
	return
}

// Dot returns the dot product of v and o (v⋅o)
func (v Vec4) Dot(o Vec4) (d float64) {
	for i := range v {
		d += v[i] * o[i]
	}
	return
}

// Within returns true if v is within the bounds of o
// considering both values as their absolute.
func (v Vec4) Within(o Vec4) bool {
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
func (v Vec4) Eq(o Vec4) bool {
	for i := range v {
		if v[i] != o[i] {
			return false
		}
	}
	return true
}

// Mag returns the L2 norm of v
func (v Vec4) Mag() float64 {
	return sqrtfloat64(v.Dot(v))
}

// Unit returns the (l2) normalized vector of v
// or the zero vector if v is zero.
func (v Vec4) Unit() Vec4 {
	d := v.Mag()
	switch d {
	case 1, 0:
		return v
	}
	return v.Div(d)
}

// Get a slice of the underlying values
func (v Vec4) Get() []float64 {
	return v[:]
}

// Set vector values
func (v *Vec4) Set(vs ...float64) {
	initVec4(v, len(vs))
	for i, val := range vs {
		if i >= len(*v) {
			return
		}
		(*v)[i] = val
	}
}
