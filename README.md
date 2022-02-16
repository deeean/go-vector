# go-vector
A Vector library for 2D and 3D applications.

[![github-actions](https://github.com/deeean/go-vector/actions/workflows/ci.yml/badge.svg)](https://github.com/deeean/go-vector)
[![codecov](https://codecov.io/gh/deeean/go-vector/branch/master/graph/badge.svg?token=ITVBDT948V)](https://codecov.io/gh/deeean/go-vector)

## Installation
```shell
go get -u github.com/deeean/go-vector
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/deeean/go-vector/vector3"
)

func main() {
	// basic usage
	a := vector3.New(1, 2, 3)
	b := vector3.New(1, 2, 3)
	fmt.Println(a.Add(b))

	// method chaining
	c := vector3.New(0, 3, 5).
		DivScalars(1, 2, 4).
		MulScalars(0, 5, 3).
		Magnitude()
	fmt.Println(c)
}
```

## Documentation

### Vector2
```go
func New(x float64, y float64) *Vector2

func Dot(ihs *Vector2, rhs *Vector2) float64

func Lerp(a *Vector2, b *Vector2, t float64) *Vector2

func Distance(a *Vector2, b *Vector2) float64

func Reflect(ihs *Vector2, rhs *Vector2) *Vector2

func (v *Vector2) Copy() *Vector2

func (v *Vector2) Set(x float64, y float64) *Vector2

func (v *Vector2) Add(other *Vector2) *Vector2

func (v *Vector2) AddScalar(scalar float64) *Vector2

func (v *Vector2) AddScalars(x float64, y float64) *Vector2

func (v *Vector2) Sub(other *Vector2) *Vector2

func (v *Vector2) SubScalar(scalar float64) *Vector2

func (v *Vector2) SubScalars(x float64, y float64) *Vector2

func (v *Vector2) Mul(other *Vector2) *Vector2

func (v *Vector2) MulScalar(scalar float64) *Vector2

func (v *Vector2) MulScalars(x float64, y float64) *Vector2

func (v *Vector2) Div(other *Vector2) *Vector2

func (v *Vector2) DivScalar(scalar float64) *Vector2

func (v *Vector2) DivScalars(x float64, y float64) *Vector2

func (v *Vector2) Distance(other *Vector2) float64

func (v *Vector2) Dot(other *Vector2) float64

func (v *Vector2) Lerp(other *Vector2, t float64) *Vector2

func (v *Vector2) Magnitude() float64

func (v *Vector2) Normalize() *Vector2

func (v *Vector2) Reflect(other *Vector2) *Vector2

func (v *Vector2) Equals(other *Vector2) bool

func (v *Vector2) ToString() string
```

### Vector3
```go
func New(x float64, y float64, z float64) *Vector3

func Dot(ihs *Vector3, rhs *Vector3) float64

func Cross(ihs *Vector3, rhs *Vector3) *Vector3

func Lerp(a *Vector3, b *Vector3, t float64) *Vector3

func Distance(a *Vector3, b *Vector3) float64

func Reflect(ihs *Vector3, rhs *Vector3) *Vector3

func (v *Vector3) Copy() *Vector3

func (v *Vector3) Set(x float64, y float64, z float64) *Vector3

func (v *Vector3) Add(other *Vector3) *Vector3

func (v *Vector3) AddScalar(scalar float64) *Vector3

func (v *Vector3) AddScalars(x float64, y float64, z float64) *Vector3

func (v *Vector3) Sub(other *Vector3) *Vector3

func (v *Vector3) SubScalar(scalar float64) *Vector3

func (v *Vector3) SubScalars(x float64, y float64, z float64)

func (v *Vector3) Mul(other *Vector3) *Vector3

func (v *Vector3) MulScalar(scalar float64) *Vector3

func (v *Vector3) MulScalars(x float64, y float64, z float64) *Vector3

func (v *Vector3) Div(other *Vector3) *Vector3

func (v *Vector3) DivScalar(scalar float64) *Vector3

func (v *Vector3) DivScalars(x float64, y float64, z float64) *Vector3

func (v *Vector3) Distance(other *Vector3) float64

func (v *Vector3) Dot(other *Vector3) float64

func (v *Vector3) Cross(other *Vector3) *Vector3

func (v *Vector3) Lerp(other *Vector3, t float64) *Vector3

func (v *Vector3) Magnitude() float64

func (v *Vector3) Normalize() *Vector3

func (v *Vector3) Reflect(other *Vector3) *Vector3

func (v *Vector3) Equals(other *Vector3) bool

func (v *Vector3) ToString() string
```
