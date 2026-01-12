package main

import "fmt"

func main() {
	var name string
	var age int
	var height float32

	fmt.Printf("Your name: ")
	fmt.Scan(&name)

	fmt.Printf("Your age: ")
	fmt.Scan(&age)

	fmt.Printf("Your height (in meters): ")
	fmt.Scan(&height)

	// Calcula após ler os valores
	var ofLegalAge bool = age >= 18
	var tallPerson bool = height >= 1.80

	fmt.Println("\n--- Results ---")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %.2f m\n", height)
	fmt.Printf("Of legal age: %t\n", ofLegalAge)
	fmt.Printf("Tall person: %t\n", tallPerson)
}
