package nike

import "factory/utils"

type Nike struct{}

func (n *Nike) MakeShoe() utils.IShoe {
	return &NikeShoe{
		Shoe: utils.Shoe{
			Logo: "nike",
			Size: 7,
		},
	}
}

func (n *Nike) MakeShirt() utils.IShirt {
	return &NikeShirt{
		Shirt: utils.Shirt{
			Logo: "nike",
			Size: 38,
		},
	}
}
