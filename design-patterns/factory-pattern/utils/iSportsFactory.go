package utils

type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}
