package person

import (
	"testing"
)

func TestCreatePerson(t *testing.T) {
	var person1 Person = Person{Fio{"Lyahov", "Stepan", "Romanovich"}, 21}
	t.Log(person1)

	var person2 Person = Person{fio: Fio{
		name: "Stepan",
	}}
	t.Log(person2)

	var person3 *Person = &Person{fio: Fio{
		name: "Stepan",
	}}
	t.Log(person3)
}
