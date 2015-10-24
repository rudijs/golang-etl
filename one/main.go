package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	orders := extract()
	orders = transform(orders)
	load(orders)
	fmt.Println(time.Since(start))
}

// Product from csv
type Product struct {
	PartNumber string
	UnitCost   float64
	UnitPrice  float64
}

// Order from csv plus Product info
type Order struct {
	CustomerNumber int
	PartNumber     string
	Quantity       int

	UnitCost  float64
	UnitPrice float64
}

func extract() []*Order {
	result := []*Order{}

	f, _ := os.Open("./../orders.txt")
	defer f.Close()
	r := csv.NewReader(f)

	for record, err := r.Read(); err == nil; record, err = r.Read() {
		order := new(Order)
		order.CustomerNumber, _ = strconv.Atoi(record[0])
		order.PartNumber = record[1]
		order.Quantity, _ = strconv.Atoi(record[2])
		result = append(result, order)
	}

	return result

}

func transform(orders []*Order) []*Order {
	f, _ := os.Open("./../productList.txt")
	defer f.Close()
	r := csv.NewReader(f)

	records, _ := r.ReadAll()

	productList := make(map[string]*Product)

	for _, record := range records {
		product := new(Product)
		product.PartNumber = record[0]
		product.UnitCost, _ = strconv.ParseFloat(record[1], 64)
		product.UnitPrice, _ = strconv.ParseFloat(record[2], 64)
		productList[product.PartNumber] = product
	}

	for idx := range orders {
		time.Sleep(3 * time.Millisecond)
		o := orders[idx]
		o.UnitCost = productList[o.PartNumber].UnitCost
		o.UnitPrice = productList[o.PartNumber].UnitPrice
	}

	return orders

}

func load(orders []*Order) {
	f, _ := os.Create("./dest.txt")
	defer f.Close()

	fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s\n",
		"Part Number", "Quantity", "Unit Cost",
		"Unit Price", "Total Cost", "Total Price")

	for _, o := range orders {
		time.Sleep(1 * time.Millisecond)
		fmt.Fprintf(f, "%20s%15d%12.2f%12.2f%15.2f%15.2f\n",
			o.PartNumber, o.Quantity, o.UnitCost, o.UnitPrice,
			o.UnitCost*float64(o.Quantity),
			o.UnitPrice*float64(o.Quantity))
	}
}
