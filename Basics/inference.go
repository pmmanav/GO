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

func for_loop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)
	for_loop2()
}

func for_loop2() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)
}
