package entity

type Status string

const (
	StatusNone Status = "None"
)

func (s Status) String() string {
	return string(s)
}
