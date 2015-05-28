package main

import "fmt"

func main() {
	var i int // declares a variable, later assign with i = 0
	i = 0

	j := 0 // declares & assigns the variable, defines type based on how variable is used

	s := make([]int, 10) // creates "slice" of 10 integers

	s[1] = 1
	s[2] = 2

	s2 := []int{1, 2, 3, 4} // creates slice of [1,2,3,4]
	s3 := s2[1:3]
	s2 = append(s2, 5)

	s3 = append(s3, s2...) // need to do s2..., otherwise, will try to add array as 1 object.  Like *
	fmt.Println(s[1], s[2], i, j, s2[4], s3[1])
	fmt.Println(s3)

	m := make(map[string]int) // hash[keytype] valuetype
	m["foo"] = 1
	m["bar"] = 2

	fmt.Println(m["foo"])

	m2 := map[string]int{
		"foo": 1,
		"bar": 2}

	fmt.Println(m2["bar"])

	value, ok := m2["foo"]
	if !ok {
		m2["foo"] = 9
	}

	fmt.Println(value, ok, "wooooo")
	fmt.Printf("the value is %d\n", value)

	type Foo struct {
		value int
	}
	type Bar struct {
		Lorem int
		ipsum string
		bleh  rune
		blah  *Foo
	}
	f := Foo{8}

	b := Bar{5, "hello", 't', &f}
	b2 := new(Bar)
	b2.Lorem = 100

	fmt.Println(f)
	fmt.Println(b)
	fmt.Println(b.Lorem)
	fmt.Println(b2)

	b4 := new(Bar)   // creates a pointer to bar
	b5 := Bar{}      // creates actual bar
	b6 := Bar{Lorem: 0,
		ipsum: "hello"} // creates actual bar

	fmt.Println(b4, b5, b6)

	b7 := b6
	b7.Lorem = 222

	fmt.Println(b6, b7)

}
