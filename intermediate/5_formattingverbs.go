package intermediate

import "fmt"

func main() {
	//--General Formatting Verbs
	// %v Prints the value in teh default format
	// %#v Prints the value in Go-Syntax format
	//%T Prints the type of the value
	// %% Prints the % sign

	//Underscores are allowed in numeric literals in Go to improve readability,
	//and they can be ignored by the compiler
	//It is to make large numbers easy to read
	i := 1_03_5.5
	string := "Hello World!"

	fmt.Printf("%v\n", i)
	//difference will be noticed in string
	fmt.Printf("%#v\n", i)
	fmt.Printf("%T\n", i)
	//% will be printed in end
	fmt.Printf("%v%%\n", i)

	fmt.Println("----String--------")
	fmt.Printf("%v\n", string)
	//getting string in double quote which is in go syntax
	fmt.Printf("%#v\n", string)
	fmt.Printf("%Tv\n", string)

	// --- Integer Formatting Verbs
	// %b Base 2
	// %d Base 10
	// %+d Base 10 and always show sign
	// %o Base 8
	// %O Base 8, with leading 0o
	// %x Base 16, lowercase
	// %X Bae 16, uppercase, noticeable in large values with FF
	// %#x Base 16, with leading 0x
	// %4d Pad with spaces (width 4, right justified)
	// %-4d Pad with spaces (width 4, left justified)
	// %04d Pad with zeroes

	int := 255

	fmt.Println("=========Integer formatting===========")
	fmt.Printf("%b\n", int)
	fmt.Printf("%d\n", int)
	fmt.Printf("%+d\n", int)
	fmt.Printf("%o\n", int)
	fmt.Printf("%O\n", int)
	fmt.Printf("%x\n", int)
	fmt.Printf("%X\n", int)
	fmt.Printf("%#x\n", int)
	fmt.Printf("%4d\n", int)
	fmt.Printf("%-4d\n", int)
	fmt.Printf("%04d\n", int)

	//  -- String Formatting Verbs
	// %s Prints the value as plain string
	// %q Prints the value as a double-quoted string
	// %8s Prints the value in plain string (width 8, right justified)
	// %-8s Prints the value as plain string (width 8, left justified)
	// %x Prints the value as hex dump of byte values
	// % x Prints the value as hex dump with spaces

	fmt.Println("=========Strings=============")
	txt := "World"
	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	// -- Boolean Formatting Verbs
	// %t Value of the boolean operator in true or false format (same as using %v)

	t := true
	f := false

	fmt.Println("======Boolean=======")
	fmt.Printf("%t\n", t)
	fmt.Printf("%t\n", f)
	fmt.Printf("%v\n", t)

	// --Floating Formatting Verbs
	// Verb Description
	// %e Scientific notation with 'e' as exponent
	// %f Decimal point, no exponent
	// %.2f Default width, precision 2
	// %6.2f width 6, precision 2
	// %g Exponent as needed , only necessary digits

	flt := 9180000.00

	fmt.Println("==========Floating=========")
	fmt.Printf("%e\n", flt)
	fmt.Printf("%f\n", flt)
	fmt.Printf("%.2f\n", flt)
	//Minimum number of total characters(including the '.' decimal)
	fmt.Printf("%6.2f\n", flt)
	fmt.Printf("%g\n", flt)

}
