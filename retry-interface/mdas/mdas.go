package mdas

// Keep interfaces small
type Adder interface {
	Add(int, int) int
}

type Differ interface {
	Dif(int, int) int
}

type Product interface {
	product(int, int) int
}

type Quotient interface {
	quotient(int, int) int
}

// Struct implementing both
type MdasWriter struct {
	Offset int
}

func (m MdasWriter) Add(a, b int) int {
	return a + b + m.Offset
}

func (m MdasWriter) Dif(a, b int) int {
	return a - b + m.Offset
}

func (m MdasWriter) product(a, b int) int {
	return a*b + m.Offset
}

func (m MdasWriter) quotient(a, b int) int {
	return a/b + m.Offset
}
