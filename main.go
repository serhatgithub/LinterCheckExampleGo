package main

import "fmt"

func camel_Case() {
	x := true

	if x {
		return
	} else { // Gereksiz 'else'. Zaten yukarıda return var.
		fmt.Println("Lint bunu sevmez")
	}
}
