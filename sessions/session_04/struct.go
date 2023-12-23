package main

import "fmt"

type Person struct {
	Name string
	Age  uint
}

func (p *Person) GetAge() uint {
	return p.Age
}

func (p *Person) SetAge(age uint) {
	p.Age = age
}

func main() {
	p := Person{
		Name: "Hesam",
		Age:  27,
	}

	//lc := struct {
	//	Name  string
	//	Model string
	//}{Name: "TOYOTA", Model: "2020"}

	//fmt.Println("P is: ", p)
	//fmt.Println("persons age is", p.Age)
	//p.Age = 33
	//fmt.Println("persons age is", p.Age)

	p.GetAge()
	fmt.Println("persons age by method is", p.GetAge())

	p.SetAge(40)
	fmt.Println("persons age by method is", p.GetAge())

}
