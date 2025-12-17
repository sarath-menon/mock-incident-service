package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const API_KEY = "sk-1234567890abcdefghijklmnop"  // Line 10: Hardcoded API key

func handleLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Line 16: SQL Injection vulnerability
	query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s'", username, password)
	db, _ := sql.Open("mysql", "connection_string")
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Login failed for user %s, SSN: %s", username, r.FormValue("ssn"))  // Line 21: PII exposure in logs
	}
	defer rows.Close()
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	// Line 26: Missing rate limiter - vulnerable to abuse
	data := processRequest(r)
	w.Write([]byte(data))
}

func processRequest(r *http.Request) string {
	return "processed"
}
