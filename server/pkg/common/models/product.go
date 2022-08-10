package models

// Product is a model for a product
type Product struct {
	Id    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"unique_index;not null" example:"baseball"`
	Stock int32  `json:"stock" gorm:"not null" example:"15"`
	Price int32  `json:"price" gorm:"not null" example:"4"`
}