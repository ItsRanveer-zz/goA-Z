// go7.go
package main

import (
	"fmt"
	"math"
)

func main() {
	//Structs (Mutable. Allows you to define your own type with collection of fields or properties and also define method of that type)
	rect1 := Rectangle{topY: 0, leftX: 50, width: 10, height: 5}  //Creating a struct instance. You can also give as Rectangle{0,50,10,5}, Omitted field will be 0
	fmt.Printf("Type of Instance: %T \n", rect1)                  //Type of rect1
	fmt.Printf("Instance of Rectangle: %v \n", rect1)             //%v is a printing verb used to print the struct's instance
	fmt.Printf("Instance of Rectangle with field: %+v \n", rect1) //%+v will include field name's also
	fmt.Printf("Syntax representation:  %#v \n", rect1)           //%#v will print the Syntax represantation
	fmt.Println("Rectangle width is = ", rect1.width, " wide")
	fmt.Println("Area of rectangle is = ", rect1.area())                         //Calling struct method
	rect2 := &rect1                                                              //Creating a pointer to the previous struct
	rect2.width = 11                                                             //Any change to this struct pointer will change the value of struct coz this is pointing to that.
	fmt.Println("Actual Struct - ", rect1, ", Pointer to that struct - ", rect2) //Printing the struct and the pointer

	//Interface (To achive Polymorphism)
	square := Square{20, 20}                       //Creating instance of Square
	circle := Circle{4}                            //Creating instance of Circle
	fmt.Println("Square Area = ", square.area())   // Directly calling area() through square type, so area implemented by Square type will be called
	fmt.Println("Square Area = ", getArea(square)) // Square and Circle both implement Shape, so there instances can be used as the argument to the getArea and getPerimeter
	fmt.Println("Square Perimeter = ", getPerimeter(square))
	fmt.Println("Circle Area = ", getArea(circle))
	fmt.Println("Circle Perimeter = ", getPerimeter(circle))

}

//Struct (Defining a struct)
type Rectangle struct {
	leftX, topY, height, width float64 //Defining attributes to struct
}

//Defining method of struct Rectangle with a pointer receiver type
func (rect *Rectangle) area() float64 { // Method of type Rectangle(Its own method, not implemented one)
	return rect.width * rect.height
}

//Interfaces are named collections of method signatures only. (Declare a method and implement it in struct which uses that interface)
type Shape interface {
	area() float64
	perimeter() float64
}

//Structs which will implement the interface
type Square struct {
	height float64
	width  float64
}
type Circle struct {
	radius float64
}

//To implement an interface, we just need to implement all the methods in the interface. Here we implement Shape on Square(All method)
func (s Square) area() float64 {
	return s.height * s.width
}
func (s Square) perimeter() float64 {
	return 2*s.height + 2*s.width
}

//Here we implement Shape on Circle(All method)
func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

//A generic functions taking type Shape(interface) which can be a Circle or Rectangle
func getArea(shape Shape) float64 {
	return shape.area()
}
func getPerimeter(shape Shape) float64 {
	return shape.perimeter()
}
