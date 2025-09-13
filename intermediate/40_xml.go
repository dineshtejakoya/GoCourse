package intermediate

import (
	"encoding/xml"
	"fmt"
	"log"
)

//extensible markup language used for encoding documents in a format that is both human and machine readable
//it is widely used for data interchange b/w system and configuration files
//go provides encodign/xml package to encode and decode data
//before json with rest api, it was xml the first choice for transferring data

type Person struct {
	//XMLName is a special field used by encoding/xml package, it is used to determine the name of xml element for marshalling and unmarshalling data
	//it is used to specify xml element name of the struct
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age,omitempty"`
	Email   string   `xml:"-"`
	//- to omit data permanently place - in metadata
	//City    string   `city:"city"`
	Address Address `xml:"address"`
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main() {

	// person := Person{Name: "John", City: "London", Email: "emiail@example.com"}
	person := Person{Name: "John", Email: "email@example.com", Address: Address{City: "Oakland", State: "CA"}}
	xmlData, err := xml.Marshal(person)
	if err != nil {
		log.Fatalln("Error Marshalling data into XML:", err)
	}

	//response we can see all in one line here
	fmt.Println(string(xmlData))

	//the last parameter would be indent spaces , it can be tab(\t) depending up on our choice
	xmlData1, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatalln("Error Marshalling data into XML:", err)
	}
	//XML Data with indent
	fmt.Println(string(xmlData1))

	//xmlRaw := `<person><name>Jane</name><age>25</age></person>`"
	xmlRaw := `<person><name>John</name><age>25</age><address><city>"SanFrancisco"</city><state>CA</state></address></person>`
	var personxml Person
	err = xml.Unmarshal([]byte(xmlRaw), &personxml)
	if err != nil {
		log.Fatalln("Error unmarshalling XML:", err)
	}
	fmt.Println(personxml)

	//by default any  root element we have will be stored inside local string
	fmt.Println("Local String", personxml.XMLName.Local)
	fmt.Println("Namespace", person.XMLName.Space)

	fmt.Println("========Book=========")
	book := Book{
		ISBN:       "583-9293-88292",
		Title:      "Go Bootcamp",
		Author:     "Ashish",
		Pseudo:     "Pseudo",
		PseudoAttr: "Pseudo Attribute",
	}

	//^ is a prefix we added to check whether it is coming in every line or not
	xmlDataAttr, err := xml.MarshalIndent(book, "^", "  ")
	if err != nil {
		log.Fatalln("Error Marshalling data:", err)
	}

	fmt.Println(string(xmlDataAttr))
}

// Assigning attributes to any element
type Book struct {
	XMLName xml.Name `xml:"book"`
	//adding attr will let the xml know this not a child element but attribute of the parent element
	ISBN   string `xml:"isbn,attr"`
	Title  string `xml:"title,attr"`
	Author string `xml:"author,attr"`
	//if we want to create attribute values to child elements as well, we need to create separate struct for child elements as well
	//because in the same struct all attributes that you defined are going to be assigned to the primary element of that struct
	Pseudo     string `xml:"pseudo"`
	PseudoAttr string `xml:"pseudoattr,attr"`
}
