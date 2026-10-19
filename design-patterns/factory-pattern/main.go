package main

import (
	"errors"
	"factory/adidas"
	"factory/nike"
	"factory/utils"
	"fmt"
)

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	printShoeDetails(adidasFactory.MakeShoe())
	printShirtDetails(adidasFactory.MakeShirt())

	nikeFactory, _ := GetSportsFactory("nike")
	printShoeDetails(nikeFactory.MakeShoe())
	printShirtDetails(nikeFactory.MakeShirt())

}

func printShoeDetails(s utils.IShoe) {
	fmt.Printf("Logo: %v\n", s.GetLogo())
	fmt.Printf("Size: %v\n", s.GetSize())
}

func printShirtDetails(s utils.IShirt) {
	fmt.Printf("Logo: %v\n", s.GetLogo())
	fmt.Printf("Size: %v\n", s.GetSize())
}

func GetSportsFactory(brand string) (utils.ISportsFactory, error) {
	switch brand {
	case "adidas":
		return &adidas.Adidas{}, nil
	case "nike":
		return &nike.Nike{}, nil
	default:
		return nil, errors.New("wrong brand type passed")
	}
}
