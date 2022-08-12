package models

// Product is a model for a product
type Product struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"unique;not null" example:"baseball"`
	Stock int32  `json:"stock,omitempty" gorm:"not null; default:null" example:"15"`
	Price int32  `json:"price,omitempty" gorm:"not null; default:null" example:"4"`
}
