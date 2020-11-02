package main

import (
	"fmt"

	"github.com/goSql/config"
	"github.com/goSql/entities"
	"github.com/goSql/models"
	"github.com/gofrs/uuid"
)

func main() {
	test_CallFindAll()
	//test_CallFind()
	//test_Insert()
}

func test_Insert() {
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		uuid, _ := uuid.NewV4()
		id := uuid.String()

		product := entities.Product{
			Id:       id,
			Name:     "fridge",
			Quantity: 2,
			Price:    45000.95,
		}
		productModel.Create(&product)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("updated", product.Id)
		}

	}
}

func test_CallFind() {
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product, err := productModel.Find("b969a8e7-bc68-4fc4-87d4-6c8c24")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(product)
			fmt.Println("product List")

			fmt.Println("Id", product.Id)
			fmt.Println("Name", product.Name)
			fmt.Println("Quantity", product.Quantity)
			fmt.Println("Price", product.Price)
			fmt.Println("----------------------------------")

		}
	}
}

func test_CallFindAll() {
	db, err := config.GetMySQLDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err := productModel.FindAll()

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(products)
			fmt.Println("product List")
			for _, product := range products {
				fmt.Println("Id", product.Id)
				fmt.Println("Name", product.Name)
				fmt.Println("Quantity", product.Quantity)
				fmt.Println("Price", product.Price)
				fmt.Println("----------------------------------")
			}
		}
	}
}
