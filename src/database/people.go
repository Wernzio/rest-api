package Database

import "errors"

var people = []Person{
	Person{ID: 1, Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}},
	Person{ID: 2, Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}},
	Person{ID: 3, Firstname: "Francis", Lastname: "Sunday"},
}
var index = len(people)

func GetPeople() *[]Person {
	return &people
}

func GetPerson(id int) (*Person, error) {
	for _, item := range people {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, errors.New("Entity not found")
}

func CreatePerson(person *Person) *Person {
	index++
	person.ID = index
	people = append(people, *person)
	return person
}

func DeletePerson(id int) bool {
	deleted := false
	for index, item := range people {
		if item.ID == id {
			people = append(people[:index], people[index+1:]...)
			deleted = true
			break
		}
	}
	return deleted
}
