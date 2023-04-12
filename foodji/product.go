package foodji

// ProductProvider give access to products
type ProductProvider interface {
	ListProducts() ([]Product, int, error)
	GetProduct(id string) (*Product, error)
	CreateProduct(product Product) (Product, error)
}

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
