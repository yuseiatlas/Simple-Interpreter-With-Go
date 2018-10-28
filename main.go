package main

import "fmt"

func main() {
	input := `let five = 5;
	let ten = 10;

	let add = fn(x, y) {
		x + y;
	let result = add(five, ten);
	};`
	fmt.Println("The result is: \n", input)
}
