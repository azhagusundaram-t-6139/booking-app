package adidas

import "factory/utils"

type Adidas struct{}

func (a *Adidas) MakeShoe() utils.IShoe {
	return &AdidasShoe{
		Shoe: utils.Shoe{
			Logo: "adidas",
			Size: 7,
		},
	}

}

func (a *Adidas) MakeShirt() utils.IShirt {
	return &AdidasShirt{
		Shirt: utils.Shirt{
			Logo: "adidas",
			Size: 38,
		},
	}
}
