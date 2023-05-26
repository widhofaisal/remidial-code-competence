package model


type Item struct {
	ID          string `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Stock       int    `json:"stock" form:"stock"`
	Price       int    `json:"price" form:"price"`
	CategoryID  int    `json:"category_id" form:"category_id"`
}

// type Response struct {
// 	Name        string `json:"name" form:"name"`
// 	Description string `json:"description" form:"description"`
// 	Stock       int    `json:"stock" form:"stock"`
// 	Price       int    `json:"price" form:"price"`
// 	CategoryID  int    `json:"category_id" form:"category_id"`
// }
