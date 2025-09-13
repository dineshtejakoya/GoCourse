package intermediate

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("math error: square root of negative number")
	}
	//compute the square root
	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error empty data")
	}
	return nil
}

func main() {
	// result, err := sqrt(16)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result)

	// result1, err1 := sqrt(-16)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }
	// fmt.Println(result1)

	// data := []byte{}
	// if err2 := process(data); err2 != nil {
	// 	fmt.Println("Error:", err2)
	// 	return
	// }
	// fmt.Println("Data Processed successfully")

	// err3 := eprocess()
	// if err3 != nil {
	// 	fmt.Println("Error:", err3)
	// 	return
	// }

	if err := readData(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data read successfully")
}

// custom errors
type myError struct {
	message string
}

// go's builtin package has Error() interface
// by using this we can format custom interfaces as we please
func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func eprocess() error {
	return &myError{"custom error message"}
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

func readConfig() error {
	return errors.New("Config Error")
}
