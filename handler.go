package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type MyHandler struct {
	service Service
}

func NewMyHandler(service Service) *MyHandler {
	return &MyHandler{service: service}
}

func (h *MyHandler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Query parameter is required", http.StatusBadRequest)
		return
	}

	vacancies, err := h.service.SearchVacancy(query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error searching for vacancies: %v", err), http.StatusInternalServerError)
		return
	}

	_, err = h.service.SearchVacancy(query)
	if err != nil {
		fmt.Printf("Error saving search history: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vacancies)
}

func (h *MyHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vacancyID, ok := vars["id"]
	if !ok {
		http.Error(w, "Vacancy ID is required", http.StatusBadRequest)
		return
	}

	vacancy, err := h.service.GetVacancy(vacancyID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting vacancy: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vacancy)
}

func (h *MyHandler) ListHandler(w http.ResponseWriter, r *http.Request) {
	vacancies, err := h.service.ListVacancies()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting list of vacancies: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vacancies)
}

func (h *MyHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vacancyID, ok := vars["id"]
	if !ok {
		http.Error(w, "Vacancy ID is required", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteVacancy(vacancyID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting vacancy: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *MyHandler) ListSearchHistoryHandler(w http.ResponseWriter, r *http.Request) {
	history, err := h.service.ListSearchHistory()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting search history: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}

func (h *MyHandler) DeleteSearchHistoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	historyID, ok := vars["id"]
	if !ok {
		http.Error(w, "History ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(historyID)
	if err != nil {
		http.Error(w, "Invalid History ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteSearchHistory(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting search history: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
