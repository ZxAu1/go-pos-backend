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
	mux.HandleFunc("/menu_items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.GetMenuItems(w, r)
		} else {
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

	mux.HandleFunc("/menu_categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// case http.MethodGet:
		// 	controllers.GetMenuCategories(w, r)
		case http.MethodPost:
			controllers.AddMenuCategory(w, r)
		case http.MethodPut:
			controllers.UpdateMenuCategory(w, r)
		case http.MethodDelete:
			controllers.DeleteMenuCategory(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
