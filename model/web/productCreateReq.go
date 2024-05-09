package web

type ProductCreateReq struct {
	Id          int
	Name        string `validate:"required" json:"name"`
	Sku         string `validate:"required" json:"sku"`
	Category    string `validate:"required" json:"category"`
	ImageUrl    string `validate:"required" json:"imageUrl"`
	Notes       string `validate:"required" json:"notes"`
	Price       int    `validate:"required" json:"price"`
	Stock       int    `validate:"required" json:"stock"`
	Location    string `validate:"required" json:"location"`
	IsAvailable bool   `validate:"required" json:"isAvailable"`
}
