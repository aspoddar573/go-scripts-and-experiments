package helloWorld

import (
	"fmt"
	"math/rand"
)

func GFG(i func(p, q string) string) {
	fmt.Println(i("Geeks", "for"))
}

func kchBhi() {
	fmt.Println("Hello World")
	fmt.Println("It works like a dream!")
	fmt.Println("The number is ", rand.Intn(10), ". Okay?")

	value := func(p, q string) string {
		return p + q + "Geeks"
	}
	GFG(value)

	var temp = 65
	var p *int = &temp

	fmt.Println(temp)
	fmt.Println(*p)

	temp = 47

	fmt.Println(temp)
	fmt.Println(*p)

	*p = 368

	fmt.Println(temp)
	fmt.Println(*p)

	value2 := 25
	ptr2 := &value2
	fmt.Println(ptr2)
	fmt.Println(*ptr2)

}

func HelloBolo() {
	kchBhi()
}
