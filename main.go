package main

import (
	"fmt"
	"log"
	"net/http"

	"pos-backend/database"
	"pos-backend/routes"
)

func main() {
	// เชื่อมต่อฐานข้อมูล
	database.Connect()

	// สร้าง Router
	router := routes.SetupRouter()

	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
