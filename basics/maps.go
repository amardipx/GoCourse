package basics

import "fmt"

func main() {
	// var mapVariable map[keyType]variableType
	// mapVariable = make(map[keyType]valueType)

	// using map literal
	// mapVariable = map[keyType][valueType] {
	// 	key1 : value1,
	// 	key2 : value2
	// }

	myMap := make(map[string]int)

	myMap["key1"] = 9
	myMap["key2"] = 10
	fmt.Println(myMap)
	// delete(myMap, "key1")
	// fmt.Println(myMap)
	// clear(myMap)
	// fmt.Println(myMap)

	value, unknownvalue := myMap["key1"]
	fmt.Println(value)
	fmt.Println(unknownvalue)

	_, ok := myMap["key2"]
	if ok {
		fmt.Println("A value exists for key2")
	} else {
		fmt.Println("No value exists for key2")
	}

	myMap2 := map[string]int{
		"a": 1,
		"b": 2,
	}

	fmt.Println(myMap2)

	for k, v := range myMap2 {
		fmt.Println(k, v)
	}

	// nested maps
	myMap3 := make(map[string]map[string]int)
	myMap3["map1"] = myMap2
	fmt.Println(myMap3)
}
