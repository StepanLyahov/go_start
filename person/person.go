package person

type Person struct {
	fio Fio
	age int8
}

type Fio struct {
	name      string
	firstname string
	lastname  string
}
