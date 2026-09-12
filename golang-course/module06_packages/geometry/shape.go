// Package geometry provides basic shapes and area calculations.
// It demonstrates package organization and visibility rules in Go.
package geometry

import "math"

// ─────── 1. EXPORTED VS UNEXPORTED (VISIBILITY) ───────
// In Go, visibility is determined by the first letter of the identifier.
// - Capitalized (e.g., Circle, Area) -> Exported (Public)
// - lowercase (e.g., pi, calculateCache) -> Unexported (Private to the package)

// internalPi is unexported. It cannot be accessed outside the 'geometry' package.
const internalPi = math.Pi

// Shape is an exported interface.
type Shape interface {
	Area() float64
}

// Circle is an exported struct.
type Circle struct {
	// Radius is exported.
	Radius float64
	
	// cachedArea is unexported. Other packages cannot read or write this field directly.
	cachedArea float64 
}

// NewCircle is an exported constructor function.
// This is the idiomatic Go way to initialize structs that have unexported fields
// or require complex setup.
func NewCircle(radius float64) *Circle {
	return &Circle{
		Radius: radius,
		cachedArea: internalPi * radius * radius,
	}
}

// Area is an exported method on Circle.
func (c *Circle) Area() float64 {
	// If radius changed, we might need to recalculate. 
	// For this simple example, we just return the cache if it's there.
	if c.cachedArea == 0 {
		c.cachedArea = internalPi * c.Radius * c.Radius
	}
	return c.cachedArea
}
