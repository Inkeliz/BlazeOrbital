package main

import (
	"math"
)

type Point struct {
	X, Y float32
}

func NewPoint(x, y float32) Point {
	return Point{X: x, Y: y}
}

// LengthSquared is public float LengthSquared => X* X + Y * Y;
func (a Point) LengthSquared() float32 {
	return a.X*a.X + a.Y*a.Y
}

// Length is public float Length => (float)Math.Sqrt(LengthSquared);
func (a Point) Length() float32 {
	return float32(math.Sqrt(float64(a.LengthSquared())))
}

// Direction is public Point Direction => this / Length;
func (a Point) Direction() Point {
	return a.DivScale(a.Length())
}

// Equal is public static bool operator ==(Point a, Point b) => (a.X == b.X) && (a.Y == b.Y);
func (a Point) Equal(b Point) bool {
	return a == b
}

// Different is public static bool operator !=(Point a, Point b) => (a.X != b.X) || (a.Y != b.Y);
func (a Point) Different(b Point) bool {
	return a != b
}

// Add is public static Point operator +(Point a, Point b) => new Point(a.X + b.X, a.Y + b.Y);
func (a Point) Add(b Point) Point {
	return Point{X: a.X + b.X, Y: a.Y + b.Y}
}

// Sub is public static Point operator -(Point a, Point b) => new Point(a.X - b.X, a.Y - b.Y);
func (a Point) Sub(b Point) Point {
	return Point{X: a.X - b.X, Y: a.Y - b.Y}
}

// Mul is public static Point operator *(Point a, Point b) => new Point(a.X * b.X, a.Y * b.Y);
func (a Point) Mul(b Point) Point {
	return Point{X: a.X * b.X, Y: a.Y * b.Y}
}

// MulScale is public static Point operator *(Point a, float scale) => new Point(a.X * scale, a.Y * scale);
func (a Point) MulScale(scale float32) Point {
	return Point{X: a.X * scale, Y: a.Y * scale}
}

// DivScale is public static Point operator /(Point a, float scale) => new Point(a.X / scale, a.Y / scale);
func (a Point) DivScale(scale float32) Point {
	return Point{X: a.X / scale, Y: a.Y / scale}
}
