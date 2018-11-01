package main

import ("fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")

	str1 := "string"
	str2 := "string2"
	str3 := "string3"
	stringLength, _ := fmt.Println(str1, str2, str3)
 	    fmt.Println(stringLength)

	const nonvar int = 1

	var aString string = "ffdf"
	fmt.Println(aString)
	fmt.Println(strings.EqualFold("ddd", "DDD"))

	var p *int
	i := 5
	p = &i
	fmt.Println("value of p:", *p)

	var colors[3] string
	colors[0] = "Red"
	colors[1] = "Black"
	colors[2] = "Blue"
	fmt.Println(colors)
}
