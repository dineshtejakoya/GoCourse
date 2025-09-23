package concurrency

import (
	"fmt"
	"sync"
)

type person struct {
	name string
	age  int
}

func main() {
	var pool = sync.Pool{
		//New field will create a new instance if the object pool is empty
		New: func() interface{} {
			fmt.Println("Creating a new Person. ")
			return &person{}
		},
	}

	//Get an object from the pool
	person1 := pool.Get().(*person)
	person1.name = "John"
	person1.age = 18
	fmt.Println("Got Person:", person1)

	fmt.Printf("Person1 - Name: %s, Age: %d\n", person1.name, person1.age)

	pool.Put(person1)
	fmt.Println("Returned Person to poool")

	person2 := pool.Get().(*person)
	fmt.Println("Got perso2:", person2)

	person3 := pool.Get().(*person)
	fmt.Println("Got person3:", person3)
	person3.name = "Jane"

	//Returning object to the pool again
	pool.Put(person2)
	pool.Put(person3)
	fmt.Println("Returned Persons to pool")

	person4 := pool.Get().(*person)
	fmt.Println("Got Person4: ", person4)

	person5 := pool.Get().(*person)
	fmt.Println("Got person5:", person5)
}
