package main

import "fmt"

var w int = 42

func add(x, y int) int {
	return x + y
	// x and y are parameters ( variables inside a function are called parameters and its values are arguments)
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	a := 42 // here 'a' is short declaration variable and its scope is just insisde this func. anywhere else this will error out

	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"

	fmt.Println(b, c, d, f)

	// all the variables defined should be used, otherwise it'll throw error except for blank(_) variable

	fmt.Printf("%X\n", a)
  fmt.Printf("%b\n", a)


  fmt.Printf("%v of type %T\n", w, w)

	// functions have parenthesis at the end to execute eg: func main() where as constant does not have them eg: Pi

	fmt.Println(add(5, 13))

	fmt.Println(swap("hello", "world"))

	fmt.Println(split(17))
}