package models

import "store/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func GetAllProducts() []Product {
	db := db.ConnectDB()
	resultQuery, err := db.Query("select * from products order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for resultQuery.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = resultQuery.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}

	defer db.Close()
	return products
}

func CreateProduct(name, description string, price float64, amount int) {
	db := db.ConnectDB()

	scriptInsert, err := db.Prepare("insert into products(name, description, price, amount) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	scriptInsert.Exec(name, description, price, amount)
	defer db.Close()
}

func UpdateProduct(id int, name, description string, price float64, amount int) {
	db := db.ConnectDB()

	scriptUpdate, err := db.Prepare("update products set name=$2, description=$3, price=$4, amount=$5 where id=$1")
	if err != nil {
		panic(err.Error())
	}

	scriptUpdate.Exec(id, name, description, price, amount)
	defer db.Close()
}

func DeleteProcuct(id string) {
	db := db.ConnectDB()

	scriptDelete, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}
	scriptDelete.Exec(id)
	defer db.Close()
}

func GetProduct(id string) Product {
	db := db.ConnectDB()
	result, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for result.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = result.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount
	}
	defer db.Close()
	return product
}
