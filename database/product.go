package database

import (
	"log"

	"github.com/KhizarShabir1/foodji-tinder/foodji"
)

var _ foodji.ProductProvider = (*Provider)(nil)

func (p *Provider) ListProducts() ([]foodji.Product, int, error) {
	var results []foodji.Product
	rows, err := p.db.Query("SELECT id, name FROM product")
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var prod foodji.Product
		err = rows.Scan(&prod.ID, &prod.Name)
		if err != nil {
			return nil, 0, err
		}
		results = append(results, prod)
	}

	err = rows.Err()
	if err != nil {
		return nil, 0, err
	}

	return results, len(results), nil
}

func (p *Provider) GetProduct(id string) (*foodji.Product, error) {
	var prod foodji.Product

	// Prepare the SQL statement to select the product by ID
	stmt, err := p.db.Prepare("SELECT id, name FROM product WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the statement with the product ID parameter
	err = stmt.QueryRow(id).Scan(&prod.ID, &prod.Name)
	if err != nil {
		return nil, err
	}

	return &prod, nil
}

func (p *Provider) CreateProduct(product foodji.Product) (foodji.Product, error) {
	log.Println(product)

	_, err := p.db.Exec("INSERT INTO product (id, name) VALUES ($1, $2)", product.ID, product.Name)
	if err != nil {
		log.Println("panic created while creating product")
		panic(err)
	}

	log.Println("Product created successfully")
	return product, nil
}
