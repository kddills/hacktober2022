package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	if len(os.Args) > 1 {
		name = os.Args[1]
	} else {
		name = "Jin"
	}

	fmt.Println(HelloWorld(name))
}

// Hello world project!
func HelloWorld(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
