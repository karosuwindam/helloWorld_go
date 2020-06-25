package helloworld

import "fmt"

func Puts(i int) string {
	var output string
	switch i {
	case 0:
		output = "Hello World"
	case 1:
		output = "Hello World !!"
	default:
		output = "Hello World !"
	}
	fmt.Println(output)
	return output
}
