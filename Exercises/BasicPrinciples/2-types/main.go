package main

import "fmt"

func main() {
	// int
	var age int = 30

	// float
	var height float32 = 1.76

	// bolean
	var ofLegalAge bool = age >= 18
	var tallPerson bool = height >= 1.80

	fmt.Println(age)
	fmt.Println(height)
	fmt.Println(ofLegalAge)
	fmt.Println(tallPerson)

	fmt.Printf("Tipo de ofLegalAge: %T\n", ofLegalAge)
}
