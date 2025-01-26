package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type VMCreateRequest struct {
	Name     string `json:"name"`
	Capacity string `json:"capacity"`
	Region   string `json:"region"`
}

type SQLMigrateRequest struct {
	ServerID         string `json:"server_id"`
	SourceRegion     string `json:"source_region"`
	DestinationRegion string `json:"destination_region"`
}

type VMMigrateRequest struct {
	VMID            string `json:"vm_id"`
	SourceRegion    string `json:"source_region"`
	DestinationRegion string `json:"destination_region"`
}

func createVM(w http.ResponseWriter, r *http.Request) {
	var req VMCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Simulate duplicate resource creation check
	if req.Name == "existing-vm" {
		http.Error(w, "VM already exists", http.StatusConflict)
		return
	}

	// Simulate VM creation logic
	log.Printf("Creating VM: %+v", req)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "VM created successfully"})
}

func migrateSQL(w http.ResponseWriter, r *http.Request) {
	var req SQLMigrateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Simulate migration logic
	log.Printf("Migrating SQL Server: %+v", req)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "SQL server migration initiated"})
}

func migrateVM(w http.ResponseWriter, r *http.Request) {
	var req VMMigrateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Simulate migration logic
	log.Printf("Migrating VM: %+v", req)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "VM migration initiated"})
}

func main() {
	http.HandleFunc("/api/v1/vm/create", createVM)
	http.HandleFunc("/api/v1/sql/migrate", migrateSQL)
	http.HandleFunc("/api/v1/vm/migrate", migrateVM)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

