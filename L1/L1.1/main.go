package main

type Human struct {
	Name string
}

func (h *Human) SayHello() {
	println(
		"Hello, my name is",
		h.Name,
	)
}

type Action struct {
	Human
}

func main() {
	gleb := Human{Name: "Gleb"}

	action := Action{Human: gleb}

	action.SayHello()
}
