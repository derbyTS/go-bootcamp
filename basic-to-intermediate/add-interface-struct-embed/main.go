package main

import (
	"fmt"
	"math"
)

// --- THE INTERFACE ---
// PATTERN: ABSTRACTION
// Like: type Context interface { ... }
// This defines the "Contract" that all our math objects must follow.
type Shape interface {
	Area() float64
	Name() string
}

// --- 1. THE BASE (Internal Utility) ---
// PATTERN: NULL OBJECT / PROXY - Not quite right according to CGPT
// Like: type emptyCtx struct { ... }
// VIRTUE: It provides "default" behavior so other structs don't have to
// write every single method if they don't want to.
type baseShape struct{}

func (baseShape) Area() float64 { return 0 }
func (baseShape) Name() string  { return "Generic Shape" }

// --- 2. THE IDENTITY (Empty Embedding) ---
// PATTERN: MARKER / TYPE IDENTITY - Not quite right according to CGPT
// Like: type todoCtx struct { emptyCtx }
// VIRTUE: It uses 0 bytes of memory but has a unique "Type" and a
// custom String/Name representation for debugging.
type todoShape struct{ baseShape }

func (todoShape) Name() string { return "context.TODO_SHAPE" }

// --- 3. THE DECORATOR (Interface Wrapping) ---
// PATTERN: DECORATOR / MIDDLEWARE
// Like: type withoutCancelCtx struct { c Context }
// VIRTUE: It "holds" a real Shape inside it. It shadows (overrides)
// certain methods while letting others pass through to the original.
type multiplierShape struct {
	s      Shape // Embedding the Interface
	factor float64
}

// SHADOWING: We intercept the Area() call to change the result.
func (m multiplierShape) Area() float64 {
	return m.s.Area() * m.factor
}

func (m multiplierShape) Name() string {
	return fmt.Sprintf("Boosted %s (x%.1f)", m.s.Name(), m.factor)
}

// --- 4. THE OVERRIDE (Concrete Embedding) ---
// PATTERN: COMPOSITION OVER INHERITANCE
// Like: type timerCtx struct { cancelCtx; deadline time.Time }
// VIRTUE: It embeds the base to "get in the door" of the interface,
// then adds its own fields and overrides the logic it actually cares about.
type circle struct {
	baseShape // PROMOTION: Gets Name() for free if we didn't override it
	radius    float64
}

// SHADOWING: This "wins" over the baseShape.Area()
func (c circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) Name() string {
	return "Perfect Circle"
}

// --- EXECUTION ---

func main() {
	// 1. Identity Pattern: Calling a Factory-style creation
	// In context.go, this would be: ctx := context.TODO()
	var todo todoShape
	measure(todo)

	// 2. Override Pattern: Creating a concrete specialized type
	c := circle{radius: 10}
	measure(c)

	// 3. Decorator Pattern: Wrapping an object to add behavior
	// In context.go, this would be: ctx, cancel := context.WithCancel(parent)
	bigCircle := multiplierShape{
		s:      c,
		factor: 5.0,
	}
	measure(bigCircle)
}

func measure(s Shape) {
	fmt.Printf("[%s]\n", s.Name())
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Println("--------------------")
}
