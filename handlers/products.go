package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/crazyoptimist/go-microservice-poc/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, conErr := strconv.Atoi(vars["id"])
	if conErr != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
	}

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	updateErr := data.UpdateProduct(id, prod)
	if updateErr == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if updateErr != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// validate the product
		validationErr := prod.Validate()
		if validationErr != nil {
			http.Error(rw, fmt.Sprintf("Error validating product: %s", validationErr), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
