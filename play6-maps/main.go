package main

import "fmt"

func main() {
	numbers := map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println(numbers)

	numbers = make(map[int]string) // to initialize an empty map
	numbers[4] = "four"
	numbers[5] = "five"
	fmt.Println(numbers)

	delete(numbers, 5)
	fmt.Println(numbers)

	colors := map[string]string{"red": "0xffff", "blue": "0x44353"}
	printMap(colors)

}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}
