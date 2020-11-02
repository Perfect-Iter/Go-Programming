package models

import (
	"database/sql"

	"github.com/goSql/entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) Create(product *entities.Product) error {
	_, err := productModel.Db.Exec("INSERT INTO products(id, name, quantity, price) values(?,?,?,?)", product.Id, product.Name, product.Quantity, product.Price)
	if err != nil {
		return err
	}
	return nil

}

func (productModel ProductModel) Find(id string) (entities.Product, error) {
	rows, err := productModel.Db.Query("SELECT * FROM products WHERE id = ?")

	if err != nil {
		return entities.Product{}, err
	} else {
		product := entities.Product{}
		for rows.Next() {
			var id string
			var name string
			var quantity int
			var price float64

			err2 := rows.Scan(&id, &name, &quantity, &price)
			if err2 != nil {
				return entities.Product{}, err2
			} else {
				product = entities.Product{id, name, quantity, price}
			}
		}
		return product, nil
	}
}

func (productModel ProductModel) FindAll() ([]entities.Product, error) {
	rows, err := productModel.Db.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	} else {
		products := []entities.Product{}
		for rows.Next() {
			var id string
			var name string
			var quantity int
			var price float64

			err2 := rows.Scan(&id, &name, &quantity, &price)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{id, name, quantity, price}
				products = append(products, product)
			}
		}
		return products, nil
	}
}
