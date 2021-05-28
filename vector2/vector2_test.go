package vector2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	a := New(2, 4)

	assert.Equal(t, a.X, 2.0)
	assert.Equal(t, a.Y, 4.0)
}

func TestVector2_Copy(t *testing.T) {
	a := New(2, 4)
	res := a.Copy()

	assert.Equal(t, a.X, res.X)
	assert.Equal(t, a.Y, res.Y)
}

func TestVector2_Set(t *testing.T) {
	a := New(2, 4)
	a.Set(5, 4)

	assert.Equal(t, a.X, 5.0)
	assert.Equal(t, a.Y, 4.0)
}

func TestVector2_Add(t *testing.T) {
	a := New(1, 2)
	b := New(4, 3)
	res := a.Add(b)

	assert.Equal(t, res.X, 5.0)
	assert.Equal(t, res.Y, 5.0)
}

func TestVector2_AddScalar(t *testing.T) {
	a := New(1, 2)
	res := a.AddScalar(1.0)

	assert.Equal(t, res.X, 2.0)
	assert.Equal(t, res.Y, 3.0)
}

func TestVector2_AddScalars(t *testing.T) {
	a := New(1, 2)
	res := a.AddScalars(1, 2, 3)

	assert.Equal(t, res.X, 2.0)
	assert.Equal(t, res.Y, 4.0)
}

func TestVector2_Sub(t *testing.T) {
	a := New(1, 2)
	b := New(4, 3)
	res := a.Sub(b)

	assert.Equal(t, res.X, -3.0)
	assert.Equal(t, res.Y, -1.0)
}

func TestVector2_SubScalar(t *testing.T) {
	a := New(1, 2)
	res := a.SubScalar(1.0)

	assert.Equal(t, res.X, 0.0)
	assert.Equal(t, res.Y, 1.0)
}

func TestVector2_SubScalars(t *testing.T) {
	a := New(1, 2)
	res := a.SubScalars(1, 2, 3)

	assert.Equal(t, res.X, 0.0)
	assert.Equal(t, res.Y, 0.0)
}

func TestVector2_Mul(t *testing.T) {
	a := New(1, 2)
	b := New(4, 3)
	res := a.Mul(b)

	assert.Equal(t, res.X, 4.0)
	assert.Equal(t, res.Y, 6.0)
}

func TestVector2_MulScalar(t *testing.T) {
	a := New(1, 2)
	b := 4.0
	res := a.MulScalar(b)

	assert.Equal(t, res.X, 4.0)
	assert.Equal(t, res.Y, 8.0)
}

func TestVector2_MulScalars(t *testing.T) {
	a := New(1, 2)
	res := a.MulScalars(1.0, 2.0, 3.0)

	assert.Equal(t, res.X, 1.0)
	assert.Equal(t, res.Y, 4.0)
}

func TestVector2_Div(t *testing.T) {
	a := New(1, 2)
	b := New(4, 3)
	res := a.Div(b)

	assert.Equal(t, res.X, 0.25)
	assert.Equal(t, res.Y, 0.6666666666666666)
}

func TestVector2_DivScalar(t *testing.T) {
	a := New(1, 2)
	b := 4.0
	res := a.DivScalar(b)

	assert.Equal(t, res.X, 0.25)
	assert.Equal(t, res.Y, 0.5)
}

func TestVector2_DivScalars(t *testing.T) {
	a := New(1, 2)
	res := a.DivScalars(1.0, 2.0, 3.0)

	assert.Equal(t, res.X, 1.0)
	assert.Equal(t, res.Y, 1.0)
}

func TestVector2_Distance(t *testing.T) {
	a := New(1, 2)
	b := New(4, 3)
	res := a.Distance(b)

	assert.Equal(t, res, 3.1622776601683795)
}

func TestVector2_Dot(t *testing.T) {
	a := New(1, 2)
	b := New(4, 3)
	res := a.Dot(b)

	assert.Equal(t, res, 10.0)
}

func TestVector2_Magnitude(t *testing.T) {
	a := New(3, 2)
	res := a.Magnitude()

	assert.Equal(t, res, 3.605551275463989)
}

func TestVector2_Normalize(t *testing.T) {
	a := New(3, 2)
	b := New(0.000003, 0.000002)
	res := a.Normalize()
	res2 := b.Normalize()

	assert.Equal(t, res.X, 0.8320502943378437)
	assert.Equal(t, res.Y, 0.5547001962252291)
	assert.True(t, b.Equals(res2))
}

func TestVector2_Reflect(t *testing.T) {
	a := New(2, 1)
	b := New(6, 6)
	res := a.Reflect(b)

	assert.Equal(t, res.X, -66.0)
	assert.Equal(t, res.Y, -30.0)
}

func TestVector2_Lerp(t *testing.T) {
	a := New(0, 0)
	b := New(10, 10)
	res := a.Lerp(b, 0.5)

	assert.Equal(t, res.X, 5.0)
	assert.Equal(t, res.Y, 5.0)
}

func TestVector2_Equals(t *testing.T) {
	a := New(5, 3)
	b := New(4, 3)

	assert.False(t, a.Equals(b))
}

func TestDistance(t *testing.T) {
	a := New(1, 2)
	b := New(4, 3)
	res := Distance(a, b)

	assert.Equal(t, res, 3.1622776601683795)
}

func TestDot(t *testing.T) {
	a := New(1, 2)
	b := New(4, 3)
	res := Dot(a, b)

	assert.Equal(t, res, 10.0)
}

func TestReflect(t *testing.T) {
	a := New(2, 1)
	b := New(6, 6)
	res := Reflect(a, b)

	assert.Equal(t, res.X, -66.0)
	assert.Equal(t, res.Y, -30.0)
}

func TestLerp(t *testing.T) {
	a := New(0, 0)
	b := New(10, 10)
	res := Lerp(a, b, 0.5)

	assert.Equal(t, res.X, 5.0)
	assert.Equal(t, res.Y, 5.0)
}
