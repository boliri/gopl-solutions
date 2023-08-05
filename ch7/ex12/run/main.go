// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"ecommerce"
)

//go:embed tpl/list.html
var rawTpl string
var tableTpl *template.Template = template.Must(template.New("items").Parse(rawTpl))

var db = ecommerce.NewDatabase()

func main() {
	http.HandleFunc("/", landing)
	http.HandleFunc("/create", create)
	http.HandleFunc("/get", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func landing(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	tableTpl.Execute(w, db.GetAll())
}

func create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if item == "" {
		http.Error(w, "missing \"item\" parameter", http.StatusBadRequest)
		return
	}

	if price == "" {
		http.Error(w, "missing \"price\" parameter", http.StatusBadRequest)
		return
	}

	priceF, err := strconv.ParseFloat(price, 32)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid value for \"price\" parameter: %s", price), http.StatusBadRequest)
		return
	}

	err = db.Insert(item, float32(priceF))
	if err != nil {
		switch err.(type) {
		case ecommerce.ItemAlreadyExists:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "item \"%s\" inserted", item)
}

func read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if item == "" {
		landing(w, req)
		return
	}

	entry, err := db.Get(item)

	if err != nil {
		switch err.(type) {
		case ecommerce.MissingItem:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	tableTpl.Execute(w, entry)
}

func update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if item == "" {
		http.Error(w, "missing \"item\" parameter", http.StatusBadRequest)
		return
	}

	if price == "" {
		http.Error(w, "missing \"price\" parameter", http.StatusBadRequest)
		return
	}

	priceF, err := strconv.ParseFloat(price, 32)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid value for \"price\" parameter: %s", price), http.StatusBadRequest)
		return
	}

	err = db.Update(item, float32(priceF))
	if err != nil {
		switch err.(type) {
		case ecommerce.MissingItem:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "item \"%s\" updated", item)
}

func delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if item == "" {
		http.Error(w, "missing \"item\" parameter", http.StatusBadRequest)
		return
	}

	err := db.Delete(item)
	if err != nil {
		switch err.(type) {
		case ecommerce.MissingItem:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "item \"%s\" deleted", item)
}
