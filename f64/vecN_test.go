package f64

import (
	"testing"
)

func TestVecN(t *testing.T) {
	a := NewVecN(1, 2, 3, 4)
	b := NewVecN(0, 1, 2, 3)
	one := NewVecN(1, 1, 1, 1)

	if !a.Add(b).Eq(NewVecN(1, 3, 5, 7)) {
		t.Fatal("add failed")
	}
	if !a.Sub(b).Eq(one) {
		t.Fatal("add failed")
	}
	if NewVecN(0, 1, 0, 0).Mag() != 1.0 {
		t.Fatalf("dot failed")
	}
	if !one.Inv().Eq(one) {
		t.Fatalf("inv failed")
	}
	if !one.MulV(one).Eq(one) {
		t.Fatalf("mulv failed")
	}
	if !one.Mul(2).Div(2).Eq(one) {
		t.Fatalf("mul->div failed")
	}
	if !one.Neg().Abs().Eq(one) {
		t.Fatalf("neg->abs failed")
	}
	if !one.Within(one.Mul(1.1)) {
		t.Fatalf("within failed")
	}
	if absfloat64(NewVecN(2, 2, 2, 2).Mul(100).Unit().Mag()-1.0) > 1e-6 {
		t.Fatalf("unit mag failed")
	}

	zv := NewVecN()
	zv.Set(1, 1, 1, 1)
	if !zv.Eq(one) {
		t.Fatalf("Set failed")
	}
}
