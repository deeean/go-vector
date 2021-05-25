package vector3

import "math"

const Epsilon = 0.00001

type Vector3 struct {
	X, Y, Z float64
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
	return New(v.X + other.X, v.Y + other.Y, v.Z + other.Z)
}

func (v *Vector3) Sub(other *Vector3) *Vector3 {
	return New(v.X - other.X, v.Y - other.Y, v.Z - other.Z)
}

func (v *Vector3) Mul(other *Vector3) *Vector3 {
	return New(v.X * other.X, v.Y * other.Y, v.Z * other.Z)
}

func (v *Vector3) MulByScalar(scalar float64) *Vector3 {
	return New(v.X * scalar, v.Y * scalar, v.Z * scalar)
}

func (v *Vector3) Div(other *Vector3) *Vector3 {
	return New(v.X / other.X, v.Y / other.Y, v.Z / other.Z)
}

func (v *Vector3) DivByScalar(scalar float64) *Vector3 {
	return New(v.X / scalar, v.Y / scalar, v.Z / scalar)
}

func (v *Vector3) Distance(other *Vector3) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	dz := v.Z - other.Z
	return math.Sqrt(dx * dx + dy * dy + dz * dz)
}

func (v *Vector3) Dot(other *Vector3) float64 {
	return v.X * other.X + v.Y * other.Y + v.Z * other.Z
}

func (v *Vector3) Cross(other *Vector3) *Vector3 {
	return New(
		v.Y * other.Z - v.Z * other.Y,
		v.Z * other.X - v.X * other.Z,
		v.X * other.Y - v.Y * other.X,
	)
}

func (v *Vector3) Lerp(other *Vector3, t float64) *Vector3 {
	return New(
		v.X + (other.X - v.X) * t,
		v.Y + (other.Y - v.Y) * t,
		v.Z + (other.Z - v.Z) * t,
	)
}

func (v *Vector3) Magnitude() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

func (v *Vector3) Normalize() *Vector3 {
	m := v.Magnitude()

	if (m > Epsilon) {
		return v.DivByScalar(m)
	} else {
		return v.Copy()
	}
}

func (v *Vector3) Reflect(other *Vector3) *Vector3 {
	factor := -2.0 * v.Dot(other);
	return New(factor * v.X + other.X,
		factor * v.Y + other.Y,
		factor * v.Z + other.Z)
}

func (v *Vector3) Equals(other *Vector3) bool {
	return v.X == other.X && v.Y == other.Y && v.Z == other.Z
}