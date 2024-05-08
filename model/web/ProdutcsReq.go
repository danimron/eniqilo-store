package web

type ProductsReq struct {
	Name    	string `validate:"required" json:"name"`
	Sku 		string `validate:"required" json:"sku"`
	Category    string `validate:"required" json:"category"`
	imageUrl    string `validate:"required" json:"imageUrl"`
	Notes     	string `validate:"required" json:"notes"`
	Price     	string `validate:"required" json:"price"`
	Stock		string `validate:"required" json:"stock"`
	Location    string `validate:"required" json:"location"`
	IsAvailable bool `validate:"required" json:"isAvailable"`
}
