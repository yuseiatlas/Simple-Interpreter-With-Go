package main

import (
	"fmt"
	"github.com/yuseiatlas/interpreter/repl"
	"os"
	"os/user"
)

func main() {

	/* User dynamic input */
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)

	/* Static input */
	//input := `let five = 5;
	//		let ten = 10;
	//		let add = fun(x, y) {
	//			x + y;
	//		};
	//		let result = add(five, ten);
	//		!-/*5;
	//		5 < 10 > 5;
	//		if (5 < 10) {
	//			return true;
	//		} else if (5 == 5) {
	//			return true;
	//		} else if(5 != 100) {
	//			return false;
	//		} else {
	//			return false;
	//		}
	//		while (5 < hello) {
	//			return false;
	//		}
	//		struct myStruct {
	//			a: string
	//			b: double
	//		}`
	//repl.PrintResult(input)
}
