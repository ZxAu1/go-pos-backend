package controllers

import (
	"encoding/json"
	"net/http"

	"strconv"

	"pos-backend/database"
	"pos-backend/models"
)

func AddMenuCategory(w http.ResponseWriter, r *http.Request) {
	var mc models.MenuCategory
	if err := json.NewDecoder(r.Body).Decode(&mc); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// แปลง mc.Name (map) เป็น JSON string
	nameBytes, err := json.Marshal(mc.Name)
	if err != nil {
		http.Error(w, "Failed to marshal name", http.StatusInternalServerError)
		return
	}
	nameJSONString := string(nameBytes)

	res, err := database.DB.Exec("INSERT INTO menu_categories (name) VALUES (?)", nameJSONString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	mc.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mc)
}

// UpdateMenuCategory แก้ไขหมวดหมู่ตาม id
func UpdateMenuCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid menu category ID", http.StatusBadRequest)
		return
	}

	var mc models.MenuCategory
	if err := json.NewDecoder(r.Body).Decode(&mc); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	nameBytes, err := json.Marshal(mc.Name)
	if err != nil {
		http.Error(w, "Failed to marshal name", http.StatusInternalServerError)
		return
	}
	nameJSONString := string(nameBytes)

	_, err = database.DB.Exec("UPDATE menu_categories SET name=? WHERE id=?", nameJSONString, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mc.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mc)
}

// GetCustomers returns all customers
func GetMenuItems(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT id, category_id, option_ids, name, price, discounted_price, is_orderable, status, description 
		FROM menu_items
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []models.MenuItem
	for rows.Next() {
		var item models.MenuItem
		if err := rows.Scan(
			&item.ID,
			&item.CategoryID,
			&item.OptionIDs,
			&item.Name,
			&item.Price,
			&item.DiscountedPrice,
			&item.IsOrderable,
			&item.Status,
			&item.Description,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
func DeleteMenuCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	result, err := database.DB.Exec("DELETE FROM menu_categories WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Unable to verify deletion", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "No record found to delete", http.StatusNotFound)
		return
	}

	// ส่ง JSON กลับเมื่อสำเร็จ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "success",
	})
}
