package main

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// product defines the Product model
type product struct {
	gorm.Model
	Title string `json:"title"`
	Code  string `json:"code"`
	Price int    `json:"price"`
}

func (p *product) fineOne(id int) int64 {
	res := db.First(&p, "id = ?", id)
	return res.RowsAffected
}

func (p *product) update(tp product) {
	p.Title = tp.Title
	p.Code = tp.Code
	p.Price = tp.Price

	db.Save(&p)
}

func listAllProducts(w http.ResponseWriter, r *http.Request) {
	products := []product{}
	db.Find(&products)

	res := response{}
	res.statusCode(200).data(products).serve(w)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	p := product{}
	if err := parseBody(r, &p); err != nil {
		panic(err)
	}

	db.Create(&p)

	res := response{}
	res.statusCode(201).data(p).serve(w)
}

func showSingleProduct(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	id, _ := strconv.Atoi(q["id"][0])

	p := &product{}
	no := p.fineOne(id)
	if no <= 0 {
		p = nil
	}

	res := response{}
	res.statusCode(200).data(p).serve(w)
}

func updateSingleProduct(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	id, _ := strconv.Atoi(q["id"][0])

	tp := product{}
	if err := parseBody(r, &tp); err != nil {
		panic(err)
	}

	p := &product{}
	no := p.fineOne(id)
	if no <= 0 {
		p = nil
	}

	p.update(tp)

	res := response{}
	res.statusCode(201).data(p).serve(w)
}
func deleteSingleProduct(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	id, _ := strconv.Atoi(q["id"][0])

	p := &product{}
	no := p.fineOne(id)
	if no <= 0 {
		p = nil
	}

	db.Delete(&p)

	res := response{}
	res.statusCode(201).data("product deleted").serve(w)
}
