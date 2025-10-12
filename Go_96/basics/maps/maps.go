package main

import (
	"fmt"
	"maps"
)

func main() {

	// var m map[keyType]valueType
	// mapVariable = make(map[keyType]valueType)
	// mapVariable =  map[keyType]valueType{key1: value1, key2: value2}

	myMap := make(map[string]int)
	fmt.Println("Initial map:", myMap)

	myMap["key1"] = 9
	myMap["key2"] = 18

	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"])

	myMap["code"] = 21
	fmt.Println("Map after adding code:", myMap)

	delete(myMap, "key1")
	fmt.Println("Map after deleting key2:", myMap)

	myMap["key3"] = 30
	myMap["key4"] = 40
	myMap["key5"] = 50
	fmt.Println("Map before deleting:", myMap)

	// clear(myMap)
	// fmt.Println("Map after clearing:", myMap)

	_, unknownvalue := myMap["key2"]
	// fmt.Println("Value:", value)
	fmt.Println("Unknown Value:", unknownvalue)

	myMap2 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("MyMap2:", myMap2)

	myMap3 := map[string]int{"a": 1, "b": 2, "c": 3}

	if maps.Equal(myMap, myMap2) {
		fmt.Println("The maps are equal.")
	} else {
		fmt.Println("The maps are not equal.")
	}

	if maps.Equal(myMap2, myMap3) {
		fmt.Println("The maps are equal.")
	} else {
		fmt.Println("The maps are not equal.")
	}

	for k, v := range myMap {
		fmt.Println("Key:", k, "Value:", v)
	}

	_, ok := myMap["key3"]
	if ok {
		fmt.Println("Key3 exists in the map.")
	} else {
		fmt.Println("Key3 does not exist in the map.")
	}

	var myMap4 map[string]string
	if myMap4 == nil {
		fmt.Println("myMap4 is nil")
	} else {
		fmt.Println("myMap4 is not nil")
	}

	val := myMap4["key"]
	fmt.Println("Value:", val)

	// myMap4["key"] = "value"
	// fmt.Println("myMap4 after assignment:", myMap4)

	myMap4 = make(map[string]string)
	myMap4["key"] = "value"
	fmt.Println("myMap4 after make and assignment:", myMap4)

	fmt.Println("myMap4 length is: ", len(myMap4))

	myMap5 := make(map[string]map[string]string)
	myMap5["map1"] = myMap4

	fmt.Println("myMap5:", myMap5)

}
