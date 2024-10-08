package handlers

import (
	"encoding/json"
	"log"
	"main/repos"
	"net/http"
	"strconv"
)

type SupplierHandler struct {
	repo *repos.SupplierRepo
}

func NewSupplierHandler(repo *repos.SupplierRepo) *SupplierHandler {
	return &SupplierHandler{repo: repo}
}

func (h *SupplierHandler) GetSuppliersHandler(w http.ResponseWriter, r *http.Request) {
	// Parse pagination parameters from query string
	pageParam := r.URL.Query().Get("page")
	pageSizeParam := r.URL.Query().Get("pageSize")

	// Convert the query parameters to integers (default to page 1 and pageSize 10 if not provided)
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(pageSizeParam)
	if err != nil || pageSize < 1 {
		pageSize = 9
	}

	suppliers, err := h.repo.PaginatedGetSuppliers(page, pageSize)
	if err != nil {
		http.Error(w, "Failed to fetch suppliers", http.StatusInternalServerError)
		return
	}
	log.Printf("Fetched %d suppliers", len(suppliers))

	//the total suppliers is 104
	//totalSuppliers, err := h.repo.GetTotalOfSuppliers()
	//if err != nil {
	//	http.Error(w, "Failed to fetch total suppliers", http.StatusInternalServerError)
	//	return
	//}

	// Paginate the suppliers
	//start := (page - 1) * pageSize
	//
	//if start >= totalSuppliers {
	//	http.Error(w, "No more suppliers to fetch", http.StatusNotFound)
	//	return
	//}
	//end := start + pageSize
	//if end > totalSuppliers {
	//	end = totalSuppliers
	//}
	//paginatedSuppliers := suppliers[start:end]

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(suppliers)
	if err != nil {
		log.Printf("Error encoding suppliers: %v", err)
		http.Error(w, "Failed to fetch suppliers", http.StatusInternalServerError)
		return
	}
}
