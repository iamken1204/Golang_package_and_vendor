package struven

type Ven struct {
	ID int
}

func NewVen(id int) Ven {
	v := Ven{id}
	return v
}
