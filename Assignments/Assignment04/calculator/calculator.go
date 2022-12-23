package calculator

func Add(a *int, b *int) *int {
	add := *a + *b
	return &add
}

func Sub(a *int, b *int) *int {
	sub := *a - *b
	return &sub
}

func Mul(a *int, b *int) *int {
	mul := *a * *b
	return &mul
}

func Div(a *int, b *int) (*int, *int) {
	q := *a / *b
	r := *a % *b
	return &q, &r
}
