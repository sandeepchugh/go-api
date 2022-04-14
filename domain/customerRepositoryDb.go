package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sandeepchugh/profileapi/errs"
	"log"
	"time"
)

type CustomerRepositoryDb struct {
	db *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {

	selectQuery := "SELECT customer_id,name,city,zipcode,date_of_birth, status FROM customers"

	rows, err := d.db.Query(selectQuery)
	if err != nil {
		log.Println("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateOfBirth, &customer.Status)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("customer not found")
			} else {
				log.Println("Error while scanning customers" + err.Error())
				return nil, errs.NewUnexpectedError("unexpected database error")
			}
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	selectQuery := "SELECT customer_id,name,city,zipcode,date_of_birth, status FROM customers WHERE customer_id = ?"

	row := d.db.QueryRow(selectQuery, id)
	var customer Customer
	err := row.Scan(&customer.Id, &customer.Name, &customer.City, &customer.Zipcode, &customer.DateOfBirth, &customer.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("Error while scanning customer with id:  " + id + ". " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &customer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	db, err := sql.Open("mysql", "root:rootpass@/banking")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(10)

	return CustomerRepositoryDb{db}
}
