package intermediate

import "fmt"

func main() {
	// sequence := adder()

	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())

	// sequence2 := adder()
	// fmt.Println(sequence2())

	adderMain := func() func(int) int { // declaring an anonymous function which returnes a function which takes an integer argument and returns an integer.
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}()

	// using the closure adderMain
	fmt.Println(adderMain(10))
	fmt.Println(adderMain(11))
	fmt.Println(adderMain(12))
	fmt.Println(adderMain(13))
}

// func adder() func() int {
// 	i := 0
// 	fmt.Println("previous value of i:", i)
// 	return func() int {
// 		i++
// 		fmt.Println("added 1 to i")
// 		return i
// 	}
// }
