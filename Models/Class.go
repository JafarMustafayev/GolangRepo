package Models

type Class struct {
	ID       int
	Name     string
	Type     int
	MaxCount int
	Students []Student
}
