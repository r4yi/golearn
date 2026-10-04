package main

import (
	"fmt"
	"sort"
	// "strings"
)

func main() {
	// greeting := "hello there friends!"

	// fmt.Println(strings.Contains(greeting, "hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))

	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "th'"))

	// fmt.Println(strings.Split(greeting, " there "))

	// fmt.Println("original value:", greeting)

	ages:= []int{45, 52, 56, 54, 23, 50, 55, 23}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 56)
	fmt.Println(index)


	names := []string{"yoshi", "mario", "peach", "boweser", "luigi"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "mario"))

}
