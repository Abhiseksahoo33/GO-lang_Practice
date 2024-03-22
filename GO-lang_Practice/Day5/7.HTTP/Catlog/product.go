package main

import (
	"fmt"
	"time"
)

type Product struct {
	Pid      int       `json:"pid"`
	Name     string    `json:"name"`
	Catagory string    `json:"catagory"`
	Price    float64   `json:"price"`
	Expiry   time.Time `json:"expiry"`
	Desc     string    `json:"desc"`
}

var ProductDetails []Product = []Product{Product{Pid: 1, Name: "laptop", Catagory: "Electronics", Price: 145000.00, Expiry: time.Now().Add(365 * time.Hour * 24)},
	Product{Pid: 2, Name: "Playstation", Catagory: "Electronics", Price: 216255.00, Expiry: time.Now().Add(365 * time.Hour * 24)}}

func GetAllProduct() []Product {
	return ProductDetails
}

func GetProduct(pid int) (Product, error) {
	for _, v := range ProductDetails {
		if v.Pid == pid {
			return v, nil
		}
	}
	p := Product{}
	return p, fmt.Errorf("Not Found")
}

func InsertProduct(p Product) error {

	p, e := GetProduct(p.Pid)
	if e == nil {
		return fmt.Errorf("Product Already Exist")
	} else {
		ProductDetails = append(ProductDetails, p)
	}

	return nil
}

func DeleteProduct(pid int) error {
	_, e := GetProduct(pid)
	if e != nil {
		return fmt.Errorf("Product does not Exist!!!")
	} else {
		idx := -1
		for i, v := range ProductDetails {
			if v.Pid == pid {
				idx = i
				break
			}
		}
		if idx == -1 {
			return fmt.Errorf("Internal Server Error")
		}
		ProductDetails = append(ProductDetails[:idx-1], ProductDetails[idx+1:]...)
	}
	return nil
}
