package vector2

import (
	"fmt"
	"math"
)

type Vector2 struct {
	X, Y float64
}

func Dot(ihs *Vector2, rhs *Vector2) float64 {
	return ihs.X*rhs.X + ihs.Y*rhs.Y
}

func Lerp(a *Vector2, b *Vector2, t float64) *Vector2 {
	return New(
		a.X+(b.X-a.X)*t,
		a.Y+(b.Y-a.Y)*t,
	)
}

func Distance(a *Vector2, b *Vector2) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func Reflect(ihs *Vector2, rhs *Vector2) *Vector2 {
	factor := -2.0 * Dot(ihs, rhs)
	return New(
		factor*ihs.X+rhs.X,
		factor*ihs.Y+rhs.Y,
	)
}

func New(x float64, y float64) *Vector2 {
	return &Vector2{X: x, Y: y}
}

func (v *Vector2) Copy() *Vector2 {
	return New(v.X, v.Y)
}

func (v *Vector2) Set(x float64, y float64) *Vector2 {
	v.X = x
	v.Y = y
	return v
}

func (v *Vector2) Add(other *Vector2) *Vector2 {
	return New(v.X+other.X, v.Y+other.Y)
}

func (v *Vector2) AddScalar(scalar float64) *Vector2 {
	return New(v.X+scalar, v.Y+scalar)
}

func (v *Vector2) AddScalars(x float64, y float64) *Vector2 {
	return New(v.X+x, v.Y+y)
}

func (v *Vector2) Sub(other *Vector2) *Vector2 {
	return New(v.X-other.X, v.Y-other.Y)
}

func (v *Vector2) SubScalar(scalar float64) *Vector2 {
	return New(v.X-scalar, v.Y-scalar)
}

func (v *Vector2) SubScalars(x float64, y float64) *Vector2 {
	return New(v.X-x, v.Y-y)
}

func (v *Vector2) Mul(other *Vector2) *Vector2 {
	return New(v.X*other.X, v.Y*other.Y)
}

func (v *Vector2) MulScalar(scalar float64) *Vector2 {
	return New(v.X*scalar, v.Y*scalar)
}

func (v *Vector2) MulScalars(x float64, y float64) *Vector2 {
	return New(v.X*x, v.Y*y)
}

func (v *Vector2) Div(other *Vector2) *Vector2 {
	return New(v.X/other.X, v.Y/other.Y)
}

func (v *Vector2) DivScalar(scalar float64) *Vector2 {
	return New(v.X/scalar, v.Y/scalar)
}

func (v *Vector2) DivScalars(x float64, y float64) *Vector2 {
	return New(v.X/x, v.Y/y)
}

func (v *Vector2) Distance(other *Vector2) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (v *Vector2) Cross(other *Vector2) float64 {
	return v.X*other.Y - v.Y*other.X
}

func (v *Vector2) Dot(other *Vector2) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v *Vector2) Lerp(other *Vector2, t float64) *Vector2 {
	return New(
		v.X+(other.X-v.X)*t,
		v.Y+(other.Y-v.Y)*t,
	)
}

func (v *Vector2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector2) Normalize() *Vector2 {
	m := v.Magnitude()

	if m > 0.0 {
		return v.DivScalar(m)
	} else {
		return v.Copy()
	}
}

func (v *Vector2) Reflect(other *Vector2) *Vector2 {
	factor := -2.0 * v.Dot(other)
	return New(
		factor*v.X+other.X,
		factor*v.Y+other.Y,
	)
}

func (v *Vector2) Equals(other *Vector2) bool {
	return v.X == other.X && v.Y == other.Y
}

func (v *Vector2) String() string {
	return fmt.Sprintf("Vector2(%f, %f)", v.X, v.Y)
}
