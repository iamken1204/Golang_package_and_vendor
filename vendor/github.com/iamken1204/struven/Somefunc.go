package struven

import "math/rand"

func NewRandomVen() Ven {
	v := NewVen(rand.Int())
	return v
}
