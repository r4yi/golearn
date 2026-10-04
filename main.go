package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}

	names := [4]string{"yoshi", "mario", "peach", "bowo"}
	names[1] = "gibran"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood, flexible)
	var scores = []int{100, 50, 60}
	scores[2] = 67

	fmt.Println(scores)

	scores = append(scores, 85)

	fmt.Println(scores)

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo:= names[1:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)

}
