package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func inference() {
	var name = "Pratyush Mani Manav"
	var age = 34

	fmt.Println(name)
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.ValueOf(name))
	fmt.Println(age)
	fmt.Println(reflect.TypeOf(age))
	fmt.Println(reflect.ValueOf(age))
	fmt.Printf("%s is %d years old \n", name, age)

}

func shorthand() {
	name := "Prasang manav"
	age := 36
	fmt.Println(name)
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.ValueOf(name))
	fmt.Println(age)
	fmt.Println(reflect.TypeOf(age))
	fmt.Println(reflect.ValueOf(age))
	fmt.Printf("%s is %d years old \n", name, age)

}

func UserInput() {
	var name string
	fmt.Print("Enter Your Name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello ", name)
}

func conditionals() {
	rand.Seed(time.Now().UnixNano())
	num := -5 + rand.Intn(10)
	if num > 0 {
		fmt.Println("The number is positive")
	} else if num == 0 {

		fmt.Println("The number is zero")
	} else {

		fmt.Println("The number is negative")
	}
}
