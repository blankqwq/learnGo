package main

type Test struct {
	Name string
	Age  int
}

func main() {
	var t *Test = &Test{}
	println(t)
	t.Name = "123"
	println(t.Age, t.Name)
}
