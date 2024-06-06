package vector3

import (
	"fmt"
	"math"
)

type Vector3 struct {
	X, Y, Z float64
}

func Dot(ihs *Vector3, rhs *Vector3) float64 {
	return ihs.X*rhs.X + ihs.Y*rhs.Y + ihs.Z*rhs.Z
}

func Cross(ihs *Vector3, rhs *Vector3) *Vector3 {
	return New(
		ihs.Y*rhs.Z-ihs.Z*rhs.Y,
		ihs.Z*rhs.X-ihs.X*rhs.Z,
		ihs.X*rhs.Y-ihs.Y*rhs.X,
	)
}

func Lerp(a *Vector3, b *Vector3, t float64) *Vector3 {
	return New(
		a.X+(b.X-a.X)*t,
		a.Y+(b.Y-a.Y)*t,
		a.Z+(b.Z-a.Z)*t,
	)
}

func Distance(a *Vector3, b *Vector3) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func Reflect(ihs *Vector3, rhs *Vector3) *Vector3 {
	factor := -2.0 * Dot(ihs, rhs)
	return New(
		factor*ihs.X+rhs.X,
		factor*ihs.Y+rhs.Y,
		factor*ihs.Z+rhs.Z,
	)
}

func New(x float64, y float64, z float64) *Vector3 {
	return &Vector3{X: x, Y: y, Z: z}
}

func (v *Vector3) Copy() *Vector3 {
	return New(v.X, v.Y, v.Z)
}

func (v *Vector3) Set(x float64, y float64, z float64) *Vector3 {
	v.X = x
	v.Y = y
	v.Z = z
	return v
}

func (v *Vector3) Add(other *Vector3) *Vector3 {
	return New(v.X+other.X, v.Y+other.Y, v.Z+other.Z)
}

func (v *Vector3) AddScalar(scalar float64) *Vector3 {
	return New(v.X+scalar, v.Y+scalar, v.Z+scalar)
}

func (v *Vector3) AddScalars(x float64, y float64, z float64) *Vector3 {
	return New(v.X+x, v.Y+y, v.Z+z)
}

func (v *Vector3) Sub(other *Vector3) *Vector3 {
	return New(v.X-other.X, v.Y-other.Y, v.Z-other.Z)
}

func (v *Vector3) SubScalar(scalar float64) *Vector3 {
	return New(v.X-scalar, v.Y-scalar, v.Z-scalar)
}

func (v *Vector3) SubScalars(x float64, y float64, z float64) *Vector3 {
	return New(v.X-x, v.Y-y, v.Z-z)
}

func (v *Vector3) Mul(other *Vector3) *Vector3 {
	return New(v.X*other.X, v.Y*other.Y, v.Z*other.Z)
}

func (v *Vector3) MulScalar(scalar float64) *Vector3 {
	return New(v.X*scalar, v.Y*scalar, v.Z*scalar)
}

func (v *Vector3) MulScalars(x float64, y float64, z float64) *Vector3 {
	return New(v.X*x, v.Y*y, v.Z*z)
}

func (v *Vector3) Div(other *Vector3) *Vector3 {
	return New(v.X/other.X, v.Y/other.Y, v.Z/other.Z)
}

func (v *Vector3) DivScalar(scalar float64) *Vector3 {
	return New(v.X/scalar, v.Y/scalar, v.Z/scalar)
}

func (v *Vector3) DivScalars(x float64, y float64, z float64) *Vector3 {
	return New(v.X/x, v.Y/y, v.Z/z)
}

func (v *Vector3) Distance(other *Vector3) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	dz := v.Z - other.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (v *Vector3) Dot(other *Vector3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v *Vector3) Cross(other *Vector3) *Vector3 {
	return New(
		v.Y*other.Z-v.Z*other.Y,
		v.Z*other.X-v.X*other.Z,
		v.X*other.Y-v.Y*other.X,
	)
}

func (v *Vector3) Lerp(other *Vector3, t float64) *Vector3 {
	return New(
		v.X+(other.X-v.X)*t,
		v.Y+(other.Y-v.Y)*t,
		v.Z+(other.Z-v.Z)*t,
	)
}

func (v *Vector3) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vector3) Normalize() *Vector3 {
	m := v.Magnitude()

	if m > 0.0 {
		return v.DivScalar(m)
	} else {
		return v.Copy()
	}
}

func (v *Vector3) Reflect(other *Vector3) *Vector3 {
	factor := -2.0 * v.Dot(other)
	return New(
		factor*v.X+other.X,
		factor*v.Y+other.Y,
		factor*v.Z+other.Z,
	)
}

func (v *Vector3) Equals(other *Vector3) bool {
	return v.X == other.X && v.Y == other.Y && v.Z == other.Z
}

func (v *Vector3) String() string {
	return fmt.Sprintf("Vector3(%f, %f, %f)", v.X, v.Y, v.Z)
}
