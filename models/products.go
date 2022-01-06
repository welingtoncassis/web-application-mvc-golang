package models

import "store/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func GetAll() []Product {
	db := db.ConnectDB()
	resultQuery, err := db.Query("select * from products")
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

func Create(name, description string, price float64, amount int) {
	db := db.ConnectDB()

	scriptInsert, err := db.Prepare("insert into products(name, description, price, amount) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	scriptInsert.Exec(name, description, price, amount)
	defer db.Close()
}
