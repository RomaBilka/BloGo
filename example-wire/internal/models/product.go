package models

type ProductId int

type Product struct {
	Id          ProductId
	Name        string
	Description string
	Image       string
	Price       int
	Count       int
}
