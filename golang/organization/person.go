package organization

type Identificable interface {
	ID() string
}

type Person struct {
}

func (p Person) ID() string {
	return "123456"
}
