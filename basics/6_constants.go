package basics

//untyped constant
const pi = 3.14
const GRAVITY = 9.81

func main() {
	//typed constant
	const days int = 7

	//we can declare multiple constants
	const (
		monday        = 1
		tuesday       = 2
		wednesday     = 3
		thursday  int = 4
	)

	//constants must be initialized with values that can be determined at compile time
	//no short declartion (:=) for constants

}
