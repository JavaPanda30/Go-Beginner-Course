package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHttpHandleFunc(s.handleAccount)).Methods("GET", "POST")
	router.HandleFunc("/account/{id}", makeHttpHandleFunc(s.handleAccount)).Methods("GET", "DELETE")
	router.HandleFunc("/transfer", makeHttpHandleFunc(s.handleTransfer)).Methods("POST")

	log.Printf("Server Running on Port %v", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}




func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAccount(w, r)
	case "POST":
		return s.handleCreateAccount(w, r)
	case "DELETE":
		return s.handleDeleteAccount(w, r)
	default:
		return fmt.Errorf("method not allowed: %s", r.Method)
	}
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		return fmt.Errorf("account ID required")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid account ID")
	}

	account, err := s.store.GetAccountByID(id)
	if err != nil {
		return fmt.Errorf("account not found")
	}

	return writeJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	var acc Account
	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		return fmt.Errorf("invalid request payload")
	}

	if err := s.store.CreateAccount(&acc); err != nil {
		return fmt.Errorf("failed to create account")
	}

	return writeJSON(w, http.StatusCreated, acc)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid account ID")
	}

	if err := s.store.DeleteAccount(id); err != nil {
		return fmt.Errorf("failed to delete account")
	}

	return writeJSON(w, http.StatusOK, map[string]string{"message": "account deleted"})
}






func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	var transfer struct {
		FromID int64 `json:"from_id"`
		ToID   int64 `json:"to_id"`
		Amount int64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&transfer); err != nil {
		return fmt.Errorf("invalid request payload")
	}

	if transfer.Amount <= 0 {
		return fmt.Errorf("transfer amount must be greater than zero")
	}

	fromAcc, err := s.store.GetAccountByID(transfer.FromID)
	if err != nil {
		return fmt.Errorf("sender account not found")
	}

	if fromAcc.Balance < transfer.Amount {
		return fmt.Errorf("insufficient balance")
	}

	toAcc, err := s.store.GetAccountByID(transfer.ToID)
	if err != nil {
		return fmt.Errorf("receiver account not found")
	}

	fromAcc.Balance -= transfer.Amount
	toAcc.Balance += transfer.Amount

	if err := s.store.UpdateAccount(fromAcc); err != nil {
		return fmt.Errorf("failed to update sender account")
	}
	if err := s.store.UpdateAccount(toAcc); err != nil {
		return fmt.Errorf("failed to update receiver account")
	}

	return writeJSON(w, http.StatusOK, map[string]string{"message": "transfer successful"})
}






func writeJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
