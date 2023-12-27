package seeders

import (
	"fmt"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ProductSeed(count int64) []Product {
	products := []Product{}

	for i := 0; i <= int(count); i++ {
		product := Product{
			ID:   i,
			Name: fmt.Sprintf("user%d", i),
		}
		products = append(products, product)
	}
	return products

}
