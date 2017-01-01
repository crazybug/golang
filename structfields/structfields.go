package main

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	speciality string
	Skills
}

func main() {
	mark := Student{Human: Human{"mark", 25, 120}, speciality: "Computer Science"}

	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)
	fmt.Println("Get Human weight:", mark.Human.weight)

	mark.speciality = "AI"
	fmt.Println("His speciality is ", mark.speciality)

	mark.Skills = []string{"anatomy"}
	fmt.Println("His skills are ", mark.Skills)
	mark.Skills = append(mark.Skills, "physics", "golang")
	fmt.Println("His skills now are ", mark.Skills)

}
