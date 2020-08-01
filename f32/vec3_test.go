/* FILE WAS AUTO-GENERATED FROM f32/vecN_test */

/* FILE WAS AUTO-GENERATED FROM f64/vecN_test */

package f32

import (
	"fmt"
	"testing"
)

func TestVec3(t *testing.T) {
	a := NewVec3(1, 2, 3, 4)
	b := NewVec3(0, 1, 2, 3)
	one := NewVec3(1, 1, 1, 1)

	if !a.Add(b).Eq(NewVec3(1, 3, 5, 7)) {
		t.Fatal("add failed")
	}
	if !a.Sub(b).Eq(one) {
		t.Fatal("add failed")
	}
	if NewVec3(0, 1, 0, 0).Mag() != 1.0 {
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
	if absfloat32(NewVec3(2, 2, 2, 2).Mul(100).Unit().Mag()-1.0) > 1e-6 {
		fmt.Println(NewVec3(2, 2, 2, 2).Mul(100).Unit().Mag())
		t.Fatalf("unit mag failed")
	}
}
