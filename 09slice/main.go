package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitlist = []string{"Apple", "Tomato", "Banana"}
	fmt.Printf("Type of fruitlist is %T\n", fruitlist)
	fruitlist = append(fruitlist, "Mango", "Peach")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:])
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:3])
	fmt.Println(fruitlist)

	highscore := make([]int, 4)
	highscore[0] = 234
	highscore[1] = 945
	highscore[2] = 465
	highscore[3] = 867
	fmt.Println(highscore)

	highscore = append(highscore, 555, 666, 321)
	fmt.Println(highscore)

	fmt.Println(sort.IntsAreSorted(highscore))

	sort.Ints(highscore)
	fmt.Println(highscore)
	fmt.Println(sort.IntsAreSorted(highscore))

	var courses = []string{"Java", "Python", "C++", "JavaScript"}
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
