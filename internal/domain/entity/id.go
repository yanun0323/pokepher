package entity

type ID int

func (id ID) Int() int {
	return int(id)
}
