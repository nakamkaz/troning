package main

import (
	"fmt"
)

/** Nand impl with no use !And() **/
func Nand(x1, x2 bool) bool {
	retval := false
	if (x1 == true) && (x2 == true) {
		retval = false
	} else {
		retval = true
	}
	return retval
}

/** And implemented by Nand */
func And(x1, x2 bool) bool {
	// goal:  (x1 && x2)
	return !Nand(x1,x2)
}

/** Or impl by Nand */
func Or(x1, x2 bool) bool {
	// goal:  (x1 || x2)
	retval :=false
	return Nand(x1,x2)
}

/** Xor impl by Nand */
func Xor(x1, x2 bool) bool {
	// goal:  (x1 != x2)

	return (x1 != x2)
}

func Test1() {
	fmt.Println("And 0,0", And(false, false))
	fmt.Println("And 1,0", And(true, false))
	fmt.Println("And 0,1", And(false, true))
	fmt.Println("And 1,1", And(true, true))

	fmt.Println("Or 0,0", Or(false, false))
	fmt.Println("Or 1,0", Or(true, false))
	fmt.Println("Or 0,1", Or(false, true))
	fmt.Println("Or 1,1", Or(true, true))

	fmt.Println("Xor 0,0", Xor(false, false))
	fmt.Println("Xor 1,0", Xor(true, false))
	fmt.Println("Xor 0,1", Xor(false, true))
	fmt.Println("Xor 1,1", Xor(true, true))

	fmt.Println("Nand 0,0", Nand(false, false))
	fmt.Println("Nand 1,0", Nand(true, false))
	fmt.Println("Nand 0,1", Nand(false, true))
	fmt.Println("Nand 1,1", Nand(true, true))
}
func main() {
	fmt.Println("And/Or/Xor implemented by Nand")
	Test1()
}
