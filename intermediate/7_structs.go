package intermediate

import "fmt"

//Structs in Go are composite data types that allow you to group together different types of variables under a single name
//similar to classes in object oriented languages, more light weight and do not support inheritance

//usually structs would be made outside main function
//in go you cannot define a struct and its methods directly inside a main function
type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	//embedded anonymous field
	PhoneHomeCell
}

type Address struct {
	city    string
	country string
}

type PhoneHomeCell struct {
	home string
	cell string
}

func main() {

	//structs can be initialized using struct literal
	//if we omit any field in a struct it will be initialized with a zero value
	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city:    "London",
			country: "U.K",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "3434343423",
			cell: "3433084093",
		},
	}

	//Anonymous field in action
	//they promote the embedded fields to the outer struct
	//here we have direct access to cell
	fmt.Println(p.cell)
	//whereas in case city we need to use normally
	fmt.Println(p.address.city)

	p1 := Person{
		firstName: "Dinesh Teja",
		age:       25,
	}

	p2 := Person{
		firstName: "Dinesh Teja",
		age:       25,
	}
	fmt.Println("Are p1 and p2 equal:", p2 == p1)
	// p1.address.city = "Guntur"
	// p1.address.country = "Bharat"

	fmt.Println(p.firstName)
	fmt.Println(p1.firstName)
	fmt.Println(p.fullName())
	fmt.Println(p.address.country)
	fmt.Println(p.address)
	fmt.Println(p1.address.country)
	//Anonymous Structs
	// we pass the field values immediately
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "pseudoemail@example.com.org",
	}
	fmt.Println(user.username)
	fmt.Println("Before Increment:", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("After Increment:", p1.age)

	fmt.Println("=====Comparing two structs=========")
	fmt.Println("Are p1 and p equal:", p == p1)
	fmt.Println("Are p1 and p2 equal:", p2 == p1)

}

//p is an instance of struct Person we are accessing firstname and lastname of it
//receiver is associated to struct person
func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

//creating a method which is going to increase age by 1
func (p *Person) incrementAgeByOne() {
	p.age++
}
