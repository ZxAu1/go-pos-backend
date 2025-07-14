package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"pos-backend/database"
	"pos-backend/models"
)

// GetCustomers returns all customers
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT customer_id, name, active, price_child, price_adult FROM customer")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var c models.Customer
		if err := rows.Scan(&c.CustomerID, &c.Name, &c.Active, &c.PriceChild, &c.PriceAdult); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		customers = append(customers, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

// AddCustomer creates a new customer
func AddCustomer(w http.ResponseWriter, r *http.Request) {
	var c models.Customer
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	res, err := database.DB.Exec(
		"INSERT INTO customer (name, active, price_child, price_adult) VALUES (?, ?, ?, ?)",
		c.Name, c.Active, c.PriceChild, c.PriceAdult,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	c.CustomerID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

// UpdateCustomer updates an existing customer
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	var c models.Customer
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(
		"UPDATE customer SET name=?, active=?, price_child=?, price_adult=? WHERE customer_id=?",
		c.Name, c.Active, c.PriceChild, c.PriceAdult, id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.CustomerID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}

// DeleteCustomer removes a customer
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("DELETE FROM customer WHERE customer_id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
