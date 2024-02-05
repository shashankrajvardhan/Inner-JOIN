package main

import (
	"fmt"
	"log"
)

func join() {
	query := `
	SELECT orders.orderid, customer.customerid, customer.customername, orders.orderdate
	FROM orders
	INNER JOIN customer ON orders.customerid = customer.customerid;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			orderid, customerid     int
			orderdate, customername string
		)

		err := rows.Scan(&orderid, &customerid, &customername, &orderdate)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("orderid: %d, customerid: %d, customername: %s, orderdate: %s\n", orderid, customerid, customername, orderdate)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
