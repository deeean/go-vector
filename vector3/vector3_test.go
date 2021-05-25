package vector3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	a := New(2, 4, 6)

	assert.Equal(t, a.X, 2.0)
	assert.Equal(t, a.Y, 4.0)
	assert.Equal(t, a.Z, 6.0)
}

func TestVector3_Copy(t *testing.T) {
	a := New(2, 4, 6)
	res := a.Copy()

	assert.Equal(t, a.X, res.X)
	assert.Equal(t, a.Y, res.Y)
	assert.Equal(t, a.Z, res.Z)
}

func TestVector3_Set(t *testing.T) {
	a := New(2, 4, 6)
	a.Set(5, 4, 5)

	assert.Equal(t, a.X, 5.0)
	assert.Equal(t, a.Y, 4.0)
	assert.Equal(t, a.Z, 5.0)
}

func TestVector3_Add(t *testing.T) {
	a := New(1, 2, 3)
	b := New(4, 3, 1)
	res := a.Add(b)

	assert.Equal(t, res.X, 5.0)
	assert.Equal(t, res.Y, 5.0)
	assert.Equal(t, res.Z, 4.0)
}

func TestVector3_AddScalar(t *testing.T) {
	a := New(1, 2, 3)
	res := a.AddScalar(1.0)

	assert.Equal(t, res.X, 2.0)
	assert.Equal(t, res.Y, 3.0)
	assert.Equal(t, res.Z, 4.0)
}

func TestVector3_AddScalars(t *testing.T) {
	a := New(1, 2, 3)
	res := a.AddScalars(1, 2, 3)

	assert.Equal(t, res.X, 2.0)
	assert.Equal(t, res.Y, 4.0)
	assert.Equal(t, res.Z, 6.0)
}

func TestVector3_Sub(t *testing.T) {
	a := New(1, 2, 3)
	b := New(4, 3, 1)
	res := a.Sub(b)

	assert.Equal(t, res.X, -3.0)
	assert.Equal(t, res.Y, -1.0)
	assert.Equal(t, res.Z, 2.0)
}

func TestVector3_SubScalar(t *testing.T) {
	a := New(1, 2, 3)
	res := a.SubScalar(1.0)

	assert.Equal(t, res.X, 0.0)
	assert.Equal(t, res.Y, 1.0)
	assert.Equal(t, res.Z, 2.0)
}

func TestVector3_SubScalars(t *testing.T) {
	a := New(1, 2, 3)
	res := a.SubScalars(1, 2, 3)

	assert.Equal(t, res.X, 0.0)
	assert.Equal(t, res.Y, 0.0)
	assert.Equal(t, res.Z, 0.0)
}

func TestVector3_Mul(t *testing.T) {
	a := New(1, 2, 3)
	b := New(4, 3, 1)
	res := a.Mul(b)

	assert.Equal(t, res.X, 4.0)
	assert.Equal(t, res.Y, 6.0)
	assert.Equal(t, res.Z, 3.0)
}

func TestVector3_MulByScalar(t *testing.T) {
	a := New(1, 2, 3)
	b := 4.0
	res := a.MulByScalar(b)

	assert.Equal(t, res.X, 4.0)
	assert.Equal(t, res.Y, 8.0)
	assert.Equal(t, res.Z, 12.0)
}

func TestVector3_Div(t *testing.T) {
	a := New(1, 2, 3)
	b := New(4, 3, 1)
	res := a.Div(b)

	assert.Equal(t, res.X, 0.25)
	assert.Equal(t, res.Y, 0.6666666666666666)
	assert.Equal(t, res.Z, 3.0)
}

func TestVector3_DivByScalar(t *testing.T) {
	a := New(1, 2, 3)
	b := 4.0
	res := a.DivByScalar(b)

	assert.Equal(t, res.X, 0.25)
	assert.Equal(t, res.Y, 0.5)
	assert.Equal(t, res.Z, 0.75)
}

func TestVector3_Distance(t *testing.T) {
	a := New(1, 2, 3)
	b := New(4, 3, 1)
	res := a.Distance(b)

	assert.Equal(t, res, 3.7416573867739413)
}

func TestVector3_Dot(t *testing.T) {
	a := New(1, 2, 3)
	b := New(4, 3, 1)
	res := a.Dot(b)

	assert.Equal(t, res, 13.0)
}

func TestVector3_Cross(t *testing.T) {
	a := New(1, 2, 3)
	b := New(4, 3, 1)
	res := a.Cross(b)

	assert.Equal(t, res.X, -7.0)
	assert.Equal(t, res.Y, 11.0)
	assert.Equal(t, res.Z, -5.0)
}

func TestVector3_Magnitude(t *testing.T) {
	a := New(3, 2, 3)
	res := a.Magnitude()

	assert.Equal(t, res, 4.69041575982343)
}

func TestVector3_Normalize(t *testing.T) {
	a := New(3, 2, 3)
	b := New(0.000003, 0.000002, 0.000003)
	res := a.Normalize()
	res2 := b.Normalize()

	assert.Equal(t, res.X, 0.6396021490668313)
	assert.Equal(t, res.Y, 0.42640143271122083)
	assert.Equal(t, res.Z, 0.6396021490668313)
	assert.True(t, b.Equals(res2))
}

func TestVector3_Reflect(t *testing.T) {
	a := New(2, 1, 0)
	b := New(6, 6, 6)
	res := a.Reflect(b)

	assert.Equal(t, res.X, -66.0)
	assert.Equal(t, res.Y, -30.0)
	assert.Equal(t, res.Z, 6.0)
}

func TestVector3_Lerp(t *testing.T) {
	a := New(0, 0, 0)
	b := New(10, 10, 10)
	res := a.Lerp(b, 0.5)

	assert.Equal(t, res.X, 5.0)
	assert.Equal(t, res.Y, 5.0)
	assert.Equal(t, res.Z, 5.0)
}

func TestVector3_Equals(t *testing.T) {
	a := New(5, 3, 5)
	b := New(4, 3, 3)

	assert.False(t, a.Equals(b))
}

func TestDistance(t *testing.T) {
	a := New(1, 2, 3)
	b := New(4, 3, 1)
	res := Distance(a, b)

	assert.Equal(t, res, 3.7416573867739413)
}

func TestDot(t *testing.T) {
	a := New(1, 2, 3)
	b := New(4, 3, 1)
	res := Dot(a, b)

	assert.Equal(t, res, 13.0)
}

func TestCross(t *testing.T) {
	a := New(1, 2, 3)
	b := New(4, 3, 1)
	res := Cross(a, b)

	assert.Equal(t, res.X, -7.0)
	assert.Equal(t, res.Y, 11.0)
	assert.Equal(t, res.Z, -5.0)
}

func TestReflect(t *testing.T) {
	a := New(2, 1, 0)
	b := New(6, 6, 6)
	res := Reflect(a, b)

	assert.Equal(t, res.X, -66.0)
	assert.Equal(t, res.Y, -30.0)
	assert.Equal(t, res.Z, 6.0)
}

func TestLerp(t *testing.T) {
	a := New(0, 0, 0)
	b := New(10, 10, 10)
	res := Lerp(a, b, 0.5)

	assert.Equal(t, res.X, 5.0)
	assert.Equal(t, res.Y, 5.0)
	assert.Equal(t, res.Z, 5.0)
}
