package intermediate

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Error. square root of a negative number doesn't exist")
	}

	// calculate sqrt. not doing that here.
	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error: Empty data")
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

	// result1, err1 := sqrt(-2)

	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	// fmt.Println(result1)

	// data := []byte{}
	// err := process(data)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("Data processed successfully")

	// err1 := eprocess()
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }

	err := readData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("data read successfully")
}

// Custom error
type myError struct {
	message string
}

func (m *myError) Error() string {
	return fmt.Sprintf("Error %s", m.message)
}

func eprocess() error {
	return &myError{"Custom error generated"}
}

// Using errorf

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	} else {
		return nil
	}
}

func readConfig() error {
	return errors.New("config error")
}
