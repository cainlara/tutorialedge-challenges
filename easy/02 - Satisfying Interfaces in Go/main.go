package main

type Employee interface {
	Language() string
	Age() int
}

type Engineer struct {
	Name string
}

func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}

// Uncomment next fuction to make code runnable
func (e Engineer) Age() int {
	return 30
}

func main() {
	var programmers []Employee
	elliot := Engineer{Name: "Elliot"}
	programmers = append(programmers, elliot)
}
