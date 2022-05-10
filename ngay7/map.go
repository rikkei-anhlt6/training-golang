package main

import "fmt"

type User struct {
	name string
	age  int
}

// make(map[type of key]type of value)
func main() {
	map1 := make(map[string]int)
	map1["key1"] = 1
	map1["key2"] = 2
	fmt.Println(map1)
	map2 := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Println(map2)
	var map3 map[string]int //zero value = nil
	//map3["key1"] = 1        // ko thể gán cho nil map
	fmt.Println(map3)

	fmt.Println(map1["key1"]) //Retrieving value for a key from a map

	//Checking if a key exists
	value, ok := map1["key3"]
	if ok {
		fmt.Println(value)
	}
	fmt.Println("key3 not found")

	for key, value := range map1 {
		fmt.Println(key, value)
	}
	delete(map1, "key2") // delete(tên map, tên key)

	//Map of structs
	user1 := User{
		name: "lta1",
		age:  20,
	}
	user2 := User{
		name: "lta2",
		age:  30,
	}
	map4 := map[string]User{
		"user1": user1,
		"user2": user2,
	}
	fmt.Println(map4)
	fmt.Println(len(map4)) //length of the map
	//Maps are reference types
}
