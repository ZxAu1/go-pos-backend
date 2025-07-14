package routes

import (
	"net/http"

	"pos-backend/controllers"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetCustomers(w, r)
		case http.MethodPost:
			controllers.AddCustomer(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/customer", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			controllers.UpdateCustomer(w, r)
		case http.MethodDelete:
			controllers.DeleteCustomer(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
